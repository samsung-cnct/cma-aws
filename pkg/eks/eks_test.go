package eks

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestCreateEksBadToken(t *testing.T) {
	ekscluster := New(AwsCredentials{
		AccessKeyId:     "BOGUS",
		SecretAccessKey: "BOGUS",
		Region:          "us-east-1",
	}, "ekstest1", "us-west-2", "/tmp/kubeconfig")

	azs := []string{"us-west-2b", "us-west-2c", "us-west-2d"}

	nodepools := make([]NodePool, 1)
	nodepools[0].Name = "nodepool1"
	nodepools[0].Nodes = 2
	nodepools[0].Type = "m5.large"

	output, err := ekscluster.CreateCluster(CreateClusterInput{
		Name:              ekscluster.clusterName,
		Version:           "1.11",
		Region:            ekscluster.region,
		AvailabilityZones: azs,
		NodePools:         nodepools,
	})
	fmt.Printf("output = %s", output.CmdOutput)
	if err == nil {
		t.Errorf("want error, got nil")
	}
}

func TestGetClusterBadToken(t *testing.T) {
	eks := New(AwsCredentials{
		AccessKeyId:     "BOGUS",
		SecretAccessKey: "BOGUS",
		Region:          "us-east-1",
	}, "ekstest1", "us-west-2", "")

	output, err := eks.GetCluster(GetClusterInput{
		Name:   eks.clusterName,
		Region: eks.region,
	})
	fmt.Printf("output = %s", output.CmdOutput)

	if err == nil {
		t.Errorf("want error, got nil")
	}
}

type MockEksNotFound struct {
	awsCreds       AwsCredentials
	clusterName    string
	region         string
	kubeConfigPath string
}

func (m *MockEksNotFound) GetCluster(in GetClusterInput) (GetClusterOutput, error) {
	return GetClusterOutput{CmdOutput: "[✖]  unable to describe control plane 'foo1': ResourceNotFoundException: No cluster found for name: foo1. status code: 404, request id: 851e214f-561f-11e9-bd9d-a1debad5d25d"},
		errors.New("exit status 1")
}
func TestGetClusterNotFound(t *testing.T) {
	eks := MockEksNotFound{
		awsCreds: AwsCredentials{
			AccessKeyId:     "BOGUS",
			SecretAccessKey: "BOGUS",
			Region:          "us-east-1",
		},
		clusterName:    "foo1",
		region:         "us-west-2",
		kubeConfigPath: ""}

	output, err := eks.GetCluster(GetClusterInput{
		Name:   eks.clusterName,
		Region: eks.region,
	})
	fmt.Printf("output = %s error = %v", output.CmdOutput, err)
	if err == nil {
		t.Errorf("want error, got nil")
	}
	if strings.Contains(output.CmdOutput, "status code: 404") == false {
		t.Errorf("did not find expected 404 status code")
	}
}

func TestGetKubeConfigDataBadToken(t *testing.T) {
	eks := New(AwsCredentials{
		AccessKeyId:     "BOGUS",
		SecretAccessKey: "BOGUS",
		Region:          "us-east-1",
	}, "ekstest1", "us-west-2", "")

	output, err := eks.GetKubeConfigData("/tmp/test1config")
	fmt.Printf("output = %s", output.String())

	if err == nil {
		t.Errorf("want error, got nil")
	}
}

func TestClusterExistsBadCreds(t *testing.T) {
	eks := New(AwsCredentials{
		AccessKeyId:     "BOGUS",
		SecretAccessKey: "BOGUS",
		Region:          "us-east-1",
	}, "ekstest1", "us-west-2", "")

	exists, output := eks.ClusterExists(eks.clusterName, eks.region)
	fmt.Printf("exists = %v", exists)

	if exists {
		t.Errorf("expecting exists == false")
	}
	if !strings.Contains(output, "InvalidClientTokenId") {
		t.Errorf("expecting InvalidClientTokenId")
	}
}
