#!/bin/bash -xe

# Create a temp dir
CONFIG_DIR=$(mktemp -d)

# deletes the temp directory
function cleanup {
  rm -rf "${CONFIG_DIR}"
}
# register the cleanup function to be called on the EXIT signal
trap cleanup EXIT

# Get Frontend Configs
cd frontend
kustomize build k8s/prod > ${CONFIG_DIR}/frontend.yaml
cd ..

# Get Backend Configs
cd backend
kustomize build k8s/prod >> ${CONFIG_DIR}/backend.yaml
cd ..

# Run them against our constraints
cp validate/policy.yaml ${CONFIG_DIR}/prod.yaml

# Run the validations
kpt fn source ${CONFIG_DIR} | kpt fn eval --image gcr.io/kpt-fn/gatekeeper:v0.1 -
