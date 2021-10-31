#

This document is the implementation of the [ADR Security and Compliance](../adr/0015-security-and-compliance.md).

| Description                                                      | ADR                                                 |
| -----------------------------------------------------------------| ----------------------------------------------------|
| Terraform state storage not using custom key                     |                                                     |
| Kubernetes API available online                                  | [0007](../adr/0007-management-connectivity.md)      |
| Vault API available online                                       | [0018](../adr/0018-vault-setup.md)                  |
| CI/CD permissions are overly broad                               | [0018](../adr/0018-vault-setup.md)                  |
| Azure Foundation Key Vault available online                      | [0019](../adr/0019-foundation-azure.md)             |
| The management role used in platform creation is broad           | [0019](../adr/0019-foundation-azure.md)             |
| AKS/EKS cluster available online                                 | [0023](../adr/0023-kubernetes-azure.md)             |
|                                                                  | [0024](../adr/0024-kubernetes-aws.md)               |
| Pods in AKS/EKS cluster have unrestricted access to the Internet | [0023](../adr/0023-kubernetes-azure.md)             |
|                                                                  | [0024](../adr/0024-kubernetes-aws.md)               |
| Kubernetes dashboard disabled in AKS/EKS                         | [0023](../adr/0023-kubernetes-azure.md)             |
|                                                                  | [0024](../adr/0024-kubernetes-aws.md)               |
| SSH node access is NOT allowed on AKS/EKS                        | [0023](../adr/0023-kubernetes-azure.md)             |
|                                                                  | [0024](../adr/0024-kubernetes-aws.md)               |
| Disks in AKS/EKS are encrypted with custom encryption key        | [0023](../adr/0023-kubernetes-azure.md)             |
|                                                                  | [0024](../adr/0024-kubernetes-aws.md)               |
| AKS/EKS uses default Azure Linux OS for nodes                    | [0023](../adr/0023-kubernetes-azure.md)             |
|                                                                  | [0024](../adr/0024-kubernetes-aws.md)               |
| AKS audit logs are available for 5 days                          | [0023](../adr/0023-kubernetes-azure.md)             |
|                                                                  | [0024](../adr/0024-kubernetes-aws.md)               |
| AKS/EKS instance metadata service disabled                       | [0023](../adr/0023-kubernetes-azure.md)             |
|                                                                  | [0024](../adr/0024-kubernetes-aws.md)               |
| AKS CA cert rotation is not enabled                              | [0023](../adr/0023-kubernetes-azure.md)             |
| AKS/EKS host encryption enabled                                  | [0023](../adr/0023-kubernetes-azure.md)             |
|                                                                  | [0024](../adr/0024-kubernetes-aws.md)               |
| No WAF enabled on the load balancer                              | [0025](../adr/0026-kubernetes-aws-loadbalancing.md) |
| SSL policies on load balancer set to default (permissive)        |                                                     |
