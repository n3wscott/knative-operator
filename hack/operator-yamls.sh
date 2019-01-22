#!/usr/bin/env bash

kubectl create configmap knative-istio-crds-yaml \
    --from-file=yaml/istio-crds-v0.3.0.yaml \
    --from-file=yaml/istio-v0.3.0.yaml

