# 2. Monorepo

Date: 2021-08-25

## Status

Status: Accepted on 2021-08-25

## Context

We need to decide where and how to store the code and documentation for the Opsteady platform.

## Decision

We are using a monorepo to store all the infrastructure as code, code, configuration, and documentation. This makes it easier to reuse the code and gives a better and faster overview of the components and decisions made to the platform.

There can be reasons to move certain aspects out of the repository but those should be discussed and explained in a separate ADR when it happens.

## Consequences

The repository might become big and could potentially have lots of tags that would require cleanup.
