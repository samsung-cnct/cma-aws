package apiserver

import (
	"io/ioutil"
	"os"
	"strings"

	eks "github.com/samsung-cnct/cma-aws/pkg/eks"
	pb "github.com/samsung-cnct/cma-aws/pkg/generated/api"
	"github.com/samsung-cnct/cma-aws/pkg/util/awsutil"
	"github.com/samsung-cnct/cma-aws/pkg/util/awsutil/models"
	"github.com/samsung-cnct/cma-aws/pkg/util/cluster"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func matchEksStatus(status string) pb.ClusterStatus {
	switch status {
	case "ACTIVE":
		return pb.ClusterStatus_RUNNING
	case "CREATING":
		return pb.ClusterStatus_PROVISIONING
	default:
		return pb.ClusterStatus_STATUS_UNSPECIFIED
	}
}

func (s *Server) CreateCluster(ctx context.Context, in *pb.CreateClusterMsg) (*pb.CreateClusterReply, error) {
	// Quick validation
	if in.Provider.GetAws() == nil {
		return nil, status.Error(codes.InvalidArgument, "AWS Configuration was not provided")
	}
	if (in.Provider.GetAws().DataCenter.AvailabilityZones != nil) &&
		(len(in.Provider.GetAws().DataCenter.AvailabilityZones) > 0) &&
		(len(in.Provider.GetAws().DataCenter.AvailabilityZones) < 2) {
		return nil, status.Error(codes.InvalidArgument,
			"Must be at least 2 DataCenter.AvailabilityZones, if specified")
	}

	// check if cluster exists or invalid token
	e := eks.New(eks.AwsCredentials{
		AccessKeyId:     in.Provider.GetAws().Credentials.SecretKeyId,
		SecretAccessKey: in.Provider.GetAws().Credentials.SecretAccessKey,
		Region:          in.Provider.GetAws().Credentials.Region,
	}, in.Name, in.Provider.GetAws().DataCenter.Region, "")
	exists, exout := e.ClusterExists(in.Name, in.Provider.GetAws().DataCenter.Region)
	if exists {
		return nil, status.Error(codes.AlreadyExists, "cluster already exists")
	} else {
		if strings.Contains(exout, "InvalidClientTokenId") {
			return nil, status.Error(codes.Unauthenticated, exout)
		}
	}

	// create cluster
	go func() {
		_ = s.CreateEksCluster(ctx, in)
	}()

	return &pb.CreateClusterReply{
		Ok: true,
		Cluster: &pb.ClusterItem{
			Id:     in.Name,
			Name:   in.Name,
			Status: pb.ClusterStatus_PROVISIONING,
		},
	}, nil
}

func (s *Server) CreateEksCluster(ctx context.Context, in *pb.CreateClusterMsg) error {
	// create ssh key in AWS credential region, and store as secret
	credentials := generateCredentials(in.Provider.GetAws().Credentials)
	keyName, err := cluster.ProvisionAndSaveSSHKey(in.Name, credentials)
	if err != nil {
		logger.Errorf("Error creating AWS SSH key: %v", err)
	}

	e := eks.New(eks.AwsCredentials{
		AccessKeyId:     in.Provider.GetAws().Credentials.SecretKeyId,
		SecretAccessKey: in.Provider.GetAws().Credentials.SecretAccessKey,
		Region:          in.Provider.GetAws().Credentials.Region,
	}, in.Name, in.Provider.GetAws().DataCenter.Region, "")

	nodepools := make([]eks.NodePool, len(in.Provider.GetAws().InstanceGroups))
	nodepools[0].Name = in.Provider.GetAws().InstanceGroups[0].Name
	nodepools[0].Nodes = in.Provider.GetAws().InstanceGroups[0].DesiredQuantity
	nodepools[0].Type = in.Provider.GetAws().InstanceGroups[0].Type
	nodepools[0].MinNodes = in.Provider.GetAws().InstanceGroups[0].MinQuantity
	nodepools[0].MaxNodes = in.Provider.GetAws().InstanceGroups[0].MaxQuantity

	createout, err := e.CreateCluster(eks.CreateClusterInput{
		Name:              in.Name,
		Version:           in.Provider.K8SVersion,
		Region:            in.Provider.GetAws().DataCenter.Region,
		SSHKeyName:        keyName,
		AvailabilityZones: in.Provider.GetAws().DataCenter.AvailabilityZones,
		NodePools:         nodepools,
	})
	if err != nil {
		logger.Errorf("CreateCluster Error: %v CmdOutput: %s", err, createout.CmdOutput)

		// store error in threadsafe map for use later by GetCluster
		// GetCluster will report the error and delete it from the map.
		key := in.Name + in.Provider.GetAws().DataCenter.Region + in.Provider.GetAws().Credentials.SecretKeyId
		s.ErrorMap.Store(key, cluster.ErrorValue{
			CmdError:  err,
			CmdOutput: createout.CmdOutput,
		})

		// try to remove ssh key
		ssherr := cluster.RemoveSSHKey(in.Name, credentials)
		if ssherr != nil {
			logger.Warningf("Error removing ssh keys: %v", ssherr)
		}

		return err
	}
	return nil
}

// getErrorCodeString returns the error code and msg string to use for the error
func (s *Server) getErrorCodeString(cmdOutput string, errorMapKey string) (codes.Code, string) {
	var c codes.Code = codes.Unknown
	var msg string = cmdOutput
	if strings.Contains(cmdOutput, "InvalidClientTokenId") {
		return codes.Unauthenticated, cmdOutput
	}
	if strings.Contains(cmdOutput, "AlreadyExistsException") {
		return codes.AlreadyExists, cmdOutput
	}
	if strings.Contains(cmdOutput, "ResourceNotFoundException") {
		c = codes.NotFound
	}
	// check for saved error (create saves error)
	value, found := s.ErrorMap.Load(errorMapKey)
	if found {
		msg = "Error output: " + value.CmdOutput
		s.ErrorMap.Delete(errorMapKey)
	}
	return c, msg
}

func (s *Server) GetCluster(ctx context.Context, in *pb.GetClusterMsg) (*pb.GetClusterReply, error) {
	e := eks.New(eks.AwsCredentials{
		AccessKeyId:     in.Credentials.SecretKeyId,
		SecretAccessKey: in.Credentials.SecretAccessKey,
		Region:          in.Credentials.Region,
	}, in.Name, in.Region, "")
	output, err := e.GetCluster(eks.GetClusterInput{
		Name:   in.Name,
		Region: in.Region,
	})
	if err != nil {
		logger.Errorf("GetCluster Error: %v  CmdOutput: %s", err, output.CmdOutput)
		key := in.Name + in.Region + in.Credentials.SecretKeyId
		c, msg := s.getErrorCodeString(output.CmdOutput, key)
		if c == codes.NotFound {
			// a not found exception is thrown for a short period after create cluster
			clusterCreating, _ := e.ClusterCreateInProgress(in.Name, in.Region)
			if clusterCreating {
				return &pb.GetClusterReply{
					Ok: true,
					Cluster: &pb.ClusterDetailItem{
						Id:         in.Name,
						Name:       in.Name,
						Status:     pb.ClusterStatus_PROVISIONING,
						Kubeconfig: "",
					},
				}, nil
			}
		}
		return nil, status.Error(c, msg)
	}

	// GetKubeConfig
	file, err := ioutil.TempFile("/tmp", in.Name)
	if err != nil {
		logger.Errorf("Error creating tempory file for kubeconfig: %v", err)
	}
	defer os.Remove(file.Name())
	kubeConfigBuf, err := e.GetKubeConfigData(file.Name())
	if err != nil {
		logger.Warningf("GetKubeConfigData returning error: %v", err)
	}

	return &pb.GetClusterReply{
		Ok: true,
		Cluster: &pb.ClusterDetailItem{
			Id:         in.Name,
			Name:       in.Name,
			Status:     pb.ClusterStatus(matchEksStatus(output.Status)),
			Kubeconfig: kubeConfigBuf.String(),
		},
	}, nil
}

func (s *Server) DeleteCluster(ctx context.Context, in *pb.DeleteClusterMsg) (*pb.DeleteClusterReply, error) {
	stackId := in.Name
	credentials := generateCredentials(in.Credentials)

	err := awsutil.DeleteCFStack(stackId, credentials)
	if err != nil {
		return nil, err
	}

	// TODO: Should we continue to clean up if the initial thing fails?
	err = cluster.CleanupClusterInK8s(stackId)
	if err != nil {
		return nil, err
	}
	err = awsutil.DeleteKey(stackId, credentials)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteClusterReply{Ok: true, Status: "Deleting"}, nil
}

func (s *Server) GetClusterList(ctx context.Context, in *pb.GetClusterListMsg) (reply *pb.GetClusterListReply, err error) {
	reply = &pb.GetClusterListReply{}
	return
}

func generateCredentials(credentials *pb.AWSCredentials) awsmodels.Credentials {
	return awsmodels.Credentials{
		Region:          credentials.Region,
		AccessKeyId:     credentials.SecretKeyId,
		SecertAccessKey: credentials.SecretAccessKey,
	}
}
