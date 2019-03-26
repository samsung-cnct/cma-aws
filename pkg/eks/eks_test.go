package eks

import (
	"fmt"
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
