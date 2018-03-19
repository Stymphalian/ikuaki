#!/usr/bin/env bash
cd api/lobby/main
go main.go --port=42657

npm install -g grpcc
npm uninstall -g grpcc
grpcc --proto=./ikuaki.proto --address=127.0.0.1:42657 -i

glide init
glide update --strip-vendor


# protoc \
#   --proto_path=$GOPATH/src/google.golang.org/grpc/reflection/ \
#   --proto_path=$GOPATH/src/github.com \
#   --go_out=plugins=grpc,\
# Mgrpc_reflection_v1alpha/reflection.proto=google.golang.org/grpc/reflection/grpc_reflection_v1alpha,\
# MStymphalian.ikuaki.api.protosfat=github.com/Stymphalian/ikuaki/api/protosfat\
# :$GOPATH/src \
# Stymphalian/ikuaki/api/protos/ikuaki.proto

# protoc \
#   --proto_path=$GOPATH/src/github.com \
#   --go_out=plugins=grpc:$GOPATH/src \
#   Stymphalian/ikuaki/api/protosfat/fattyghost.proto

TODO:
Setup a DB
  DAO for world, agent
ML for Agents
Tests


protoc \
  --proto_path=$GOPATH/src/google.golang.org/grpc/reflection/ \
  --proto_path=$GOPATH/src \
  --go_out=plugins=grpc,\
Mgrpc_reflection_v1alpha/reflection.proto=google.golang.org/grpc/reflection/grpc_reflection_v1alpha,\
:$GOPATH/src \
  github.com/Stymphalian/ikuaki/api/protos/ikuaki.proto

protoc \
  --proto_path=$GOPATH/src \
  --go_out=plugins=grpc:$GOPATH/src \
  github.com/Stymphalian/ikuaki/api/protosfat/fattyghost.proto



# Create all the stuff
go install github.com/Stymphalian/ikuaki/tools/grpcc

docker network create -d bridge ikuaki-network
docker volume create ikuaki-redis-volume

pushd $GOPATH/src/github.com/Stymphalian/ikuaki/bin/
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ikuaki-agent github.com/Stymphalian/ikuaki/api/agent/main
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ikuaki-world github.com/Stymphalian/ikuaki/api/world/main
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ikuaki-lobby github.com/Stymphalian/ikuaki/api/lobby/main
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ikuaki-athrun github.com/Stymphalian/ikuaki/athrun
popd

# Build all the images
pushd $GOPATH/src/github.com/Stymphalian/ikuaki/
docker build -q -t=stymphalian/ikuaki-world:v1 -f=data/ikuaki-world.dockerfile .
docker build -q -t=stymphalian/ikuaki-agent:v1 -f=data/ikuaki-agent.dockerfile .
docker build -q -t=stymphalian/ikuaki-lobby:v1 -f=data/ikuaki-lobby.dockerfile .
docker rmi $(docker images --filter "dangling=true" -q --no-trunc)
popd
docker image prune --force

docker run --name ikuaki-world -p 11111:8080 --rm -it --network=ikuaki-network stymphalian/ikuaki-world:latest
docker run --name ikuaki-agent -p 11112:8081 --rm -it --network=ikuaki-network stymphalian/ikuaki-agent:latest
docker run --name ikuaki-lobby -p 11113:8083 --rm -it --network=ikuaki-network stymphalian/ikuaki-lobby:latest


## DOCKER COMPOSE
## -----------------------------------------------------------------------------
## -----------------------------------------------------------------------------
for TF in $(docker images | grep 'stymphalian' | awk '{print $1}'); do 
  docker save $TF:latest | docker-machine ssh ikuaki-master "docker load"; 
  docker save $TF:latest | docker-machine ssh ikuaki-worker-1 "docker load"; 
  docker save $TF:latest | docker-machine ssh ikuaki-worker-2 "docker load";
done

docker stack deploy -c data/docker-compose.yaml ikuaki
docker stack rm ikuaki
docker stack ps ikuaki

## USING DOCKER SWARM (with VMs)
## START Docker swarm commands
## -----------------------------------------------------------------------------
docker-machine create --driver virtualbox ikuaki-master
docker-machine create --driver virtualbox ikuaki-worker-1
docker-machine create --driver virtualbox ikuaki-worker-2

docker-machine ls

MASTER_IP=$(docker-machine inspect ikuaki-master --format='{{.Driver.IPAddress}}')
docker-machine ssh ikuaki-master "docker swarm init --advertise-addr=${MASTER_IP}:2376"

MASTER_JOIN_TOKEN=$(docker-machine ssh ikuaki-master "docker swarm join-token -q worker")
docker-machine ssh ikuaki-worker-1 "docker swarm join --token ${MASTER_JOIN_TOKEN} ${MASTER_IP}:2377"
docker-machine ssh ikuaki-worker-2 "docker swarm join --token ${MASTER_JOIN_TOKEN} ${MASTER_IP}:2377"

eval $(docker-machine env ikuaki-master)
eval $(docker-machine env -u)

# do this in master in order to see node which are apart of the swarm
docker node ls 

docker-machine ssh ikuaki-worker-1 "docker swarm leave"
docker-machine ssh ikuaki-worker-2 "docker swarm leave"
docker-machine ssh ikuaki-master "docker swarm leave --force"

docker-machine stop ikuaki-worker-1 &
docker-machine stop ikuaki-worker-2 &
docker-machine stop ikuaki-master &

docker-machine rm ikuaki-worker-1
docker-machine rm ikuaki-worker-2
docker-machine rm ikuaki-master


## KUBERNETES
## -----------------------------------------------------------------------------
## -----------------------------------------------------------------------------
minikube start 
# Do your image builds within minikube so that it can up your pods
eval $(minikube docker-env)
eval $(minikube docker-env -u)
minikube ssh
minikube ip
minikube service ikuaki-world

kubectl create -f world_ds.yaml
kubectl create -f agent_ds.yaml
kubectl delete services,deploy ikuaki-world
kubectl delete services,deploy ikuaki-agent

# creating secerts
kubectl create secret generic ikuaki-secret \
  --from-file=data/api_key \
  --from-file=data/api_key2
kubectl get secrete ikuaki-secret -o yaml

# creating config maps
kubectl create configmap ikuaki-config \
  --from-file=data/undine.properties
kubectl create configmap ikuaki-config \
  --from-env-file=data/rainbow.properties

# creating volumes
kubectl create -f redis.pv.yaml
kubectl create -f redis.pvc.yaml
kubectl delete pvc ikuaki-redis-pvc
kubectl delete pv ikuaki-redis-pv

# Redis
kubectl create -f db_ds.yaml
kubectl delete services,deploy ikuaki-redis

kubectl create -f db.yaml
kubectl expose redis --port=6379 --target_port=6379 --type=NodePort
kubectl delete pods redis

kubectl get service $SERVICE --output='jsonpath="{.spec.ports[0].nodePort}"'

# List only the names
kubectl get pods -o name
kubectl get pods -o=jsonpath='{.items[0].metadata.name}'

# bash into pod
kubectl exec -it ikuaki-redis-5747b966d9-smllt -- /bin/bash

# proxy and API
kubectl proxy --port=8080 &
curl http://localhost:8080/api/


pushd $GOPATH/src/github.com/Stymphalian/ikuaki/
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/ikuaki-world github.com/Stymphalian/ikuaki/api/world/main
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/ikuaki-lobby github.com/Stymphalian/ikuaki/api/lobby/main
docker build -q -t=stymphalian/ikuaki-lobby:v1 -f=data/ikuaki-lobby.dockerfile .
docker build -q -t=stymphalian/ikuaki-world:v1 -f=data/ikuaki-world.dockerfile .
docker rmi $(docker images --filter "dangling=true" -q --no-trunc)
popd

kubectl delete services,deploy ikuaki-world
kubectl create -f world_d.yaml
