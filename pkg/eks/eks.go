package eks

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	DeleteClusterTimeoutSeconds = 1200
	GetClusterTimeoutSeconds    = 30
)

type EKS struct {
	awsCreds       AwsCredentials
	clusterName    string
	region         string
	kubeConfigPath string
}

func init() {
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
	cmd := cmd.New("eksctl",
		e.GetCreateClusterArgs(in),
		time.Duration(CreateClusterTimeoutSeconds)*time.Second,
		[]string{"AWS_ACCESS_KEY_ID=" + e.awsCreds.AccessKeyId,
			"AWS_SECRET_ACCESS_KEY=" + e.awsCreds.SecretAccessKey,
			"AWS_DEFAULT_REGION=" + e.awsCreds.Region},
	)

	output, err := cmd.Run()
	if err != nil {
		logger.Errorf("CreateCluster error running eksctl command: %v", err)
		return CreateClusterOutput{CmdOutput: output.String()}, err
	}

	// TODO: Future work - add additional nodepools (if more than one)

	return CreateClusterOutput{CmdOutput: output.String()}, nil
}

func (e *EKS) GetKubeConfigArgs(kubeconfigPath string) []string {
	args := make([]string, 0, MaxCmdArgs)
	args = append(args, "utils")
	args = append(args, "write-kubeconfig")
	// name
	args = append(args, "--name")
	args = append(args, e.clusterName)
	// datacenter region
	args = append(args, "--region")
	args = append(args, e.region)
	// kubeconfig path
	args = append(args, "--kubeconfig")
	args = append(args, kubeconfigPath)
	return args
}

// GetKubeConfig writes the kubeconfig to the input file path
func (e *EKS) GetKubeConfig(KubeConfigPath string) error {
	// get kubeconfig command
	cmd := cmd.New("eksctl",
		e.GetKubeConfigArgs(KubeConfigPath),
		time.Duration(GetClusterTimeoutSeconds)*time.Second,
		[]string{"AWS_ACCESS_KEY_ID=" + e.awsCreds.AccessKeyId,
			"AWS_SECRET_ACCESS_KEY=" + e.awsCreds.SecretAccessKey,
			"AWS_DEFAULT_REGION=" + e.awsCreds.Region},
	)
	output, err := cmd.Run()
	if err != nil {
		logger.Errorf("GetKubeConfig error running eksctl command: %v, output = %s",
			err, output.String())
		return err
	}
	return nil
}

// GetKubeConfigData returns the kubeconfig file as a byte buffer
func (e *EKS) GetKubeConfigData(KubeConfigPath string) (*bytes.Buffer, error) {
	err := e.GetKubeConfig(KubeConfigPath)
	if err != nil {
		return new(bytes.Buffer), err
	}
	// cat the file by running a command
	cmd := cmd.New("cat",
		[]string{KubeConfigPath},
		time.Duration(GetClusterTimeoutSeconds)*time.Second,
		nil,
	)
	output, err := cmd.Run()
	if err != nil {
		return &output, err
	}
	return &output, nil
}

func (e *EKS) GetClusterArgs(clusterName string, dataCenterRegion string) []string {
	args := make([]string, 0, MaxCmdArgs)
	args = append(args, "get")
	args = append(args, "cluster")
	// name
	args = append(args, "--name")
	args = append(args, clusterName)
	// datacenter region
	args = append(args, "--region")
	args = append(args, dataCenterRegion)
	// output json format
	args = append(args, "-o")
	args = append(args, "json")
	return args
}

// ClusterExists returns a bool and the command output string for get cluster
func (e *EKS) ClusterExists(clusterName string, dataCenterRegion string) (bool, string) {
	output, err := e.GetCluster(GetClusterInput{
		Name:   clusterName,
		Region: dataCenterRegion,
	})
	if err == nil {
		// found the cluster, the normal case when it exists
		return true, output.CmdOutput
	}
	return false, output.CmdOutput
}

func (e *EKS) DescribeStacksArgs(clusterName string, dataCenterRegion string) []string {
	args := make([]string, 0, MaxCmdArgs)
	args = append(args, "utils")
	args = append(args, "describe-stacks")
	// name
	args = append(args, "--name")
	args = append(args, clusterName)
	// datacenter region
	args = append(args, "--region")
	args = append(args, dataCenterRegion)
	return args
}

// ClusterCreateInProgress should be checked if GetCluster returns a ResourceNotFoundException error.
// There is a period of time eksctl get cluster will return a ResourceNotFoundException error while
// the cluster is starting to provision.
func (e *EKS) ClusterCreateInProgress(clusterName string, dataCenterRegion string) (bool, string) {
	// describe-stacks command
	cmd := cmd.New("eksctl",
		e.DescribeStacksArgs(clusterName, dataCenterRegion),
		time.Duration(GetClusterTimeoutSeconds)*time.Second,
		[]string{"AWS_ACCESS_KEY_ID=" + e.awsCreds.AccessKeyId,
			"AWS_SECRET_ACCESS_KEY=" + e.awsCreds.SecretAccessKey,
			"AWS_DEFAULT_REGION=" + e.awsCreds.Region},
	)

	output, err := cmd.Run()
	if err != nil {
		logger.Infof("ClusterCreateInProgress error from describe-stacks: %v", err)
		return false, output.String()
	}
	if strings.Contains(output.String(), "CREATE_IN_PROGRESS") {
		logger.Infof("ClusterCreateInProgress == true")
		return true, output.String()
	}
	return false, output.String()
}

// GetCluster returns eks cluster status
func (e *EKS) GetCluster(in GetClusterInput) (GetClusterOutput, error) {
	cmd := cmd.New("eksctl",
		e.GetClusterArgs(in.Name, in.Region),
		time.Duration(GetClusterTimeoutSeconds)*time.Second,
		[]string{"AWS_ACCESS_KEY_ID=" + e.awsCreds.AccessKeyId,
			"AWS_SECRET_ACCESS_KEY=" + e.awsCreds.SecretAccessKey,
			"AWS_DEFAULT_REGION=" + e.awsCreds.Region},
	)

	output, err := cmd.Run()
	if err != nil {
		logger.Infof("GetCluster error from get cluster command: %v", err)
		return GetClusterOutput{CmdOutput: output.String()}, err
	}

	// remove outer array brackets so we can parse using an unstructured map interface
	buffer := output.Bytes()
	var s string
	if (len(output.String()) > 1) && (output.String()[0] == '[') {
		s = string(buffer[1 : len(buffer)-1])
	} else {
		s = string(buffer)
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(s), &result)

	return GetClusterOutput{
		Name:             fmt.Sprintf("%v", result["Name"]),
		Version:          fmt.Sprintf("%v", result["Version"]),
		Status:           fmt.Sprintf("%v", result["Status"]),
		CreatedTimestamp: fmt.Sprintf("%v", result["CreatedAt"]),
		CmdOutput:        output.String()}, nil
}

func (e *EKS) DeleteClusterArgs(clusterName string, dataCenterRegion string) []string {
	args := make([]string, 0, MaxCmdArgs)
	args = append(args, "delete")
	args = append(args, "cluster")
	// name
	args = append(args, "--name")
	args = append(args, clusterName)
	// datacenter region
	args = append(args, "--region")
	args = append(args, dataCenterRegion)
	return args
}

func (e *EKS) DeleteCluster(in DeleteClusterInput) (DeleteClusterOutput, error) {
	cmd := cmd.New("eksctl",
		e.DeleteClusterArgs(in.Name, in.Region),
		time.Duration(DeleteClusterTimeoutSeconds)*time.Second,
		[]string{"AWS_ACCESS_KEY_ID=" + e.awsCreds.AccessKeyId,
			"AWS_SECRET_ACCESS_KEY=" + e.awsCreds.SecretAccessKey,
			"AWS_DEFAULT_REGION=" + e.awsCreds.Region},
	)

	output, err := cmd.Run()
	if err != nil {
		logger.Errorf("DeleteCluster error running eksctl command: %v", err)
	}

	return DeleteClusterOutput{CmdOutput: output.String()}, err
}
