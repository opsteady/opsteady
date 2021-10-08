# 19. Foundation Azure

Date: 2021-10-08

## Status

Status: Accepted on 2021-10-08

## Context

Describe what the foundation is and any configuration decision made.

## Decision

The foundation layer is the first layer in creating a platform. It contains resources that are considered foundational and are not changing. Some examples are a resource group, VNET, and subnets, etc...

We are going to expose the Azure Key Vault endpoint on the Internet while we are building up the platform. Putting the endpoint on a private network complicates the CI/CD setup and overall management of the platform. We understand that this is a risk but we accept this risk for now. Once the Opsteady platform matures, and we move towards onboarding of customers, we will put the Azure Key Vault on a private network with gated access controls (proxy) in place. The exact setup for this still needs to be determined.

Currently, we are using the management service principle for CI/CD purposes to gather information and change data in the management subscription, like the DNS. This needs to be replaced with a dedicated service principle with the least privilege access necessary.

[Foundation design](../images/foundation-azure-0019.drawio.png)

## Consequences

There is some work for the future to restrict access to the Azure Key Vault and creating a dedicated least privilege access to the management when creating a platform.
