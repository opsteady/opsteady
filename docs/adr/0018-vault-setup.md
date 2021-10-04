# 18. vault-setup

Date: 2021-10-06

## Status

Status: Accepted on 2021-10-06
Builds on [0007-management-connectivity,md](0007-management-connectivity.md) on 2021-10-06
Builds on [0008-vault.md](0008-vault.md) on 2021-10-06
Builds on [0017-component-configuration.md](0017-component-configuration.md) on 2021-10-06

## Context

As described in the above ("builds-on") ADRs, Vault takes a critical role in the Opsteady platform. It is the central location for configuration and credentials management of all platforms. This ADR describes how we configure and run Vault itself.

## Decision

### Authentication

We will configure Vault with two authentication mechanisms: OIDC and JWT.

The OIDC mechanism is used by human operators to access Vault. The OIDC implementation is handled by the Opsteady Azure AD. The OIDC token contains the claims that will grant access to the Vault with an appropriate policy. The primary claim which determines the level of access, is is the 'groups' claim. As an example, if you are a member of the 'platform-admin' group, you will be granted access to the Vault with full administrative privileges.

The JWT mechanism is used by Github actions CI/CD. Vault will validate the presented JWT token with the Github actions issuer and allow access to the Vault with an appropriate policy for CI/CD purposes. This policy does not grant full administrative privileges and is primarily used to generate appropriate credentials and read component configuration. Currently we are using the platform-admin role for CI/CD purposes. This needs to be replaced with the above mentioned CI/CD policy.

Of course, the default token authentication mechanism is always activated. The above mentioned authentication mechanisms actually result in a temporary token that is presented to Vault.

### Credentials

We will configure Vault with two secrets backends: AWS and Azure.

Both of these backends generate on-the-fly, temporary credentials for a particular cloud. This allows the human operator or CI/CD process to perform the tasks needed to manage cloud infrastructure.

For Azure, Vault generates a service principal which will be granted 'owner' access to the subscription that is being targeted.

For AWS, Vault generates a temporary access key for a role in the target account. This role grants the necessary permissions to manage the infrastructure in the AWS account.

### Network Access

The information stored in Vault is highly sensitive and should be protected at all costs from unauthorized access. The above mentioned authentication mechanisms are the main barrier for access to the Vault but the endpoint for Vault, which is an HTTP API service, can also be shielded on the network level.

We are going to expose the Vault endpoint on the Internet while we are building up the platform. Putting the endpoint on a private network complicates the CI/CD setup and overall management of the platform. We understand that this is a risk but we accept this risk for now. Once the Opsteady platform matures, and we move towards onboarding of customers, we will put Vault (and related security sensitive infrastructure like Azure Key Vault) on a private network with gated access controls (proxy) in place. The exact setup for this still needs to be determined.

### High-availability and Storage

Vault has several different options for persisting the data that is stored inside the Vault. We are using the internal Raft-based storage option. The other storage options require the setup of external systems and significantly increase the complexity of running Vault. The internal storage gives us all the benefits of Vault with a relatively straight-forward implementation.

The data in Vault is crucial to manage the Opsteady platforms. As a first line of defense against data loss and unavailability, we run Vault in a distributed setup with three Vault servers on different virtual machines that exist in different availability zones. This allows for the failure of one Vault server and still be able to use Vault. The data will be replicated across these servers, so there is always the possibility to restore data from one of the existing Vault servers. This already gives us a high degree of data reliability. The storage disks themselves are Azure based disks that are provisioned through Kubernetes persistent volumes. They exist on their own, separately from the virtual machines.

In case of a catastrophic failure, where we lose the entire Vault cluster, we will ensure that we have backups of the Vault data in a durable and secure storage location in the cloud. At the time of writing the backups have not been implemented yet. This is something that we will do as the Opsteady platform matures and we are moving towards onboarding customers.

### Configuration

The initial configuration and defaults for the platform components are added manually to Vault. After bootstrapping the management environment, the Vault will be seeded with its configuration values as described in the management setup documentation. We will manually configure the components in Vault for each platform that we onboard. This might be replaced by an automated solution in the future.

## Consequences

The decisions in this ADR regarding network access and storage backups mandate that we need to do this work at a later time. For now, we accept the risk (brute-force hacking or DDoS attack on the Vault API endpoint and a catastrophic data loss) that these decisions present. Some manual work is needed to configure the components for new platforms until we find a way to make this semi-automatic. CI/CD permissions need to be improved in conjunction with the roles implementation.
