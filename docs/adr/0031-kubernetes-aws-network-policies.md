# 28. Kubernetes AWS Network Policies

Date: 2021-10-31

## Status

Status: Accepted on 2021-10-31

Builds on [0024-kubernetes-aws.md](0024-kubernetes-aws.md) on 2021-10-21

## Context

Workloads on the cluster might need to be isolated on the network level. Network policies are the de-facto standard for doing this within Kubernetes. To activate network policies we need a network tool that enforces the policies. This ADR describes our choice for this tool.

## Decision

We will use the [Tigera Calico for EKS](https://aws.amazon.com/quickstart/architecture/eks-tigera-calico/) as our preferred tool. It is the recommended tool by AWS. The helm chart from EKS is not mature, so we will use the official Helm chart from Tigera. The only thing that we configure are the AWS EKS flags. This will trigger the operator to configure all the components according to EKS specifications. In terms of configuration there is not much else to configure. Under the hood, the network policies are implemented by IP tables.

## Consequences

The namespaces for the Tigera operator Helm chart are hardcoded to tigera-operator and the calico workloads will be deployed in the calico namespace. We would rather see that everything was contained in the platform namespace, but we accept these additional namespaces at this time. Using a non-official chart will not help us in the long run. We will need to do lifecycle management of this component.
