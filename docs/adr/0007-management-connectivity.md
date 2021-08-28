# 7. Management connectivity

Date: 2021-08-27

## Status

Status: Accepted on 2021-08-27
Builds on [0006-management-environment.md](0006-management-environment.md) on 2021-08-27
Builds on [0005-multi-cloud.md](0005-multi-cloud.md) on 2021-08-27

## Context

Managing user platforms requires regular lifecycle management and troubleshooting in case of (severe) issues.

We will be offering both Azure and AWS as target clouds for our platforms and we need network access to the resources that we host within those clouds. Most resources are available on a public endpoint or through the cloud API's. However, it is also possible to run cloud resources privately, meaning that network traffic will not leave the cloud provider backbone. This increases speed and security but at the cost of increased management complexity.

Providing a single 'pane of glass' for the management of private resources in both clouds is possible, but complex and costly.

If you keep the resources private you will need a way to connect to the private networks, probably with some kind of point-to-site VPN solution. This connection needs to be secured, preferably with an MFA mechanism and regular rotation of security credentials.

We've come up with three scenarios to implement our management solution for network connectivity.

### Fully transparent network layer

We can connect Azure and AWS via a site-to-site VPN. This means that it is possible to route traffic between clouds over a secure Internet connection.

Since we want to reach all platforms, which will be in separated network environments (VNET/VPC), we need some kind of routing solution that allows us to do this. Azure has [Virtual WAN](https://docs.microsoft.com/en-us/azure/virtual-wan/virtual-wan-about) (VWAN) solution and AWS has the [Transit Gateway](https://aws.amazon.com/transit-gateway/) (TGW) that allow transparent routing between site-to-site VPN, VNET's and VPC's. We could use this to create one connected environment for all the private network spaces. One of the routers (VWAN or TGW) can be used as an entry point for our VPN connection, which will allow full connectivity to all platforms from the local workstation.

Although ideal in overall management simplicity, it comes with significant downsides:

- The costs of running VWAN and TGW are significant. You pay upfront costs for having these resources and any connections attached to them. Since we expect that our management traffic will be relatively low, this would become severely unbalanced in terms of cost/benefit.
- Since all networks are connected, we will need to be very careful with overlapping network ranges in our platforms. It is not possible to reuse the same private network block for multiple platforms.
- It is quite complex to setup such a fully connected environment. We will need to monitor and deal with network connectivity on multiple layers using cloud resources. For the relatively small task of managing our platforms, this would not be appropriate at this time.

### VPN based access

To circumvent to costs and complexity of running a fully connected management network, we could opt to not have a hub-spoke model and instead connect to each platform environment when the need arises. This is a simple model that does not incur a lot of costs, since you don't pay for point-to-site VPN if you are not connected and we don't need to connect our environments to one central location. We will still be able to use private endpoints for all our services.

However, this approach also has some downsides:

- We will need additional tooling to ease the mechanism of connecting to platforms. This tooling should setup proper VPN connectivity in terms of network and credentials.
- To reduce credentials complexity we will need to use a centralized identity provider like Azure AD. However, this means that we need to use the cloud provided VPN clients, which are not supporting all operating systems (e.g. Linux is missing)

### Public endpoints with whitelisting

Instead of containing our cloud resources in a fully private environment, we could opt to make them publicly available. This significantly reduces the complexity of reaching those resources but at a cost of a degraded security boundary. Without any additional measures, we would be exposing our endpoints to DDoS attacks or brute-force hacking.

In practice, the only endpoint that we are talking about here is the Kubernetes API. Both clouds have a possibility of whitelisting IP addresses that are allowed to contact the API. This reduces the attack surface significantly and combined with the identity-based access for the API should provide enough security to get started with our platforms.

## Decision

We see multiple ways of managing our platforms from a network perspective. They are not mutually exclusive options: we could always switch to a different model and be able to manage our clusters.

We, therefore, propose to start with the most simple option by using public endpoints with whitelisting for our cloud resources. Considering that we are just starting and need to make progress in the right areas this seems like the right approach to start with.

If our platform scale increases significantly, or our security requirements increase, we can opt to switch to a more heavy-weight model where we make our cloud resources private. This will then be decided upon in a future ADR.

## Consequences

Our endpoints will be exposed to the public Internet, but we mitigate part of that threat with IP whitelisting.

We will have a very simple way of managing our platforms from local workstations and CI/CD solutions.
