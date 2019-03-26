package eks

import (
	"strconv"
	"strings"
	"time"

	"github.com/juju/loggo"
	log "github.com/samsung-cnct/cma-aws/pkg/util"
	"github.com/samsung-cnct/cma-aws/pkg/util/cmd"
)

var (
	logger loggo.Logger
)

const (
	MaxCmdArgs                  = 30
	CreateClusterTimeoutSeconds = 1200
)

type EKS struct {
	awsCreds       AwsCredentials
	clusterName    string
	region         string
	kubeConfigPath string
}

func SetLogger() {
	logger = log.GetModuleLogger("pkg.eks", loggo.INFO)
}

func New(creds AwsCredentials, name string, region string, kubeConfigPath string) *EKS {
	return &EKS{
		awsCreds:       creds,
		clusterName:    name,
		region:         region,
		kubeConfigPath: kubeConfigPath,
	}
}

func (e *EKS) GetCreateClusterArgs(in CreateClusterInput) []string {
	args := make([]string, 0, MaxCmdArgs)
	args = append(args, "create")
	args = append(args, "cluster")
	// name
	args = append(args, "--name")
	args = append(args, in.Name)
	// datacenter region
	args = append(args, "--region")
	args = append(args, in.Region)
	// zones
	if len(in.AvailabilityZones) > 0 {
		args = append(args, "--zones")
		args = append(args, strings.Join(in.AvailabilityZones, ","))
	}
	// version
	args = append(args, "--version")
	args = append(args, in.Version)
	// kubeconfig
	if e.kubeConfigPath != "" {
		args = append(args, "--kubeconfig")
		args = append(args, e.kubeConfigPath)
	}
	// nodegroup-name
	if in.NodePools[0].Name != "" {
		args = append(args, "--nodegroup-name")
		args = append(args, in.NodePools[0].Name)
	}
	// nodes
	if in.NodePools[0].Nodes != 0 {
		args = append(args, "--nodes")
		args = append(args, strconv.FormatInt(int64(in.NodePools[0].Nodes), 10))
	}
	// node-type
	if in.NodePools[0].Type != "" {
		args = append(args, "--node-type")
		args = append(args, in.NodePools[0].Type)
	}
	// nodes-min and nodes-max for auto scaling
	if in.NodePools[0].MaxNodes > 0 {
		args = append(args, "--nodes-min")
		args = append(args, strconv.FormatInt(int64(in.NodePools[0].MinNodes), 10))
		args = append(args, "--nodes-max")
		args = append(args, strconv.FormatInt(int64(in.NodePools[0].MaxNodes), 10))
	}
	return args
}

func (e *EKS) CreateCluster(in CreateClusterInput) (CreateClusterOutput, error) {
	SetLogger()

	// create cluster command
	cmd := cmd.New("eksctl",
		e.GetCreateClusterArgs(in),
		time.Duration(CreateClusterTimeoutSeconds)*time.Second,
		[]string{"AWS_ACCESS_KEY_ID=" + e.awsCreds.AccessKeyId,
			"AWS_SECRET_ACCESS_KEY=" + e.awsCreds.SecretAccessKey,
			"AWS_DEFAULT_REGION=" + e.awsCreds.Region},
	)

	output, err := cmd.Run()
	//logger.Infof("create cmd output is %s", output.String())
	if err != nil {
		logger.Errorf("CreateCluster error running eksctl command: %v", err)
		return CreateClusterOutput{CmdOutput: output.String()}, err
	}

	// TODO: Future work - add additional nodepools (if more than one)

	return CreateClusterOutput{CmdOutput: output.String()}, nil
}
