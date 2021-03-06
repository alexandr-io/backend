#!/bin/sh

USER="$1"
BASE64_CSR=$(base64 < "$1".csr | tr -d '\n')

CLUSTER_CERTIFICATE_AUTHORITY_DATA=$(kubectl config view -o jsonpath='{.clusters[0].cluster.certificate-authority-data}' --raw)

NODE_NAME=$(kubectl get node -o jsonpath='{.items[0].metadata.name}')

export USER
export BASE64_CSR
export NODE_NAME

export CLUSTER_CERTIFICATE_AUTHORITY_DATA

envsubst < csr.yaml | kubectl apply -f -

kubectl certificate approve "$1"
CLIENT_CERTIFICATE_DATA=$(kubectl get csr "$1" -o jsonpath='{.status.certificate}')
export CLIENT_CERTIFICATE_DATA
envsubst < kubeconfig.tpl > kubeconfig."$1"

cat kubeconfig."$1"