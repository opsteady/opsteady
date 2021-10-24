ARG FROM_IMAGE
FROM $FROM_IMAGE

USER root

# Used by the Opsteady cli
RUN install-packages golang-1.16 git jq

RUN getfile https://releases.hashicorp.com/terraform/1.0.9/terraform_1.0.9_linux_amd64.zip

RUN getfile https://releases.hashicorp.com/vault/1.8.2/vault_1.8.2_linux_amd64.zip

RUN getfile https://get.helm.sh/helm-v3.7.1-linux-amd64.tar.gz

RUN getfile https://storage.googleapis.com/kubernetes-release/release/v1.22.1/bin/linux/amd64/kubectl

RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64-2.0.30.zip" -o "awscliv2.zip" && \
    unzip awscliv2.zip && \
    ./aws/install

SHELL ["/bin/bash", "-o", "pipefail", "-c"]

# https://docs.microsoft.com/en-us/cli/azure/release-notes-azure-cli?tabs=azure-cli
# Ubuntu 21.04 has months old version therefore using pip to install it
RUN DEBIAN_FRONTEND=noninteractive install-packages python3-pip
RUN pip install azure-cli==2.27.2

ENV PATH="/usr/lib/go-1.16/bin:${PATH}"

ARG VAULT_CA_STORAGE_ACCOUNT
RUN if [[ ! -z "$VAULT_CA_STORAGE_ACCOUNT" ]]; then curl -o /vault-ca.pem https://$VAULT_CA_STORAGE_ACCOUNT.blob.core.windows.net/vault-ca/ca.pem; fi

RUN openssl x509 -in /vault-ca.pem -inform PEM -out /usr/local/share/ca-certificates/vault-ca.crt
RUN update-ca-certificates

RUN install-packages docker-ce-cli

RUN getfile -n hadolint https://github.com/hadolint/hadolint/releases/download/v2.7.0/hadolint-Linux-x86_64

USER 1000
