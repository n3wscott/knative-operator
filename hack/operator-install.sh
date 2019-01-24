#!/usr/bin/env bash

$(dirname $0)/operator-yamls.sh

kubectl apply -f deploy/crds/

ko apply -f deploy/
