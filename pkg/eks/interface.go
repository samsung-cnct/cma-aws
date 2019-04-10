package eks

type EKSCluster interface {
	// CreateCluster creates an eks cluster
	CreateCluster(CreateClusterInput) (CreateClusterOutput, error)
	// GetCluster returns the eks status and Cluster info
	GetCluster(GetClusterInput) (GetClusterOutput, error)
	// ListClusters retuns a list of ClusterOutput
	ListClusters(ListClustersInput) (ListClustersOutput, error)
	// DeleteCluster deletes an eks cluster
	DeleteCluster(DeleteClusterInput) (DeleteClusterOutput, error)
	// CreateNodeGroup creates an eks nodepool group
	CreateNodeGroup(CreateNodeGroupInput) error
	// DeleteNodeGroup deletes an eks nodepool group
	DeleteNodeGroup(DeleteNodeGroupInput) error
	// ScaleNodeGroup scales an eks nodepool group
	ScaleNodeGroup(ScaleNodeGroupInput) error
	// GetUpgradeVersions returns the valid eks k8s versions
	GetUpgradeVersions(GetUpgradeVersionsInput) (GetUpgradeVersionsOutput, error)
	// UpgradeCluster upgrades the version
	UpgradeCluster(UpgradeClusterInput) error
	// GetKubeConfig writes the kubeconfig to the specified path
	GetKubeConfig(KubeConfigPath string) error
}
