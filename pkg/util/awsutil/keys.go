package awsutil

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/ec2"
	"gitlab.com/mvenezia/cma-aws/pkg/util/awsutil/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
)

// This function will return back a list of key pairs
func GetKeyList() (keyPairs []*ec2.KeyPairInfo, err error) {
	if EC2Service == nil {
		err = initializeAWS()
		if err != nil {
			return
		}
	}
	keyPairList, err := EC2Service.DescribeKeyPairs(nil)
	if err != nil {
		return
	}
	keyPairs = keyPairList.KeyPairs
	return
}

// This function will create an AWS Key.
// TODO This presently does not support the idea of using a preexisting key from the user, perhaps it should support that
func CreateKey(name string, credentials awsmodels.Credentials) (privateKey string, err error) {
	service, err := createEC2ServiceFromCredentials(credentials)
	if err != nil {
		return
	}

	// Create the key
	result, err := service.CreateKeyPair(&ec2.CreateKeyPairInput{
		KeyName: aws.String(name),
	})
	// Verify things went OK
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == "InvalidKeyPair.Duplicate" {
			err = fmt.Errorf("keypair %q already exists", name)
			return
		}
		err = fmt.Errorf("unable to create key pair: %s, %v", name, err)
		return
	}

	// Return back the private key
	privateKey = *result.KeyMaterial
	return
}

func DeleteKey(name string, credentials awsmodels.Credentials) (err error) {
	service, err := createEC2ServiceFromCredentials(credentials)
	if err != nil {
		return
	}

	_, err = service.DeleteKeyPair(&ec2.DeleteKeyPairInput{
		KeyName: aws.String(name),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == "InvalidKeyPair.Duplicate" {
			err = fmt.Errorf("keypair %q does not exist", name)
			return
		}
		err = fmt.Errorf("unable to create key pair: %s, %v", name, err)
		return
	}

	return
}
