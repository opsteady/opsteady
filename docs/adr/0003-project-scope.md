# 3. Project scope

Date: 2021-08-25

## Status

Status: Accepted on 2021-08-25

## Context

What is the scope of the Opsteady project and what is part of a container platform and who are the users?

## Decision

### Scope

Every container platform needs to have a landing zone, foundation, kubernetes, capabilities, and an interface as the five layers to be most functional.

**The landing zone** is an AWS account, Azure subscription, Google project, or any other cloud environment in which the platform is placed and which enforces organizational default compliance rules. This layer is very dependent on the way the organization has set up contracts with the cloud provider, security, and compliance measures taken that we believe that this layer does not belong in Opsteady for now.

**The foundation layer** is the basic layer for the cloud, it contains low-level infrastructure components necessary for anything else. For other layers to function properly and to be able to deliver required features this layer needs to exist and therefore is part of the Opsteady scope.

**The Kubernetes layer** is the Kubernetes installation with required options and platform worker nodes. As this is the core of a container platform it is part of the Opsteady scope.

**Capabilities layer** is the implementation of any of the required capabilities needed for a platform to be useful for an organization. For Opsteady to deliver the most optimal interface layer this layer is crucial and is, therefore, part of the scope.

**Interface layer** is everything that is user-facing. It starts with an API to the platform which is the Kubernetes API and a portal. It is also the layer that implements and adds integrations and developer features to facilitate self-service. And it is the layer that facilitates operations of the platform but also makes sure operations for developers is simple but advanced.

**Rollout** is not a layer but it is the way how the platform is rolled out to different environments and customers which is a fundamental piece that brings all the other layers to life and is, therefore, part of the Opsteady scope.

### Container platform

This list might evolve but these capabilities are considered essential or important to have in a container platform: logging, metrics, dashboards, alerting, container registry, DNS, certificates, secrets, storage, container vulnerability scanning, Kubernetes (security) best practices, networking, tracing, autoscaling, advanced deployments, etc.

### Platform users

The platform team who will create and maintain the container platforms and the developers who are delivering business value to the organization are the main focus of Opsteady. But Opsteady also acknowledges and delivers integrated functionality in the platform for other roles like SRE, testers, security officers, business, and many more in an organization.

## Consequences

The scope of Opsteady is very large but we believe that you need to deliver the most necessary capabilities and more to get the best out of an organization. These capabilities need to be very well build and most importantly integrated as a whole to deliver a seamless self-service experience. Because of that, it requires a lot of knowledge, dedication, and time for an organization to deliver the desired outcome.
