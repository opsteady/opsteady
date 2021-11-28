# 38. Capabilities user authentication

Date: 2021-11-29

## Status

Status: Accepted on 2021-11-29

## Context

Users of the platform need to be able to use the platform with different roles and authorization levels, this ADR defines how we will support this.

## Decision

Kubernetes has multiple authentication methods but most of them are not enabled or can not be enabled when using AKS/EKS. They do offer their mechanisms for authenticating but they don't give enough flexibility of connecting the user IDP's to it. This means users can not define who has access to what.
One way around this is to use ServiceAccount tokens but that is very cumbersome and users will have tokes lying around.

The solution we are choosing is using the impersonation option in Kubernetes. This means actual roles/groups from external IDP can be coupled to roles in Kubernetes where an application authenticates and then checks if you are allowed to use the role, if so executes the desired command using that role. All of this is transparent for the user. To make this possible we will be using [pinniped concierge](https://pinniped.dev/) as VMware is behind it and there seems to be more traction than for [kube-oidc-proxy](https://github.com/jetstack/kube-oidc-proxy).

If and when cloud providers allow an easy way of connecting other IDPs we will switch to that.

Pinniped solves the authorization problem where it connects external IDP groups to Kubernetes roles. But pinniped can only have one IDP (OIDC) connection where we need one for ourselves using our Azure AD tenant and one or more for the users of the specific platform. To enable those options we have chosen [dex](https://dexidp.io/). Dex allows us to offer multiple IDPs on login so we can connect to multiple user IDPs on one cluster. Pinniped self offers supervisor but it does not allow us to add them programmatically as Dex does.

For more information on the pinniped architecture and the way the authentication works see [their architecture](https://pinniped.dev/docs/background/architecture/)

### Dex configuration

Dex depends on our load balancing, DNS, and certification capabilities to deliver the integrated OIDC from different IDPs. By default, we have our tenant in Azure AD connected to the cluster. Dex requires information to be stored, we are using the Kubernetes storage for that.
TODO how do we configure multiple IDPs for our users? (we also have a second one, the same tenant representing the customer for now)
TODO add info about using our build because of configuration, link to PR/issue

Although Dex gives us an option to style the UI the users get presented we have decided not to make any changes to that for now.

### Pinniped configuration

Pinniped does not have a Helm chart and needs to be applied using the raw Kubernetes manifest files. It is configured to use Dex as the OIDC provider. Although pinniped is exposed using our load balancing capability and uses the DNS from the platform it does not use our capability for generating certificates. It uses its own generated certificate and therefore the load balancer passes SSL/TLS through to pinniped to handle it.

## Consequences

It gives us a lot of flexibility to connect multiple user IDPs and allows them to define which group has access to which role in Kubernetes. We, however, have to run extra components for this and we need to make sure this is highly available as without it users have no access to the cluster.
