# 35. Capabilities certificates

Date: 2021-11-21

## Status

Status: Accepted on 2021-11-21

Builds on [0034-capabilities-dns.md](0034-capabilities-dns.md) on 2021-11-21

## Context

Services that are going to be exposed to the outside world should have a valid certificate. Besides this use cases there are many reasons to have certificates automatically generated.

## Decision

We've decided to use the [cert-manager](http://cert-manager.io/) project for managing our certificates. This operator integrates with a lot of different certificate providers using different ways of verification. We will be using the free [Let's Encrypt](https://letsencrypt.org/) provider to generate valid certificates but the cert-manager allows us to plug others as well in the future.

Let's Encrypt is a free services and therefore has some limitations on how many certificates can be requested per hour. This is not an issue in production environment where the certificates are not changing much but in a development environment this is a problem. There for we offer Let's Encrypt real valid certificates but are also connected to their staging environment which provides a certificate in the same way but is not accepted by the rest of the world as a valid one.

We are running cert-manager as a single replica as there is no requirement to have high-availability. If it is restarted it will just pick up any pending changes. We do deploy it with the highest priority class, so that it will always find a place to run. It is also always scheduled on the platform nodes and never on the user nodes.

## Consequences

We will need to lifecycle manage the cert-manager operator but it allows us to provide proper certifications. Because Let's Encrypt has a limitation we need to make sure users are aware of that.
