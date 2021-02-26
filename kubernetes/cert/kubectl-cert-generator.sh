#!/bin/sh

USER="$1"
BASE64_CSR=$(base64 < tmp.yaml | tr -d '\n')

export USER
export BASE64_CSR

envsubst < tmp.yaml | kubectl apply -f -

kubectl certificate approve "$1"
CLIENT_CERTIFICATE_DATA=$(kubectl get csr "$1" -o jsonpath='{.status.certificate}')
export CLIENT_CERTIFICATE_DATA
envsubst < kubeconfig.tpl > kubeconfig."$1"

cat kubeconfig."$1"