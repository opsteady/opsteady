# 33. DNS Capability

Date: 2021-11-10

## Status

Status: Accepted on 2021-11-10

Foundation for [0035-capabilities-certificates.md](0035-capabilities-certificates.md) on 2021-11-21

## Context

Some platform services need to be exposed outside the platform. DNS is the Internet standard for resolving services from a human-friendly name to an IP address. We want to offer automated DNS management for the platform or user services that we run in the clusters. If desired, a DNS name should be automatically created and available for the services that need it.

## Decision

We've decided to use the [external-dns](https://github.com/kubernetes-sigs/external-dns) project for managing our DNS records. This operator integrates with a lot of different DNS management providers. We will be using the Azure and AWS integrations, but in the future we might enable other providers if requested by our platform users.

With appropriate cloud credentials (via managed identities), external-dns will update DNS records in our cloud-hosted zones. DNS names can be configured through specific annotation on either a service or ingress resource.

We are running external-dns as a single replica as there is no requirement to have high-availability. If it is restarted it will just pick up any pending changes. We do deploy it with the highest priority class, so that it will always find a place to run. It is also always scheduled on the platform nodes and never on the user nodes.

## Consequences

We will need to lifecycle manage the external-dns operator. However, the provided functionality is relatively stable so we expect the maintenance burden to be relatively low. We've only enabled the AWS and Azure integrations for now. Any additional DNS providers will need to be investigated and configured.
