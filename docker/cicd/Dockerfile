ARG ACR_NAME
FROM $ACR_NAME.azurecr.io/base:1.0.0

USER root

# Used by the Opsteady cli
RUN install-packages golang-1.16 git jq

RUN getfile https://releases.hashicorp.com/terraform/1.0.8/terraform_1.0.8_linux_amd64.zip

RUN getfile https://releases.hashicorp.com/vault/1.8.2/vault_1.8.2_linux_amd64.zip

RUN getfile https://get.helm.sh/helm-v3.6.3-linux-amd64.tar.gz

RUN getfile https://storage.googleapis.com/kubernetes-release/release/v1.22.1/bin/linux/amd64/kubectl

SHELL ["/bin/bash", "-o", "pipefail", "-c"]

# https://docs.microsoft.com/en-us/cli/azure/release-notes-azure-cli?tabs=azure-cli
# Ubuntu 21.04 has months old version therefore using pip to install it
RUN DEBIAN_FRONTEND=noninteractive install-packages python3-pip
RUN pip install azure-cli==2.27.2

ENV PATH="/usr/lib/go-1.16/bin:${PATH}"

ARG VAULT_CA_STORAGE_ACCOUNT
RUN if [[ ! -z "$VAULT_CA_STORAGE_ACCOUNT" ]]; then curl -o /vault-ca.pem https://$VAULT_CA_STORAGE_ACCOUNT.blob.core.windows.net/vault-ca/ca.pem; fi

ENV VAULT_CACERT=/vault-ca.pem

USER 1000
