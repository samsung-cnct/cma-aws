#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail
set -o xtrace

function finish {
  kind delete cluster
}
trap finish EXIT

echo "The current directory is $(pwd)"

kind create cluster --wait=10m --loglevel=debug

export KUBECONFIG=$(kind get kubeconfig-path)

kubectl create clusterrolebinding superpowers --clusterrole=cluster-admin --user=system:serviceaccount:kube-system:default

echo "The kind cluster default current pods are:"
kubectl get pods --all-namespaces

helm init -c
export HELM_HOST=localhost:44134
echo $HELM_HOST
# expects tillerless helm, since HELM_HOST is defined
helm plugin install https://github.com/rimusz/helm-tiller || true
helm tiller start-ci

# install cert-manager (required by cma-aws chart)
# Note: removed --wait, it times out downloading the .tgz file
kubectl apply \
    -f https://raw.githubusercontent.com/jetstack/cert-manager/release-0.6/deploy/manifests/00-crds.yaml
kubectl create namespace cert-manager
kubectl label namespace cert-manager certmanager.k8s.io/disable-validation=true
helm repo update
helm install --name cert-manager --namespace cert-manager stable/cert-manager --debug
sleep 30
helm install --name nginx-ingress stable/nginx-ingress --debug

helm repo add cnct https://charts.cnct.io
echo "build tag for cma-aws is: ${PIPELINE_DOCKER_TAG}"
helm install --name cma-aws --set "image.repo=quay.io/samsung_cnct/cma-aws:${PIPELINE_DOCKER_TAG}" cnct/cma-aws --debug
helm install -f test/e2e/cma-values.yaml --name cluster-manager-api cnct/cluster-manager-api --debug
helm install -f test/e2e/cma-operator-values.yaml --name cma-operator cnct/cma-operator --debug

sleep 120

echo "After installing the cma, current pods are:"
kubectl get pods --all-namespaces

echo "services are:"
kubectl get services

helm tiller stop

# copy test scripts to kind container
docker cp test/e2e/ kind-control-plane:/root/
# create kubernetes job to run tests
apk add gettext
envsubst < test/e2e/run-tests-job.yaml | kubectl apply -f -

# wait for tests to complete TODO: adjust timeout as necessary
kubectl wait --for=condition=complete job/cma-aws-e2e-tests --timeout=36m

# output logs after job completes
kubectl logs job/cma-aws-e2e-tests -n pipeline-tools

#logsoutput=$(kubectl logs job/cma-aws-e2e-tests -n pipeline-tools)
#if echo $nodes | grep -o "full-test PASSED"; then
#  echo "full e2e test PASSED"
#else
#  echo "full e2e test FAILED"
#  return 1
#fi
