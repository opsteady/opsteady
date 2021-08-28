# 6. Management environment

Date: 2021-08-27

## Status

Status: Accepted on 2021-08-27
Foundation for [0007-management-connectivity.md](0007-management-connectivity.md) on 2021-08-27
Builds on [0005-multi-cloud.md](0005-multi-cloud.md) on 2021-08-27

## Context

We want to have one environment to do maintenance from in all other environments because it makes things much simpler. This environment can also be used to run globally available services.
Besides that, we also need an identity provider (AD) to manage access.

## Decision

We are running in multiple clouds and although we can have such an environment per cloud it makes sense to just have one because of simplicity.

We are choosing Azure over AWS purely on the fact that Azure offers us a better Active Directory with enterprise features like Azure AD Privileged Identity Management which might be interesting down the line.

The decision is to have a subscription in Azure named **management** that contains a small infrastructure based on AKS to be able to run any supporting tools needed for the management of other environments.
Next to that, we will use the Azure Active Directory as the main Opsteady management IDP.

This environment is not the same as the other environments as it is only used by Opsteady and needs to be very secure as it is the entry point to maintain other environments.

## Consequences

It simplifies the operations of the global service as they are only in one place.
We have a hard dependency on Azure even if no cluster would be created in Azure.
