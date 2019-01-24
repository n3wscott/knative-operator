#!/usr/bin/env bash

kubectl delete configmap knative-istio-crds-yaml

kubectl create configmap knative-istio-crds-yaml \
    --from-file=yaml/istio-crds-v0.3.0.yaml \
    --from-file=yaml/istio-v0.3.0.yaml


kubectl delete configmap knative-install-yaml

kubectl create configmap knative-install-yaml \
    --from-file=yaml/serving-v0.3.0.yaml \
    --from-file=yaml/build-v0.3.0.yaml \
    --from-file=yaml/eventing-v0.3.0.yaml



