# 2. Project guidelines

Date: 2021-08-25

## Status

Status: Accepted on 2021-08-25
Foundation for [0011-no-local-tools.md](0011-no-local-tools.md) on 2021-08-28

## Context

Every project has essential tooling and guidelines on how to use, build, run or deploy the project.

## Decision

Although tools might change overtime we do have a set of tools we consider essential:

| Tool               | Description                                                                       |
| ------------------ | --------------------------------------------------------------------------------- |
| Visual Studio Code | Supporting at least one IDE helps onboard members, VSC has very useful extensions |
| diagrams.net       | One way to draw diagrams for consistency and ease of use                          |
| docker             | Still the standard for local development containers                               |
| terraform          | Multi-cloud and multi-purpose tool for configuration and deployment               |
| kubectl            | Uniform way to communicate to Kubernetes clusters                                 |
| helm               | Uniform way of deploying releases to Kubernetes                                   |
| cloud cli          | Uniform way to talk to the cloud (AWS, Azure, Google, etc..) APIs for operations  |

Because the guidelines might change more often and extensive documentation might be required on certain tools and ways of working we will document that separately from the ADR and link it directly to this ADR to acknowledge the importance of guidelines and consistency across the project.

## Consequences

Other tools might be required which are not registered in the ADR but in the documentation only.
