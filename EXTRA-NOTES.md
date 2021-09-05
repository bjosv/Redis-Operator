# Extra notes for experiments

## Build containers
make container
> bjosv/operator:6.2-1
> bjosv/redisnode:6.2-1

## Starting cluster
kind create cluster --config chart/kind_multinode.yaml
or
kind create cluster --image kindest/node:v1.19.4 --config chart/kind_multinode.yaml
kind create cluster --image kindest/node:v1.20.0 --config chart/kind_multinode.yaml
kind create cluster --image kindest/node:v1.21.1 --config chart/kind_multinode.yaml

## Upload continers
kind load docker-image bjosv/operator:6.2-1
kind load docker-image bjosv/redisnode:6.2-1
minikube image load bjosv/operator:6.2-1
minikube image load bjosv/redisnode:6.2-1

## Install operator
helm install operator chart/redis-operator-new/

## Install DB
helm install cluster chart/redis-cluster-new/

k get rediscluster
kubectl get pods --sort-by=.metadata.creationTimestamp -o wide

## Upgrade dependencies
go get k8s.io/client-go@v0.22.1
go get k8s.io/code-generator@v0.22.1
go get k8s.io/apimachinery@v0.22.1
go get k8s.io/apiextensions-apiserver@v0.22.1

go get github.com/prometheus/client_golang@v1.11.0
go mod tidy

### Remove:
pkg/api/redis/install/install.go
pkg/api/install/install.go

### Generate API
go mod vendor
#### Codegen seems to require GOPATH usage:
cd $GOPATH/src/github.com/bjosv/
ln -s ~/git/Redis-Operator/ redis-operator
cd ~/git/Redis-Operator

./hack/update-codegen.sh
rm -rf vendor/
