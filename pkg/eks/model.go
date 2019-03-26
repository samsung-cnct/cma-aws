package eks

// CreateClusterInput is used to create a new eks cluster
type CreateClusterInput struct {
	Name    string
	Version string
	Region  string
	// SSHKeyName of a keypair in AWS credential Region
	SSHKeyName string
	// AvailabilityZones is optional, eks can select by region
	AvailabilityZones []string
	// NodePools is the worker node pool
	NodePools []NodePool
}

// CreateClusterOutput contains output from the create command
type CreateClusterOutput struct {
	CmdOutput string
}

// NodePool is the worker node pool
type NodePool struct {
	Name string
	// Type is the EC2 instance type
	Type string
	// Nodes is the number of desired nodes
	Nodes int32
	// Min and Max Nodes is used for autoscaling
	MinNodes int32
	MaxNodes int32
}

// GetClusterInput requires the cluster name and region
type GetClusterInput struct {
	Name   string
	Region string
}

// GetClusterOutput shows the eks cluster output fields
type GetClusterOutput struct {
	Name             string
	Version          string
	Status           string
	CreatedTimestamp string
	Vpc              string
	Subnets          []string
	SecurityGroups   []string
}

// ListClustersInput returns a list of all clusters
type ListClustersInput struct {
	// Region is optional here, otherwise all regions checked
	Region string
}

// ListClusterId identifies a cluster
type ListClusterId struct {
	Name   string
	Region string
}

// ListClustersOutput returns a list of ListClusterId
type ListClustersOutput struct {
	Clusters []ListClusterId
}

// DeleteClusterInput is used to delete a cluster
type DeleteClusterInput struct {
	Name   string
	Region string
}

// CreateNodeGroupInput is used to create a worker nodepool group
type CreateNodeGroupInput struct {
	ClusterName string
	Region      string
	Version     string
	NodePool    NodePool
}

// DeleteNodeGroupInput is used to delete a worker nodepool group
// Note: DeleteCluster will delete everything, including the worker nodepool
type DeleteNodeGroupInput struct {
	ClusterName string
	Region      string
	Name        string
}

// ScaleNodeGroupInput is used to change the size of a worker nodepool
type ScaleNodeGroupInput struct {
	ClusterName string
	Region      string
	Name        string
	Nodes       int32
}

// GetUpgradeVersionsInput requests a list of k8s versions supported
type GetUpgradeVersionsInput struct {
	Region string
}

// GetUpgradeVersionsOutput is used to return the supported versions
type GetUpgradeVersionsOutput struct {
	Versions []string
}

// UpgradeClusterInput is used to request a cluster upgrade
type UpgradeClusterInput struct {
	ClusterName string
	Region      string
	Version     string
}

// AwsCredentials contains the credentials to access the AWS api
type AwsCredentials struct {
	// The AccessKeyId for API Access
	AccessKeyId string
	// The SecretAccessKey for API access
	SecretAccessKey string
	// The Region for API access
	Region string
}
