ARG FROM_IMAGE
FROM $FROM_IMAGE

USER root

# Used by the Opsteady cli
RUN install-packages git jq
RUN install-packages golang-1.19

# Can't do substituions and other things if using /bin/sh in RUN
SHELL ["/bin/bash", "-o", "pipefail", "-c"]

# renovate: datasource=github-releases depName=hashicorp/terraform versioning=semver
ENV TERRAFORM_VERSION=v1.3.6
RUN getfile https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION:1}/terraform_${TERRAFORM_VERSION:1}_linux_amd64.zip

# renovate: datasource=github-releases depName=hashicorp/vault versioning=semver
ENV VAULT_VERSION=v1.12.2
RUN getfile https://releases.hashicorp.com/vault/${VAULT_VERSION:1}/vault_${VAULT_VERSION:1}_linux_amd64.zip

# renovate: datasource=github-releases depName=helm/helm versioning=semver
ENV HELM_VERSION=v3.10.3
RUN getfile https://get.helm.sh/helm-${HELM_VERSION}-linux-amd64.tar.gz

# renovate: datasource=github-releases depName=kubernetes/kubernetes versioning=semver
ENV KUBERNETES_VERSION=v1.26.0
RUN getfile https://storage.googleapis.com/kubernetes-release/release/${KUBERNETES_VERSION}/bin/linux/amd64/kubectl

# renovate: datasource=github-tags depName=aws/aws-cli versioning=semver
ENV AWSCLI_VERSION=2.9.11
RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64-${AWSCLI_VERSION}.zip" -o "awscliv2.zip" && \
    unzip awscliv2.zip && \
    ./aws/install

# renovate: datasource=pypi depName=azure-cli versioning=semver
ENV AZURECLI_VERSION=2.43.0
# https://docs.microsoft.com/en-us/cli/azure/release-notes-azure-cli?tabs=azure-cli
# Ubuntu has old version therefore using pip to install it
RUN DEBIAN_FRONTEND=noninteractive install-packages python3-pip
RUN pip install azure-cli==${AZURECLI_VERSION}

ENV PATH="/usr/lib/go-1.19/bin:${PATH}"

ARG VAULT_CA_STORAGE_ACCOUNT
RUN if [[ ! -z "$VAULT_CA_STORAGE_ACCOUNT" ]]; then curl -o /vault-ca.pem https://$VAULT_CA_STORAGE_ACCOUNT.blob.core.windows.net/vault-ca/ca.pem; fi

RUN openssl x509 -in /vault-ca.pem -inform PEM -out /usr/local/share/ca-certificates/vault-ca.crt
RUN update-ca-certificates

RUN install-packages docker-ce-cli

# renovate: datasource=github-releases depName=hadolint/hadolint versioning=semver
ENV HADOLINT_VERSION=v2.12.0
RUN getfile -n hadolint https://github.com/hadolint/hadolint/releases/download/${HADOLINT_VERSION}/hadolint-Linux-x86_64

# renovate: datasource=github-releases depName=golangci/golangci-lint versioning=semver
ENV GOLINT_VERSION=v1.50.1
RUN getfile https://github.com/golangci/golangci-lint/releases/download/${GOLINT_VERSION}/golangci-lint-${GOLINT_VERSION:1}-linux-amd64.tar.gz
RUN install-packages libc-dev gcc

RUN install-packages yamllint

# renovate: datasource=github-tags depName=rancher/k3d versioning=semver
ENV K3D_VERSION=v5.4.6
RUN curl -s https://raw.githubusercontent.com/rancher/k3d/main/install.sh | TAG=${K3D_VERSION} bash

USER 1000
