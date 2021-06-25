#!/bin/bash -xe
TMP_DIR=$(mktemp -d)
kpt fn source -o unwrap ../k8s/ > ${TMP_DIR}/manifests.yaml
cp ../validate/policy.yaml ${TMP_DIR}/policy.yaml
kpt fn eval --image $1 --image-pull-policy never ${TMP_DIR}
