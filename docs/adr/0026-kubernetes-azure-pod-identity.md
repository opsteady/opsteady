# 26. Kubernetes Azure pod identity

Date: 2021-10-28

## Status

Status: Accepted on 2021-10-28

Builds on [0023-kubernetes-azure.md](0023-kubernetes-azure.md) on 2021-10-28

## Context

Azure Kubernetes cluster needs to allow Pods to use Azure resources without providing Service Principle credentials.

## Decision

We are using [AAD pod identity](https://github.com/Azure/aad-pod-identity), see the [ADR 0023-kubernetes-azure.md](0023-kubernetes-azure.md) for why.

### Optimization

It is possible to [fine-tune](https://azure.github.io/aad-pod-identity/docs/configure/feature_flags/) how the MIC and the NMI components run but the defaults should be sufficient and as long as there is no reason to change we will leave them as is.

### Binding scope

By default the AzureIdentityBinding looks across namespaces but we don't want that because in the future we might want multi-tenancy and this breaks that. Therefor we are [only allowing binding in the same namespace](https://azure.github.io/aad-pod-identity/docs/configure/match_pods_in_namespace/).

## Consequences

Because we are not using the build in pod identity we have one more component to run.
