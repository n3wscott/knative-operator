#!/usr/bin/env bash

# meta
source $(dirname $0)/exports.sh

# assumes cluster is up and istio is installed

# Install Serving
echo
echo Installing Serving
kubectl apply --filename https://github.com/knative/serving/releases/download/v0.3.0/serving.yaml

# Install Build
echo
echo Installing Build
kubectl apply --filename https://github.com/knative/build/releases/download/v0.3.0/release.yaml

# Install Eventing
echo
echo Installing Eventing
kubectl apply --filename https://github.com/knative/eventing/releases/download/v0.3.0/release.yaml

# Install Eventing Sources
echo
echo Installing Eventing Sources
kubectl apply --filename https://github.com/knative/eventing-sources/releases/download/v0.3.0/release.yaml

# Install Monitoring
echo
echo Installing Monitoring
kubectl apply --filename https://github.com/knative/serving/releases/download/v0.3.0/monitoring.yaml
