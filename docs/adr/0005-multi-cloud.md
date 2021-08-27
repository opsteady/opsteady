# 5. Multi-cloud

Date: 2021-08-26

## Status

Status: Accepted on 2021-08-26
Foundation for [0006-management-environment.md](0006-management-environment.md) on 2021-08-27
Foundation for [0007-management-connectivity.md](0007-management-connectivity.md) on 2021-08-27

## Context

Decide if the platform should be available on multiple clouds.

## Decision

Opsteady should create and simplify the maintenance of hundreds of clusters across multiple clouds. To do that the project will use different cloud-native solutions. Starting with just one cloud might steer the project in the wrong direction which would not allow it to run on another cloud later in time. But starting with too many cloud provides at the same time slows down the project implementation. Therefore we believe that starting with AWS and Azure makes the most sense as they are the two biggest providers in Europe and the USA.

The future goal of the project is to support Google cloud, Alibaba cloud, and others.

## Consequences

Building a solution for two clouds requires more effort but also by just choosing two we might overlook something or miss opportunities.
