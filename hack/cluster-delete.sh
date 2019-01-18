#!/usr/bin/env bash

# meta
source $(dirname $0)/exports.sh

gcloud container clusters delete $CLUSTER_NAME --zone $CLUSTER_ZONE
