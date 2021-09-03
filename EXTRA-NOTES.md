# Extra notes for experiments

## Build containers
make container
> bjosv/operator:6.2-1
> bjosv/redisnode:6.2-1

## Starting cluster
kind create cluster --config chart/kind_multinode.yaml
or
kind create cluster --image kindest/node:${K8S_VERSION} --config chart/kind_multinode.yaml

## Upload continers
kind load docker-image bjosv/operator:6.2-1
kind load docker-image bjosv/redisnode:6.2-1

## Install operator
helm install operator chart/redis-operator-new/

## Install DB
helm install cluster chart/redis-cluster-new/



