ARG FROM_IMAGE
FROM $FROM_IMAGE

USER root

# Used by the Opsteady cli
RUN install-packages golang-1.16 git jq

# renovate: datasource=github-releases depName=hashicorp/terraform versioning=semver
ENV TERRAFORM_VERSION=1.0.10
RUN getfile https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip

# renovate: datasource=github-releases depName=hashicorp/vault versioning=semver
ENV VAULT_VERSION=1.8.2
RUN getfile https://releases.hashicorp.com/vault/${VAULT_VERSION}/vault_${VAULT_VERSION}_linux_amd64.zip

# renovate: datasource=github-releases depName=helm/helm versioning=semver
ENV HELM_VERSION=v3.7.1
RUN getfile https://get.helm.sh/helm-${HELM_VERSION}-linux-amd64.tar.gz

# renovate: datasource=github-releases depName=kubernetes/kubernetes versioning=semver
ENV KUBERENTES_VERSION=v1.22.3
RUN getfile https://storage.googleapis.com/kubernetes-release/release/${KUBERENTES_VERSION}/bin/linux/amd64/kubectl

# renovate: datasource=github-tags depName=aws/aws-cli versioning=semver
ENV AWSCLI_VERSION=2.0.30
RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64-${AWSCLI_VERSION}.zip" -o "awscliv2.zip" && \
    unzip awscliv2.zip && \
    ./aws/install

SHELL ["/bin/bash", "-o", "pipefail", "-c"]

# renovate: datasource=github-releases depName=Azure/azure-cli versioning=semver
ENV AZURECLI_VERSION=2.27.2
# https://docs.microsoft.com/en-us/cli/azure/release-notes-azure-cli?tabs=azure-cli
# Ubuntu 21.04 has months old version therefore using pip to install it
RUN DEBIAN_FRONTEND=noninteractive install-packages python3-pip
RUN pip install azure-cli==${AZURECLI_VERSION}

ENV PATH="/usr/lib/go-1.16/bin:${PATH}"

ARG VAULT_CA_STORAGE_ACCOUNT
RUN if [[ ! -z "$VAULT_CA_STORAGE_ACCOUNT" ]]; then curl -o /vault-ca.pem https://$VAULT_CA_STORAGE_ACCOUNT.blob.core.windows.net/vault-ca/ca.pem; fi

RUN openssl x509 -in /vault-ca.pem -inform PEM -out /usr/local/share/ca-certificates/vault-ca.crt
RUN update-ca-certificates

RUN install-packages docker-ce-cli

# renovate: datasource=github-releases depName=hadolint/hadolint versioning=semver
ENV HADOLINT_VERSION=v2.7.0
RUN getfile -n hadolint https://github.com/hadolint/hadolint/releases/download/${}/hadolint-Linux-x86_64

USER 1000
