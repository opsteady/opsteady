# 20. Foundation AWS

Date: 2021-10-08

## Status

Status: Accepted on 2021-10-08

## Context

Describe what the foundation is and any configuration decision made.

## Decision

The foundation layer is the first layer in creating a platform. It contains resources that are considered foundational and are not changing. Some examples are a VPC, hosted zone, subnets, etc.

Currently, we are using the management service principal for CI/CD purposes to gather information and change data in the management subscription, like the DNS. This needs to be replaced with a dedicated service principlal with the least privilege access.

[Foundation design](../images/foundation-aws-0021.drawio.png)

## Consequences

There is some future work to create a dedicated service principal with least-privilege access to the management environment for creating a platform.
