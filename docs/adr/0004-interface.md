# 4. Interface

Date: 2021-08-25

## Status

Status: Accepted on 2021-08-25
Builds on [0003-project-scope.md](0003-project-scope.md) on 2021-08-25

## Context

What is the purpose of the API (Kubernetes API) and the portal mentioned in the project scope as part of the interface layer?

## Decision

To facilitate self-service there needs to be an API available to service the requests. Kubernetes API is very flexible and is already being used by many projects to allow users to configure or use Kubernetes in a certain way. Besides that the users are familiar with the API, it is easy for new joiners to jump on board and because the project then only has one API it makes the Kubernetes API a very well fit for Opsteady.

Besides the API a portal is a must-have as some information is better processed visualized or because not every user of the platform is a technical user. This portal is based on [Backstage](https://backstage.io/) and it uses only the Kubernetes API as its source. But it isn't a Kubernetes resource viewer it integrates the needs of different concerns and roles in an organization for their desired information and control.

Kubernetes is the source of information where everything can be found and where this information can be viewed, created, and manipulated using the Kubernetes API and the portal that uses it.

## Consequences

Although the Kubernetes API is and can be consumed as a REST API most users don't look at it that way which can confuse those that want to communicate with it without a traditional `kubectl apply`. The good thing however is that the API has clients for every language.
Backstage is still in the early stage and is getting updated fast, this can potentially be a problem if breaking changes are introduced.
