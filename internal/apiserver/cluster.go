package apiserver

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/cloudformation"
	eks "github.com/samsung-cnct/cma-aws/pkg/eks"
	pb "github.com/samsung-cnct/cma-aws/pkg/generated/api"
	"github.com/samsung-cnct/cma-aws/pkg/util/awsutil"
	"github.com/samsung-cnct/cma-aws/pkg/util/awsutil/models"
	"github.com/samsung-cnct/cma-aws/pkg/util/cluster"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// match aws cluster status to api status enum
func matchStatus(status string) pb.ClusterStatus {
	switch status {
	case cloudformation.StackStatusCreateInProgress:
		return pb.ClusterStatus_PROVISIONING
	case cloudformation.StackStatusReviewInProgress:
		return pb.ClusterStatus_RECONCILING
	case cloudformation.StackStatusRollbackInProgress:
		return pb.ClusterStatus_RECONCILING
	case cloudformation.StackStatusUpdateInProgress:
		return pb.ClusterStatus_RECONCILING
	case cloudformation.StackStatusCreateComplete:
		return pb.ClusterStatus_RUNNING
	case cloudformation.StackStatusUpdateComplete:
		return pb.ClusterStatus_RUNNING
	case cloudformation.StackStatusDeleteInProgress:
		return pb.ClusterStatus_STOPPING
	case cloudformation.StackStatusDeleteComplete:
		return pb.ClusterStatus_STOPPING
	case cloudformation.StackStatusCreateFailed:
		return pb.ClusterStatus_ERROR
	case cloudformation.StackStatusDeleteFailed:
		return pb.ClusterStatus_ERROR
	case cloudformation.StackStatusRollbackFailed:
		return pb.ClusterStatus_ERROR
	case cloudformation.StackStatusRollbackComplete:
		return pb.ClusterStatus_ERROR
	case cloudformation.StackStatusUpdateRollbackFailed:
		return pb.ClusterStatus_ERROR
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

	// TODO: check if cluster exists when GET cluster is ready

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
		fmt.Printf("Error creating AWS SSH key: %v", err)
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
		// log to console
		fmt.Printf("CreateCluster error: %v CmdOutput: %s", err, createout.CmdOutput)

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
			fmt.Printf("Warning: error removing ssh keys: %v", ssherr)
		}

		return err
	}
	return nil
}

func (s *Server) GetCluster(ctx context.Context, in *pb.GetClusterMsg) (*pb.GetClusterReply, error) {
	stackId := in.Name
	credentials := generateCredentials(in.Credentials)

	outputs, err := awsutil.GetHeptioCFStackOutput(stackId, credentials)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil, err
	}

	var kubeconfig []byte
	if outputs.Status == cloudformation.StackStatusCreateComplete {
		kubeconfig, err = cluster.GetKubeConfig(stackId, cluster.SSHConnectionOptions{
			BastionHost: cluster.SSHConnectionHost{
				Hostname:      outputs.BastionHostPublicIp,
				Port:          "22",
				Username:      "ubuntu",
				KeySecretName: stackId,
			},
			TargetHost: cluster.SSHConnectionHost{
				Hostname:      outputs.MasterPrivateIp,
				Port:          "22",
				Username:      "ubuntu",
				KeySecretName: stackId,
			},
		})

		if err != nil {
			logger.Errorf("Error when creating cluster -->%s<-- kubeconfig, error message: %s", stackId, err)
		}
	}

	return &pb.GetClusterReply{
		Ok: true,
		Cluster: &pb.ClusterDetailItem{
			Id:         stackId,
			Name:       stackId,
			Status:     pb.ClusterStatus(matchStatus(outputs.Status)),
			Kubeconfig: string(kubeconfig),
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
