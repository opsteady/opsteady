# 27. Azure storage

Date: 2021-10-29

## Status

Status: Accepted on 2021-10-29

Builds on [0023-kubernetes-azure.md](0023-kubernetes-azure.md) on 2021-10-29

## Context

We and our users will need storage (volumes).

## Decision

Azure delivers [default csi storage driver](https://docs.microsoft.com/en-us/azure/aks/csi-storage-drivers) which allows usage of Azure Disks (local volumes) and Azure Files (network share). This is more than sufficient for now which is why we will use it.

## Consequences

No component to run our self, there could be some issues with multi-tenancy and Azure files in the future.
