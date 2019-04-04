package k8sutil

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v12 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type CoreV1SecretInterface interface {
	Secrets(namespace string) v12.SecretInterface
}

type SecretCRUDer interface {
	Create(*corev1.Secret) (*corev1.Secret, error)
	Delete(name string, options *v1.DeleteOptions) error
	Get(name string, options v1.GetOptions) (*corev1.Secret, error)
}

type Secret struct {
	secretAPI CoreV1SecretInterface
}

type SSHSecret struct {
	Secret
}

func NewSSHSecret(secretGetter CoreV1SecretInterface) SSHSecret {
	output := SSHSecret{}
	output.SetSecretGetter(secretGetter)
	return output
}

func (s *Secret) SetSecretGetter(secretAPI CoreV1SecretInterface) {
	s.secretAPI = secretAPI
	return
}

func (s *Secret) GetSecret(namespace string, name string) (secret corev1.Secret, err error) {
	secretResult, err := s.secretAPI.Secrets(namespace).Get(name, v1.GetOptions{})
	if err != nil {
		return
	}
	secret = *secretResult
	return
}

func (s *SSHSecret) Get(namespace string, name string) (secret []byte, err error) {
	secretResult, err := s.GetSecret(namespace, name)
	if err != nil {
		return
	}
	if secretResult.Type != corev1.SecretTypeSSHAuth {
		err = fmt.Errorf("secret %s is not of type %s, but rather is of type %s", name, corev1.SecretTypeSSHAuth, secretResult.Type)
		return
	}
	if len(secretResult.Data[corev1.SSHAuthPrivateKey]) == 0 {
		err = fmt.Errorf("secret %s empty", name)
	}
	secret = secretResult.Data[corev1.SSHAuthPrivateKey]
	return
}

func GetSecretList(namespace string, options v1.ListOptions) (result []corev1.Secret, err error) {
	if DefaultConfig == nil {
		DefaultConfig, _ = GenerateKubernetesConfig()
	}
	client, err := kubernetes.NewForConfig(DefaultConfig)
	if err != nil {
		return nil, err
	}

	secrets, err := client.CoreV1().Secrets(namespace).List(options)
	if err != nil {
		return nil, err
	}
	result = secrets.Items
	return result, nil
}

func DeleteSecret(name string, namespace string) (err error) {
	if DefaultConfig == nil {
		DefaultConfig, _ = GenerateKubernetesConfig()
	}
	client, err := kubernetes.NewForConfig(DefaultConfig)
	if err != nil {
		return err
	}

	err = client.CoreV1().Secrets(namespace).Delete(name, &v1.DeleteOptions{})
	return err
}

func GetSSHSecretList(namespace string) (result []corev1.Secret, err error) {
	listOption := v1.ListOptions{FieldSelector: "type=" + string(corev1.SecretTypeSSHAuth)}
	return GetSecretList(namespace, listOption)
}

func GetSecret(name string, namespace string) (secret corev1.Secret, err error) {
	if DefaultConfig == nil {
		DefaultConfig, _ = GenerateKubernetesConfig()
	}
	client, err := kubernetes.NewForConfig(DefaultConfig)
	if err != nil {
		return
	}

	secretResult, err := client.CoreV1().Secrets(namespace).Get(name, v1.GetOptions{})
	if err != nil {
		return
	}
	secret = *secretResult
	return
}

func GetSSHSecret(name string, namespace string) (secret []byte, err error) {
	secretResult, err := GetSecret(name, namespace)
	if err != nil {
		return
	}
	if secretResult.Type != corev1.SecretTypeSSHAuth {
		err = fmt.Errorf("secret %s is not of type %s, but rather is of type %s", name, corev1.SecretTypeSSHAuth, secretResult.Type)
		return
	}
	secret = secretResult.Data[corev1.SSHAuthPrivateKey]
	return
}

func DeleteSSHSecret(name string, namespace string) (err error) {
	return DeleteSecret(name, namespace)
}

func CreateSSHSecret(name string, namespace string, privateKey []byte) (err error) {
	if DefaultConfig == nil {
		DefaultConfig, _ = GenerateKubernetesConfig()
	}
	client, err := kubernetes.NewForConfig(DefaultConfig)
	if err != nil {
		return err
	}

	labelMap := make(map[string]string)
	labelMap["cmaaws"] = "true"

	dataMap := make(map[string][]byte)
	dataMap[corev1.SSHAuthPrivateKey] = privateKey

	secret := &corev1.Secret{
		ObjectMeta: v1.ObjectMeta{Name: name, Labels: labelMap},
		Type:       corev1.SecretTypeSSHAuth,
		Data:       dataMap,
	}

	_, err = client.CoreV1().Secrets(namespace).Create(secret)
	return err
}

func GetKubeconfigSecretList(namespace string) (result []corev1.Secret, err error) {
	listOption := v1.ListOptions{
		FieldSelector: "type=" + string(corev1.SecretTypeOpaque),
		LabelSelector: "kubeconfig=true",
	}
	return GetSecretList(namespace, listOption)
}

func GetKubeconfigSecret(name string, namespace string) (secret []byte, err error) {
	secretResult, err := GetSecret(name, namespace)
	if err != nil {
		return
	}
	if secretResult.Type != corev1.SecretTypeOpaque {
		err = fmt.Errorf("secret %s is not of type %s, but rather is of type %s", name, corev1.SecretTypeOpaque, secretResult.Type)
		return
	}
	if secretResult.Labels["kubeconfig"] != "true" {
		err = fmt.Errorf("secret %s does not have label kubeconfig=true", name)
		return
	}
	secret = secretResult.Data[corev1.ServiceAccountKubeconfigKey]
	return
}

func DeleteKubeconfigSecret(name string, namespace string) (err error) {
	return DeleteSecret(name, namespace)
}

func CreateKubeconfigSecret(name string, namespace string, kubeconfig []byte) (err error) {
	if DefaultConfig == nil {
		DefaultConfig, _ = GenerateKubernetesConfig()
	}
	client, err := kubernetes.NewForConfig(DefaultConfig)
	if err != nil {
		return
	}

	labelMap := make(map[string]string)
	labelMap["cmaaws"] = "true"
	labelMap["kubeconfig"] = "true"

	dataMap := make(map[string][]byte)
	dataMap[corev1.ServiceAccountKubeconfigKey] = kubeconfig

	secret := &corev1.Secret{
		ObjectMeta: v1.ObjectMeta{Name: name, Labels: labelMap},
		Type:       corev1.SecretTypeOpaque,
		Data:       dataMap,
	}

	_, err = client.CoreV1().Secrets(namespace).Create(secret)
	return
}
