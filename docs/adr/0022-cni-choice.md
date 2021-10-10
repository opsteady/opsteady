# 22. CNI choice

Date: 2021-10-10

## Status

Status: Accepted on 2021-10-10
Builds on [0009-ip-ranges.md](0009-ip-ranges.md) on 2021-10-10

## Context

Define why we are using the cloud-native CNI instead of an overlay network.

## Decision

The benefits of using the cloud-native CNI solution for Azure or AWS is that it is supported and the pods are accessible from the network, which means less complexity on the network layer. The drawbacks are the IP exhaustion as we need to be careful not to overlap when connecting with other networks or the number of Pods that can run on a node.

The benefit of network overlays like Calico or Cilium is that they remove the IP exhaustion because we can use a very big range for the pods. Besides that, they add extra features like network encryption. The drawbacks are however that there is no direct support from the cloud vendors and there is an added network layer that adds complexity.

Therefore we are using the cloud-native container network interface in the Kubernetes clusters.

## Consequences

We need to be careful with the IP ranges we use and the number of IPs left in the cluster for the pods and the nodes.
