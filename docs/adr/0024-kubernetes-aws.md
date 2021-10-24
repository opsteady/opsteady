# 22. Kubernetes AWS

Date: 2021-10-15

## Status

Status: Accepted on 2021-10-15

Builds on [0009-ip-ranges.md](0009-ip-ranges.md) on 2021-10-15
Foundation for [0025-kubernetes-bootstrap.md](0025-kubernetes-bootstrap.md) on 2021-10-21

## Context

The Opsteady platform is built on Kubernetes and targets multiple clouds (currently AWS and Azure). The Kubernetes layer builds on the [foundation layer](./0021-foundation-aws.md), both specific to AWS. Our main decision regarding Kubernetes is to use the managed EKS service from AWS. Leveraging a managed solution like EKS gives us a lot of benefits at a relatively low cost. Running Kubernetes yourself is complicated and requires a lot of ongoing (reliability) engineering effort. Since we target a cloud environment that offers a managed Kubernetes solution, we will use it. The rest of this ADR explains all the configuration details of the EKS cluster setup.

## Decision

We use EKS, the managed Kubernetes solution from AWS. The configuration decisions will be explained in the following subsections.

The architecture diagram can be found [here](../images/kubernetes-aws-0024.drawio.png).

### Network

**Kubernetes API**: The Kubernetes API will be exposed both publicly on the Internet and privately within the VPC. The private connection allows cluster workloads to always connect to the Kubernetes API. Public access is needed to connect with the cluster from outside the VPC. We understand that this can be considered a potential security risk, however we feel that the risk is acceptable. The main risks that we identify are: DDoS on the endpoint, brute-forcing access and zero-day vulnerabilities in the Kubernetes API itself. There is a possibility to IP whitelist the API endpoint. This already mitigates any DDoS attacks (on top of AWS identifying and mitigating these attacks). For brute-forcing access, we are putting monitoring in place that will alert us on such events, so that we can take appropriate action. For zero-day vulnerabilities, we rely on the fact that EKS is a managed service, used by companies all over the world, for which AWS needs to ensure security at all times. In case of such a vulnerability, we expect AWS to patch the API within an acceptable timeframe. If this turns out not to be the case, there is always the possibility to IP whitelist or temporarily close the public endpoint altogether while we wait for mitigation by AWS.

**CNI Plugin**: we choose the native AWS VPC CNI plugin. EKS provides two supported solutions to container networking: an overlay network (kubenet) or native VPC networking. We feel that the VPC option gives us the most possibilities to use cloud-native integrations. It also makes the networking model easier to understand, as there is no additional abstraction. It will require careful planning to prevent IP range exhaustion but there are enough possibilities to overcome this (potential) problem. We fully embrace the cloud-native solution and try to minimize unnecessary abstractions.

**Pods**: all pods in the Kubernetes cluster will have unconstrained outbound Internet access. In the future, we will evaluate options to block outgoing traffic. Blocking outgoing traffic has considerable implications and it needs to be a user-friendly and cost-effective solution. This will be described in a future ADR.

**Metadata Service**: AWS provides an instance metadata service (IMDS) on a link-local address. This service allows broad inspection of environment (cloud) metadata and requesting new credentials. We feel that the potential security risks do not outweigh the benefits of leaving this available and will therefore block access to this service from the Pods.

**Pod security groups**: Although AWS EKS supports it, we will not enable pod security groups at this time. Enabling them disables network policy functionality and we will solve our network segmentation differently.

**Subnets**: As described in the [foundation AWS ADR](./0021-foundation-aws.md), we have a pod subnet, which contains the nodes and pods, and a public subnet, which contains our loadbalancers and NAT gateways. Furthermore, we have a small subnet that is used by EKS on cluster creation. The exact layout can be found in the [IP ranges ADR](./0009-ip-ranges.md).

**Network policies**: we will enable Calico to support [network policies on EKS](https://docs.aws.amazon.com/eks/latest/userguide/calico.html). Calico is the recommended CNI plugin but we will need to support and manage it ourselves.

### Addons

AWS offers managed addons for EKS. These are crucial system components that we need to run ourselves otherwise. We will install the following managed addons: kube-proxy, VPC CNI plugin and CoreDNS. This considerably lowers the maintenance burden for us. On top of this we will run these self-managed addons: [CNI metrics helper](https://docs.aws.amazon.com/eks/latest/userguide/cni-metrics-helper.html), [AWS Load Balancer Controller](https://kubernetes-sigs.github.io/aws-load-balancer-controller/v2.2/) and [Metrics Server](https://docs.aws.amazon.com/eks/latest/userguide/metrics-server.html). The setup of these addons will be described in a future ADR.

### Availability

**Uptime**: EKS comes with a 99.95% availability SLA for all clusters. There is no differentiation in this. This SLA is sufficient because the Kubernetes control plane is not in the 'hotpath' of user workloads.

**Node groups**: we will limit worker node upgrades to one at a time. We prefer safety over speed. In the future, we might consider alternatives if the speed of workers upgrades becomes unacceptable (large node groups).

### Storage

**(SSD) Disks**: we support AWS EBS volumes via the out-of-tree [EBS CSI driver](https://github.com/kubernetes-sigs/aws-ebs-csi-driver). A storage class will be available for users to provision disks on the fly. These disks are tied to a single AZ only (AWS limitation), so careful consideration is needed when using them in (auto)scaling and high-availability scenarios. Our documentation will contain recommended approaches when working with EBS disks.

**File shares**: we support AWS EFS shares via the out-of-tree [EFS CSI driver](https://github.com/kubernetes-sigs/aws-efs-csi-driver). A storage class will be available for user to provision shares on the fly.

### Monitoring

**Control plane logs**: we enable the API server, controller manager and audit logs and send them to Cloudwatch for storage. To limit costs, we maximize retention to 7 days. Depending on regulatory requirements we can always increase this retention time in the future.

**Kubernetes dashboard**: we will not enable the Kubernetes dashboard because it has a history of security related issues and we will provide an alternative solution with Grafana and Prometheus.

**VPC logs**: we are not enabling VPC logs at this time. In the future, we will decide if and when we implement this, when we know more about the platform usage.

### Identity and Access Management

**Pod Identity**: we use [IAM roles for service accounts](https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html) (IRSA) to enable fine-grained access control when accessing cloud resources from Pods. Every workload in the Kubernetes cluster should be able to get its own IAM role, with a minimal set of permissions, to operate successfully. IRSA is the native solution developed by AWS to enable this functionality. We also use this mechanism for the cluster addons that we run (AWS loadbalancing, VPC CNI, etc.).

**OIDC**: To support Pod Identity, we will associate the AWS EKS cluster with an [OIDC provider](https://docs.aws.amazon.com/eks/latest/userguide/enable-iam-roles-for-service-accounts.html).

**Cluster access**: the Vault secrets role, which creates the cluster, will have full admin capabilities (baked in to EKS). Accessing the cluster via this mechanism will be subjected to an escalation procedure. Human access to the cluster will be delegated to a user-defined identity provider. This will be a custom solution that is described in a future ADR.

### Encryption

**Secrets encryption**: we have decided to enable the KMS integrated envelope encryption for Kubernetes secrets using a customer managed key. This provides additional guarantees that the secrets cannot be read if, for some reason, they are extracted from the etcd datastore. Etcd itself is already encrypted with disk encryption by AWS as part of the EKS managed offering.

**Disks**: we enable disk encryption with a customer managed key for both the Kubernetes control plane and the worker nodes.

### Worker Nodes

**Instance type**: we will use the latest generation instance types because they offer the most flexibility in a number of areas (networking, storage, etc.)

**SSH Access**: we do not allow SSH'ing directly into the worker nodes. Instead, we leverage the `kubectl debug` command to start a Pod in the same Linux namespaces as the host OS. This allows us to inspect the OS (process, filesystem, etc.) via the Kubernetes API.

**Operating System**: we use the default AWS provided EKS AMIs, which are based on Amazon Linux 2. The EKS AMIs are continuously updated by AWS to receive the latest security patches and are guaranteed to be compatible with the EKS managed service. At this point, we feel confident that these images are sufficiently tested and secured to allow for most use-cases. In principle, we will not allow direct access to the host filesystem, or any other privileged subsystem. This severely reduces the chances that the host OS will become a liability. In the future, when the need arises to run the Opsteady platform in highly regulated environments, we will look into the possibility to harden the EKS AMI image according to relevant compliance frameworks. As such, this decision does not preclude the possibility that we will use custom EKS VM images in the future. This will also include any custom configurations for AppArmor or SELinux.

**Container runtime**: we will use containerd as our container runtime. The broader Kubernetes ecosystem is moving away from docker as the container runtime and so are we.

**Windows support**: we will not support Windows nodes at this time. In the future, we might investigate this option, based on user demand.

### Compliance

**Patching**: the EKS AMIs are continuously updated by EKS but we need to activate them in the node groups. A future ADR will describe the exact process of how this enforced. Until that time we will monitor the EKS AMIs and initiate an update manually to ensure we keep using the latest version.

**Policies**: EKS does not support a generic policy engine (like OPA, Kyverno, etc.) out of the box. Our solution for policy enforcement will be described in a future ADR.

## Consequences

A lot of things have been decided at this point, but there are still relatively large open topics that we need to address in the future. We will track these items in our issue tracker and pick them up as we go. There are also some accepted security risks that might deter people from using Opsteady at this point. Some of these issues we will definitely be improved in the future and some of them will be evaulated and decided based on user demand. Either way, the concerns are recorded in our compliance and security [documentation](../opsteady/security-and-compliance.md).
