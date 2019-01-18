#!/usr/bin/env bash

# meta
source $(dirname $0)/exports.sh

./$(dirname $0)/gcp-new-project.sh

./$(dirname $0)/cluster-create.sh

./$(dirname $0)/install-istio-0.3.sh

./$(dirname $0)/install-knative-0.3.sh

# echo Pointing to:
# kubectl config current-context 