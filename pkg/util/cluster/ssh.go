package cluster

import (
	"github.com/samsung-cnct/cma-aws/pkg/util/awsutil"
	"github.com/samsung-cnct/cma-aws/pkg/util/awsutil/models"
	"github.com/samsung-cnct/cma-aws/pkg/util/k8s"
	"github.com/spf13/viper"
)

const (
	SSHK8SSecretSuffix = "-ssh"
)

func generateSSHSecretKey(clusterName string) string {
	return clusterName + SSHK8SSecretSuffix
}

func ProvisionAndSaveSSHKey(clusterName string, credentials awsmodels.Credentials) (string, error) {
	privateKey, err := awsutil.CreateKey(clusterName, credentials)
	if err != nil {
		return "", err
	}

	err = k8sutil.CreateSSHSecret(generateSSHSecretKey(clusterName), viper.GetString("kubernetes-namespace"), []byte(privateKey))
	if err != nil {
		// Let's try to roll back the create key on AWS...
		_ = awsutil.DeleteKey(clusterName, credentials)
		return "", err
	}

	return clusterName, nil
}

func RemoveSSHKey(clusterName string, credentials awsmodels.Credentials) error {
	err := awsutil.DeleteKey(clusterName, credentials)
	if err != nil {
		return err
	}
	sshSecretName := clusterName + SSHK8SSecretSuffix
	err = k8sutil.DeleteSSHSecret(sshSecretName, viper.GetString("kubernetes-namespace"))
	if err != nil {
		return err
	}
	return nil
}
