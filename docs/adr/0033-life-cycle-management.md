# 33. Life cycle management

Date: 2021-11-07

## Status

Status: Accepted on 2021-11-07

## Context

Life cycle management of a platform is very important as there are a lot of things being used that are constantly updated. Not keeping track of that and updating all the components can result in very hard or impossible upgrade paths or security risks.

## Decision

We are automating most of the life cycle management by using the [renovate bot](https://github.com/renovatebot/renovate) to automatically check for versions and update them with a Pull Request which we can review. The following is done with the renovate bot:

- Base OS image used in Docker
- Tools used by Opsteady, specified in the CI/CD Docker container
- GitHub action tools
- Terraform providers
- Helm charts
- Docker images used

EKS and AKS versions can not be checked using the renovate bot, we will do the life cycle management of this manually.

## Consequences

Everything is automated, less work for us.
