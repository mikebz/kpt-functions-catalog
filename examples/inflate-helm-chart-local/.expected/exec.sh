#!/usr/bin/env bash

kpt fn eval --image-pull-policy ifNotPresent --image gcr.io/kpt-fn/inflate-helm-chart:unstable \
--mount type=bind,src="$(pwd)",dst=/tmp/charts -- \
name=helloworld-chart \
releaseName=test \
valuesFile=/tmp/charts/helloworld-values/values.yaml