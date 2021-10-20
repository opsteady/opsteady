# 23. Kubernetes Azure

Date: 2021-10-15

## Status

Status: Accepted on 2021-10-15

Builds on [0009-ip-ranges.md](0009-ip-ranges.md) on 2021-10-15
Foundation for [0025-kubernetes-bootstrap.md](0025-kubernetes-bootstrap.md) on 2021-10-21

## Context

The Opsteady platform is built on Kubernetes and targets multiple clouds (currently AWS and Azure). The Kubernetes layer builds on the [foundation layer](./0021-foundation-azure.md), both specific to Azure. Our main decision regarding Kubernetes is to use the managed AKS service from Azure. Leveraging a managed solution like AKS gives us a lot of benefits at a relatively low cost. Running Kubernetes yourself is complicated and requires a lot of ongoing (reliability) engineering effort. Since we target a cloud environment that offers a managed Kubernetes solution, we will use it. The rest of this ADR explains all the configuration details of the AKS cluster setup.

## Decision

We use AKS, the managed Kubernetes solution from Azure. The configuration decisions will be explained in the following subsections.

The architecture diagram can be found [here](../images/kubernetes-azure-0023.drawio.png).

From the diagram you can see that we create two node pools. AKS mandates that a node pool exists when the cluster is created. Terraform implements this by specifing a node pool in the AKS cluster resource. This is unfortunate because any changes to this node pool will result in cluster recreate. To avoid this, we create a very small system node pool and create a separate platform node pool where most workloads will be running. This gives us more flexibility when we need to update settings on the node pool.

### Network

**Kubernetes API**: The Kubernetes API will be exposed publicly on the Internet. Public access is needed to connect with the cluster from outside the VNET. We understand that this can be considered a potential security risk, however we feel that the risk is acceptable. The main risks that we identify are: DDoS on the endpoint, brute-forcing access and zero-day vulnerabilities in the Kubernetes API itself. There is a possibility to IP whitelist the API endpoint. This already mitigates any DDoS attacks (on top of AWS identifying and mitigating these attacks). For brute-forcing access, we are putting monitoring in place that will alert us on such events, so that we can take appropriate action. For zero-day vulnerabilities, we rely on the fact that AKS is a managed service, used by companies all over the world, for which Azure needs to ensure security at all times. In case of such a vulnerability, we expect Azure to patch the API within an acceptable timeframe. If this turns out not to be the case, there is always the possibility to IP whitelist or temporarily close the public endpoint altogether while we wait for mitigation by Azure.

**CNI Plugin**: we choose the native Azure CNI plugin. AKS provides two supported solutions to container networking: an overlay network (kubenet) or native networking. We feel that the native networking option gives us the most possibilities to use cloud-native integrations. It also makes the networking model easier to understand, as there is no additional abstraction. It will require careful planning to prevent IP range exhaustion but there are enough possibilities to overcome this (potential) problem.

**Pods**: all pods in the Kubernetes cluster will have unconstrained outbound Internet access. In the future, we will evaluate options to block outgoing traffic. Blocking outgoing traffic has considerable implications and it needs to be a user-friendly and cost-effective solution. This will be described in a future ADR.

**Metadata Service**: AKS provides an instance metadata service (IMDS) on a link-local address. This service allows broad inspection of environment (cloud) metadata and requesting new credentials. We feel that the potential security risks do not outweigh the benefits of leaving this available and will therefore block access to this service from the Pods.

**Subnets**: As described in the [foundation Azure ADR](./0021-foundation-azure.md), we have a pod subnet, which contains the nodes and pods, and a public subnet, which contains our loadbalancers. The exact layout can be found in the [IP ranges ADR](./0009-ip-ranges.md).

**Azure Resources**: we will be using public endpoints for all Azure resources. For now, we accept the risk that this is less secure than using private endpoints. In the future, we will reconsider of (some) Azure resources need to be reachable on a private network only.

**Network policies**: AKS allows two different network policy options: [native one supported by AKS and Calico supported by the community](https://docs.microsoft.com/en-us/azure/aks/use-network-policies). We see no additional benefits in running Calico at this time. It provides some extra features but we won't be using them. Instead, we'll use the Azure supported solution.

### Addons

AKS provides kube-proxy, CoreDNS and the Azure CNI addons out of the box. We use these as-is, with no modifications (e.g. [CoreDNS](https://docs.microsoft.com/en-us/azure/aks/coredns-custom).

### Availability

**Uptime**: AKS comes with a 99.5% availability for all free clusters. This SLA is sufficient because the Kubernetes control plane is not in the 'hotpath' of user workloads. In future, we might enable the 'paid' tier to ramp up availability to 99.95%.

**Node groups**: we will limit [node pool upgrades](https://docs.microsoft.com/en-us/azure/aks/upgrade-cluster#customize-node-surge-upgrade) to one at a time. We prefer safety over speed. In the future, we might consider alternatives if the speed of workers upgrades becomes unacceptable (large node pools).

**Maintenance window**: AKS has an option to specify a [maintenance window](https://docs.microsoft.com/en-us/azure/aks/planned-maintenance) but it isn't guaranteed Azure is going to respect it, so we are not setting it.

### Storage

**(SSD) Disks**: we support Azure disks via the out-of-tree [Azure Disk CSI driver](https://github.com/kubernetes-sigs/azuredisk-csi-driver). A storage class will be available for users to provision disks on the fly. These disks are tied to a single AZ only (Azure limitation), so careful consideration is needed when using them in (auto)scaling and high-availability scenarios. Our documentation will contain recommended approaches when working with Azure disks.

**File shares**: we support Azure Files via the out-of-tree [Azure Files CSI driver](https://github.com/kubernetes-sigs/azurefile-csi-driver). A storage class will be available for user to provision shares on the fly.

### Monitoring

**Control plane logs**: we enable the API server, controller manager and audit logs and send them to a log analytics workspace. To limit costs, we maximize retention to 5 days. Depending on regulatory requirements we can always increase this retention time in the future.

**Kubernetes dashboard**: we will not enable the Kubernetes dashboard because it has a history of security related issues and we will provide an alternative solution with Grafana and Prometheus.

**VNet logs**: we are not enabling VNet logs at this time. In the future, we will decide if and when we implement this, when we know more about the platform usage.

### Identity and Access Management

**Pod Identity**: we will enable [Pod identity](https://github.com/Azure/aad-pod-identity) as a self-managed add-on. Azure currently offers this as a managed addon but it is already deprecated in favor of V2. Once that lands, we will switch over. Pod identity is the only way for pods to get access to Azure resources, as we will not support node roles.

**Cluster access**: we are using our own mechanisms for accessing the platform, that is why we will not be enabling the [AKS-managed Azure Active Directory integration](https://docs.microsoft.com/en-us/azure/aks/managed-aad). Our solution will allow access to the cluster via any user-provided identity provider. The exact details will be provided in a future ADR.

**Cluster identity**: we are using [managed identities](https://docs.microsoft.com/en-us/azure/aks/use-managed-identity) so that we don't have to handle [expiring service principals](https://docs.microsoft.com/en-us/azure/aks/update-credentials).

### Encryption

**AKS Disks**: we enable [disk encryption](https://docs.microsoft.com/en-us/azure/aks/azure-disk-customer-managed-keys) with a customer managed key for both the OS and data disks of our AKS cluster. We also enable [host encryption](https://docs.microsoft.com/en-us/azure/aks/enable-host-encryption) for the worker nodes, although we are not storing anything secret on the hosts and volumes are already encrypted.

### Worker Nodes

**Instance type**: we will use the latest generation instance types because they offer the most flexibility in a number of areas (networking, storage, etc.). Also, we will not use the [vm tuning](https://docs.microsoft.com/en-us/azure/aks/custom-node-configuration) capabilities at this time.

**SSH Access**: we do not allow SSH'ing directly into the worker nodes. Instead, we leverage the `kubectl debug` command to start a Pod in the same Linux namespaces as the host OS. This allows us to inspect the OS (process, filesystem, etc.) via the Kubernetes API.

**Operating System**: we use the default Azure provided AKS images. The Azure images are continuously updated by Azure to receive the latest security patches and are guaranteed to be compatible with the AKS managed service. At this point, we feel confident that these images are sufficiently tested and secured to allow for most use-cases. In principle, we will not allow direct access to the host filesystem, or any other privileged subsystem. This severely reduces the chances that the host OS will become a liability. In the future, when the need arises to run the Opsteady platform in highly regulated environments, we will look into the possibility to harden the Azure AKS image according to relevant compliance frameworks. As such, this decision does not preclude the possibility that we will use custom AKS VM images in the future. This will also include any custom configurations for AppArmor or SELinux.

**Container runtime**: we will use containerd as our container runtime. The broader Kubernetes ecosystem is moving away from docker as the container runtime and so are we. [AKS 1.19](https://docs.microsoft.com/en-us/azure/aks/cluster-configuration#container-runtime-configuration) already uses containerd as the container runtime.

**Windows support**: we will not support Windows nodes at this time. In the future, we might investigate this option, based on user demand.

**Virtual nodes and Kubelet**: there is an option to use [Virtual nodes](https://docs.microsoft.com/en-us/azure/aks/virtual-nodes) and [Virtual kubelet](https://github.com/virtual-kubelet/virtual-kubelet) which aren't bringing any added value at this time, so we are not using them.

**Proximity placement**: It is possible to place nodes next to each other using [proximity placement](https://docs.microsoft.com/en-us/azure/aks/reduce-latency-ppg). This however means that we can only use one availability zone, which isn't what we want.

### Compliance

**Patching**: the Azure EKS images are continuously updated by Azure but we need to activate them in the node pools. A future ADR will describe the exact process of how this enforced. Until that time we will monitor the AKS images and initiate an update manually to ensure we keep using the latest version.

**Policies**: AKS has an option to use [OPA/Gatekeeper](https://docs.microsoft.com/en-us/azure/governance/policy/concepts/policy-for-kubernetes?toc=/azure/aks/toc.json) as a policy engine to enforce rules. This kind of functionality is going to be crucial for security, best practices and more. Because of this, we want to have the full flexibility here and will therefore run our own policy engine.

**Certificate rotation**: it is possible to enable [certificate rotation](https://docs.microsoft.com/en-us/azure/aks/certificate-rotation) and although this is handy, it also the drawback that all pods will be restarted on CA rotation. As users will be connecting through our authentication mechanism, the only issue we see are the service accounts in Kubernetes. We are going to investigate options to minimise their usage but until that time we accept the risk of longer-lived certificates. In case of a security emergency, we can always trigger the CA rotation and deal with the fallout on a case-by-case basis.

## Consequences

A lot of things have been decided at this point, but there are still relatively large open topics that we need to address in the future. We will track these items in our issue tracker and pick them up as we go. There are also some accepted security risks that might deter people from using Opsteady at this point. Some of these issues we will definitely be improved in the future and some of them will be evaulated and decided based on user demand. Either way, the concerns are recorded in our compliance and security [documentation](../opsteady/security-and-compliance.md).
