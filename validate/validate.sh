#!/bin/bash -xe

kpt fn source ../k8s/ | kpt fn eval --image gcr.io/kpt-fn/gatekeeper:v0.1 -
