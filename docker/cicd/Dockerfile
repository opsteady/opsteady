ARG ACR_NAME
FROM $ACR_NAME.azurecr.io/base:1.0.0

USER root

# Used by the Opsteady cli
RUN install-packages golang-1.16 git

# https://kubernetes.io/releases/
RUN getfile https://releases.hashicorp.com/terraform/1.0.5/terraform_1.0.5_linux_amd64.zip

# https://github.com/helm/helm/releases
RUN getfile https://get.helm.sh/helm-v3.6.3-linux-amd64.tar.gz

# https://github.com/aws/aws-cli/blob/v2/CHANGELOG.rst
RUN getfile https://storage.googleapis.com/kubernetes-release/release/v1.22.1/bin/linux/amd64/kubectl

SHELL ["/bin/bash", "-o", "pipefail", "-c"]

# https://docs.microsoft.com/en-us/cli/azure/release-notes-azure-cli?tabs=azure-cli
# Ubuntu 21.04 has months old version therefore using pip to install it
RUN DEBIAN_FRONTEND=noninteractive install-packages python3-pip
RUN pip install azure-cli==2.27.2

ENV PATH="/usr/lib/go-1.16/bin:${PATH}"

USER 1000
