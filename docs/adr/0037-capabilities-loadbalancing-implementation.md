# 37. Capabilities loadbalancing implementation

Date: 2021-11-28

## Status

Status: Accepted on 2021-11-28

Builds on [0037-capabilities-loadbalancing.md](0037-capabilities-loadbalancing.md) on 2021-11-28

## Context

The Nginx ingress loadbalancer is exposed to the Internet on both AWS and Azure. This ADR describes how it is technically implemented.

## Decision

### Azure

The Nginx ingress loadbalancer will be exposed as a NodePort service on each of the cluster nodes, fronted by an Azure loadbalancer. Since the loadbalancer is only running on (some of) our platform nodes we will set the [externalTrafficPolicy]((https://www.asykim.com/blog/deep-dive-into-kubernetes-external-traffic-policies)) to 'local'. This ensures that the traffic is only ever forwarded from the loadbalancer to cluster nodes that are actually running a loadbalancer. This reduces the amount of network hops and also clearly separates platform nodes from customer nodes on the network level.

### AWS

In AWS we will also set the externalTrafficPolicy to 'local' but we will provision a network loadbalancer in IP target mode. This means that traffic is directly forwarded from the cloud loadbalancer to the loadbalancer pods. This is the most efficient network path that we can setup.

### Monitoring

The Nginx loadbalancer will be monitored with Prometheus and Grafana. A dashboard will be provided with the most important metrics. Our customers can use the same dashboard to inspect traffic. The Nginx loadbalancer pods will also be monitored for abnormal behaviour on the application level. We will not monitor for HTTP status codes, as that is something that needs to be monitored by our users. We will provide mechanisms to make this as easy as possible.

## Consequences

The loadbalancer implementation is not exactly the same on both clouds but this is due to underlying differences in the cloud-native loadbalancer features. From a customer perspective, however, there should be no discernable difference in behaviour.
