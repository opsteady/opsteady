# 9. IP ranges

Date: 2021-08-27

## Status

Status: Accepted on 2021-08-27

## Context

Decide on IP ranges to be used for the customers.

## Decision

There are 3 private ranges available:

- Range from 10.0.0.0 to 10.255.255.255 — a 10.0.0.0 network with a 255.0.0.0 or an /8 (8-bit) mask
- Range from 172.16.0.0 to 172.31.255.255 — a 172.16.0.0 network with a 255.240.0.0 (or a 12-bit) mask
- Range from 192.168.0.0 to 192.168.255.255 a 192.168.0.0 network masked by 255.255.0.0 or /16

The simplest solution would be to choose a /16 range in a 10.x.x.x network because it gives a lot of IP addresses. But there are some drawbacks to this approach. First, if there are other networks (on-premise) that need to be paired it might lead to a problem with overlapping ranges. Second, we might have the same problem if we wanted to connect all environments to a management network.

Choosing a very small range like a /21 with 2046 IP addresses could work but when split into multiple other ranges for pods and load balancers it becomes just too small.

We have decided to use /19 mask which gives us 2048 subnets in the 10.x.x.x IP range to use and optionally 128 subnets in the 172.16.x.x IP range and 8 subnets in the 192.168.x.x. By default, we will use the 10.x.x.x range as there is enough to choose from even if other networks have already IPs in that range, if for some reason collision is not fixable we will use one of the 172.16.x.x or the 192.168.x.x as a backup.

A /19 gives us around 8000 IP addresses to use which need to be divided into different ranges depending on the cloud:

AWS needs to be split up into subnet per Availability Zone (three by default) for Pods which need the most IP addresses, the load-balancer for external access, and the EKS itself as the best practice.

| Subnet address | Useable IPs | Used for     |
| -------------- | ----------- | ------------ |
| 10.0.0.0/21    | 2046        | Pod subnet A |
| 10.0.8.0/21    | 2046        | Pod subnet B |
| 10.0.16.0/21   | 2046        | Pod subnet C |
| 10.0.24.0/24   | 254         | LB subnet A  |
| 10.0.25.0/24   | 254         | LB subnet B  |
| 10.0.26.0/24   | 254         | LB subnet C  |
| 10.0.27.0/28   | 14          | EKS          |
| 10.0.27.16/28  | 14          | EKS          |
| 10.0.27.32/28  | 14          | EKS          |
| 10.0.27.48/28  | 14          | left         |
| 10.0.27.64/26  | 62          | left         |
| 10.0.27.128/25 | 126         | left         |
| 10.0.27.0/23   | 510         | left         |
| 10.0.27.0/23   | 510         | left         |

Azure doesn't need to be split as it uses all the IP address from one pool spread across multiple Availability zones.

| Subnet address | Useable IPs | Used for   |
| -------------- | ----------- | ---------- |
| 10.0.0.0/19    | 8190        | Pod subnet |

In AWS we still have some IP subnets left because of the split we need to do which we can use for other load balancers or to add some other new features in the future. If for some reason we hit a problem with the ranges we could extend those with the ranges left or we could add another /19 or a smaller range depending on the needs.

## Consequences

We need to keep a register of all the IP ranges used and unused and the IP ranges left inside the /19 range.
