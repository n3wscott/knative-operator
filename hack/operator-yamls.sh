#!/usr/bin/env bash

kubectl create configmap knative-istio-crds-yaml \
    --from-file=yaml/istio-crds-v0.3.0.yaml \
    --from-file=yaml/istio-v0.3.0.yaml


kubectl create configmap knative-install-yaml \
    --from-file=yaml/serving-v0.3.0.yaml


