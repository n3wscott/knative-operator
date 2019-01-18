#!/usr/bin/env bash

# meta
source $(dirname $0)/exports.sh

# create project
gcloud projects create $PROJECT_ID --set-as-default