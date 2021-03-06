clusters:
- cluster:
    certificate-authority-data: ${CLUSTER_CERTIFICATE_AUTHORITY_DATA}
    server: https://{{ ansible_host }}:6443
  name: ${NODE_NAME}


users:
- name: ${USER}-${NODE_NAME}
  user:
    client-certificate-data: ${CLIENT_CERTIFICATE_DATA}


contexts:
- context:
    cluster: ${NODE_NAME}
    user: ${USER}-${NODE_NAME}
  name: ${NODE_NAME}