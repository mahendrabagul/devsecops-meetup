#!/bin/bash
set -e

cd certificates || exit
./gen_certs.sh
./gen_secrets.sh

cd ..

docker build -t mahendrabagul/golang-grpc-server ./node_grpc_server && docker push mahendrabagul/golang-grpc-server
docker build -t mahendrabagul/node-grpc-client ./node_grpc_client && docker push mahendrabagul/node-grpc-client
kubectl apply -f node_grpc_server/kubernetes/configmap.yaml
kubectl apply -f node_grpc_server/kubernetes/deployment.yaml
#kubectl apply -f node_grpc_server/kubernetes/ingress.yaml
kubectl apply -f node_grpc_server/kubernetes/service.yaml
kubectl wait --namespace devsecops-meetup --for=condition=ready pod --selector=app=golang-grpc-server --timeout=120s
kubectl apply -f node_grpc_client/kubernetes/deployment.yaml
