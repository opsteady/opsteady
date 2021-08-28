# 11. No local tools

Date: 2021-08-28

## Status

Status: Accepted on 2021-08-28
Builds on [0002-project-guidelines.md](0002-project-guidelines.md) on 2021-08-28

## Context

Define how we will prevent "it works on my machine".

## Decision

Everyone is allowed to use tools available on their own machine. But by providing a Docker container with all the tooling required we can achieve consistency when people work in Visual Studio code, on the terminal and while running the CI/CD process.

The decision is that the "only" tool I need for the project is Docker and I should be able to run everything. But the decision is also that everything run in the container should also be runnable outside the container, as long as the local tools match with the ones defined in the container.

## Consequences

Everyone will need to have Docker running and the container needs to be kept up to date to facilitate development and not to bottleneck the project.
