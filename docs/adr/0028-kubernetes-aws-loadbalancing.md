# 28. Kubernetes AWS loadbalancing

Date: 2021-10-27

## Status

Status: Accepted on 2021-10-27

Builds on [0024-kubernetes-aws.md](0024-kubernetes-aws.md) on 2021-10-21

## Context

A cloud loadbalancer is needed to expose Kubernetes services (such as the ingress controller) externally. Kubernetes is moving away from in-tree cloud providers and replacing them with out-of-tree controllers, maintained by the cloud providers. This allows greater flexibility and no hard dependency on the Kubernetes release schedule.

## Decision

We will use the AWS load balancer controller (ALBC) to create a cloud loadbalancer. This is the recommended project by AWS for creating load balancers. It is well-supported and feature rich.

### High-availability

We will run two replicas with a pod disruption budget that requires one replica to be available at all times. The pods will be given a priority class of 'system-cluster-critical' to ensure that the ALBC is always scheduled, potentially evicting other workloads. We set an anti-affinity to rule to make sure that the controller replicas are never colocated onto the same node.

We enable [Pod readiness gates](https://kubernetes-sigs.github.io/aws-load-balancer-controller/v2.1/deploy/pod_readiness_gate/) to indicate that pod is registered to the ALB/NLB and healthy to receive traffic. The pod readiness gate is needed under certain circumstances to achieve full zero downtime rolling deployments.

### IAM and RBAC

We will use the EKS IAM for service accounts functionality. The service account for the ALBC will receive the necessary permissions to access cloud resources to create the load balancers. Furthermore, the service account will be given sufficient Kubernetes RBAC permissions to interact with the Kubernetes API server and listen for relevant updates to ingress resources.

We limit the ALBC to only listen for ingress resources in the 'platform' namespace. This is to ensure that the platform ingress controller is the main entrypoint for all application workloads.

### Security

The pods will run as non-root with a read-only filesystem and a deny policy on privilege escalation. We will not enable the Shield and WAF functionality now. This is mainly to save costs at this point. We will re-evaluate in the future if we need to enable this. There is nothing preventing us from doing that.

We are keeping the default [ALB](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html#describe-ssl-policies) and [NLB](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/create-tls-listener.html#describe-ssl-policies) security policies for now. In the future we might restrict this further to ensure modern TLS versions and encryption cyphers.

We are creating a shared backend security group to control backend traffic. The controller will automatically create one security group: the security group will be attached to the LoadBalancer and allow access from inbound-cidrs to the listen-ports. Also, the securityGroups for Node/Pod will be modified to allow inbound traffic from this securityGroup.

### Ingress classes

We will configure an ingress class called 'alb' on which the ALBC will listen and act. We will not create a dedicated ingress class resource.

### Endpoint slices

We are enabling the endpoint slices to ensure that endpoints never grow too large. This feature is supported and stable from version 1.21 onwards.

## Consequences

We are configuring the AWS load balancer controller with default security settings and no (expensive) addons. This means that in the future we might need to tighten these controlers. There is nothing preventing us from this.
