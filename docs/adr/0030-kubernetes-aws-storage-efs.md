# 30. Kubernetes AWS Storage EFS

Date: 2021-10-28

## Status

Status: Accepted on 2021-10-28

Builds on [0024-kubernetes-aws.md](0024-kubernetes-aws.md) on 2021-10-21

## Context

Not all workloads on the cluster are stateless. We need to provide storage options for workloads that require shared (networked) storage.

## Decision

To accomodate stateful workloads that require shared storage, we enable the AWS EFS CSI storage driver. This is an out-of-tree storage CSI driver that is developed and maintained by AWS. Kubernetes is moving away from most in-tree cloud providers, so we are following this by using this CSI driver.

### Filesystem Type

We configure all the EBS disks to be formatted with ext4 by default. There are options to change this but we will assess this in the future, based on user feedback.

### High Availability

We run two replicas of the EBS CSI driver and we add a topology spread constraint to make sure that these pods are never colocated on the same node. During a rollout, we never allow more than one replica to be unavailable.

### StorageClass options

We leave all the storage class options to default settings for now.

## Consequences

We will need to do lifecycle management of this component.
