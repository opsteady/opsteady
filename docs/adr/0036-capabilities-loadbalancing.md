# 36. Capabilities loadbalancing

Date: 2021-11-28

## Status

Status: Accepted on 2021-11-28

Foundation for [0037-capabilities-loadbalancing-implementation.md](0037-capabilities-loadbalancing-implementation.md) on 2021-11-28

## Context

Most applications will require some form of internal or external connectivity to expose their functionality. Opsteady should provide an easy path for developers to enable this connectivity.

Kubernetes provides an [Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/) abstraction to allow internal or external connectivity via some kind of loadbalancer. Work is underway to introduce the [gateway API](https://gateway-api.sigs.k8s.io/) which mitigates some of the shortcomings of the current Ingress abstraction. However, the Gateway API is still very much in the early stages of development and support for it is limited. Ingress is still the main abstraction at this point, so we will support it as the primary mechanism for exposing applications, until the Gateway API matures. Depending on the Ingress implementation, the transition to the Gateway API might be easier or harder.

There are many different implementation (or controllers) for the Ingress abstraction. Some of them are cloud native, like the [AWS Loadbalancer controller](https://github.com/kubernetes-sigs/aws-load-balancer-controller) or the [Application Gateway Ingress Controller](https://github.com/Azure/application-gateway-kubernetes-ingress), others are based on existing well-know software loadbalancer implementations like Nginx and HAProxy.

Cloud-native loadbalancers have the distinct advantage that we only have to (automatically) configure the cloud-based loadbalancers and not have to worry about operating and scaling the actual implementation. These cloud-native loadbalancers move slower in terms of feature updates and might also lag behind with the actual cluster state, making it difficult to ensure absolute zero-downtime deployments.

Running our own loadbalancer implementation gives us a large amount of flexibility. In general these solutions will be much more feature rich and more tightly integrated with the Kubernetes capabilities. The obvious downside is the burden of having the install, operate, secure and scale the loadbalancer ourselves. All external traffic directly hits the ingress loadbalancer, which makes scaling and securing the loadbalancer a primary concern. There are many choices when it comes to a community-supported loadbalancer but Nginx is the oldest and has proven to be battle-ready.

## Decision

As the Gateway API is not mature enough at this point, we will use the Nginx ingress controller as our loadbalancer implementation for Ingress resources. The Nginx controller has many useful features and is supported by the Kubernetes community. When the Gateway API matures, we will re-assess our choice and most likely switch over.

## Consequences

We will have the burden to run, operate, secure and scale the Nginx ingress loadbalancer. Developers will have full freedom to leverage the large amount of features that the Nginx ingress loadbalancer provides. When we switch over to the Gateway API we will have to make sure that the transition path is a smooth as possible, hopefully without breaking changes for our users. Given that the Kubernetes community as a whole is migrating to this target, we expect that such a smooth transition path will be possible.
