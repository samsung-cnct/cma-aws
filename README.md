CMA AWS API
[![Build Status](https://jenkins.cnct.io/buildStatus/icon?job=cma-aws/master)](https://jenkins.cnct.io/job/cma-aws/job/master/)

## Overview

The cma-aws repo provides an API to provision and manage kubernetes EKS clusters in AWS.  The API supports both REST and GRPC.  It is currently used by the [cluster-manager-api](https://github.com/samsung-cnct/cluster-manager-api).

## Getting started

See [Protocol Documentation](https://github.com/samsung-cnct/cma-aws/blob/master/docs/api-generated/api.md)


### Requirements when running with [cluster-manager-api]
When using cma-aws with the cluster-manager-api, cert-manager and nginx-ingress must be deployed.
(https://github.com/samsung-cnct/cluster-manager-api)
- Kubernetes 1.7+
- [nginx-ingress](https://github.com/helm/charts/tree/master/stable/nginx-ingress)
- [cert-manager](https://github.com/helm/charts/tree/master/stable/cert-manager)

### Deploy
```bash
$ helm install deployments/helm/cma-aws --name cma-aws
```

## Contributing
Utilizes:
- [eksctl](https://github.com/weaveworks/eksctl)
- [Protocol Buffers](https://developers.google.com/protocol-buffers/)

## Developer local container build
```
cd ./build/docker/cma-aws
docker build -t <dockerhub-user/<container-name>:<tag> -f ./Dockerfile ../../..
```
## Developer local testing
```
// minikube
minikube start

make -f build/Makefile cmaaws-bin-darwin

// start it
./cma-aws --kubeconfig ~/.kube/config
```
