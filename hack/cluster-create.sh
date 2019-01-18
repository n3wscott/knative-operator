#!/usr/bin/env bash

# meta
source $(dirname $0)/exports.sh

# Domain
#export CLUSTER_DOMAIN="knative.tech"
# gcloud beta compute addresses create "${CLUSTER_NAME}-ip" --region=$CLUSTER_REGION
# gcloud beta compute addresses list
#export STATIC_IP="35.247.117.150"
#export KNATIVE_GATEWAY="istio-ingressgateway"

# Certs
# certbot certonly --manual --preferred-challenges dns -d '*.default.knative.tech'
#export TLS_CERT_PATH="/Users/mchmarny/.gcp-keys/demo.knative.tech/ca.pem"
#export TLS_KEY_PATH="/Users/mchmarny/.gcp-keys/demo.knative.tech/pk.pem"

# API
gcloud services enable \
  cloudapis.googleapis.com \
  container.googleapis.com \
  containerregistry.googleapis.com

# Cluster
# gcloud container clusters delete $CLUSTER_NAME
gcloud container clusters create $CLUSTER_NAME \
  --zone=$CLUSTER_ZONE \
  --cluster-version=latest \
  --machine-type=n1-standard-4 \
  --enable-autoscaling --min-nodes=1 --max-nodes=$MAX_CLUSTER_NODE_SIZE \
  --enable-autorepair \
  --scopes=cloud-platform,service-control,service-management,compute-rw,storage-ro,logging-write,monitoring-write,pubsub,datastore \
  --num-nodes=$START_CLUSTER_NODE_SIZE

# Binding
kubectl create clusterrolebinding cluster-admin-binding \
--clusterrole=cluster-admin \
--user=$(gcloud config get-value core/account)
