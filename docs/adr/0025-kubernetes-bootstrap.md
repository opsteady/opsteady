# 25. Kubernetes bootstrap

Date: 2021-10-21

## Status

Status: Proposed on 2021-10-21

Builds on [0014-roles-responsibilities.md](0014-roles-responsibilities.md) on 2021-10-15
Builds on [0023-kubernetes-azure.md](0023-kubernetes-azure.md) on 2021-10-21
Builds on [0024-kubernetes-aws.md](0024-kubernetes-aws.md) on 2021-10-21

## Context

A cluster bootstrap is needed after creation of the cluster. This puts the cluster in an initial usable state. This ADR describes the resources that are configured in the bootstrap.

## Decision

### Namespaces

We create two namespaces: platform and management.

The _platform_ namespace contains all the (open-source) cluster software that we need to provide a full platform experience. Only Opsteady personnel is allowed access to this namespace.
The _management_ namespace contains all the self-service resources (nodepools, tenants, etc.) created by platform users.

### Secrets

We create a 'docker' secret for image registry access in the management subscription. This registry contains all our cluster software images.

### RBAC

We create cluster roles for our default platform roles (admin, operator and viewer). The _platform-admin_ role gets full cluster wide permission. The _platform-operator_ role gets full access within the platform namespace. The _platform-viewer_ role gets read-only access to the platform namespace.

### Priority Classes

We configure the following priority classes in the cluster: default (10000), medium (20000), high (3000) and platform (1000000).

Default is the default priority class that is assigned to deployments. Medium and high classes can be used to indicate more importance to the deployment. The cluster priority class is reserved for Opsteady platform deployments.

## Consequences

The current platform roles are not finalized yet, so the RBAC should be considered a starting point. This will most likely be updated in the future with another ADR.
