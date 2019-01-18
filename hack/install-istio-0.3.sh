#!/usr/bin/env bash

# meta
source $(dirname $0)/exports.sh

# Istio 0.3
kubectl apply -f https://github.com/knative/serving/releases/download/v0.3.0/istio-crds.yaml
kubectl apply -f https://github.com/knative/serving/releases/download/v0.3.0/istio.yaml
kubectl label namespace default istio-injection=enabled

# Validate
kubectl get pods -n istio-system
