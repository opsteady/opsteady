// Split the list of the components into a separate package so there is no cyclic dependency
package components

import (
	capabilitiesCertificatesAWS "github.com/opsteady/opsteady/capabilities/certificates/aws/cicd"
	capabilitiesCertificatesAzure "github.com/opsteady/opsteady/capabilities/certificates/azure/cicd"
	capabilitiesDNSAWS "github.com/opsteady/opsteady/capabilities/dns/aws/cicd"
	capabilitiesDNSAzure "github.com/opsteady/opsteady/capabilities/dns/azure/cicd"
	capabilitiesLoadbalancing "github.com/opsteady/opsteady/capabilities/loadbalancing/cicd"
	cli "github.com/opsteady/opsteady/cicd"
	"github.com/opsteady/opsteady/cli/component"
	dockerBase "github.com/opsteady/opsteady/docker/base/cicd"
	dockerCicd "github.com/opsteady/opsteady/docker/cicd/cicd"
	foundationAWS "github.com/opsteady/opsteady/foundation/aws/cicd"
	foundationAzure "github.com/opsteady/opsteady/foundation/azure/cicd"
	kubernetesAWSCluster "github.com/opsteady/opsteady/kubernetes/aws/cluster/cicd"
	kubernetesAWSLoadbalancing "github.com/opsteady/opsteady/kubernetes/aws/loadbalancing/cicd"
	kubernetesAWSNetworkPolicies "github.com/opsteady/opsteady/kubernetes/aws/network-policies/cicd"
	kubernetesAWSStorageEBS "github.com/opsteady/opsteady/kubernetes/aws/storage/ebs/cicd"
	kubernetesAWSStorageEFS "github.com/opsteady/opsteady/kubernetes/aws/storage/efs/cicd"
	kubernetesAzureCluster "github.com/opsteady/opsteady/kubernetes/azure/cluster/cicd"
	kubernetesAzurePodIdentity "github.com/opsteady/opsteady/kubernetes/azure/pod-identity/cicd"
	kubernetesBootstrap "github.com/opsteady/opsteady/kubernetes/bootstrap/cicd"
	managementBootstrap "github.com/opsteady/opsteady/management/bootstrap/cicd"
	managementInfra "github.com/opsteady/opsteady/management/infra/cicd"
	managementVaultConfig "github.com/opsteady/opsteady/management/vault/config/cicd"
	managementVaultInfra "github.com/opsteady/opsteady/management/vault/infra/cicd"
)

// Components contains a list of component initializers
var Components = make(map[string]component.Initialize)

func init() {
	Components["management-bootstrap"] = &managementBootstrap.ManagementBootstrap{}
	Components["management-infra"] = &managementInfra.ManagementInfra{}
	Components["management-vault-infra"] = &managementVaultInfra.ManagementVaultInfra{}
	Components["management-vault-config"] = &managementVaultConfig.ManagementVaultConfig{}
	Components["foundation-azure"] = &foundationAzure.FoundationAzure{}
	Components["foundation-aws"] = &foundationAWS.FoundationAWS{}
	Components["kubernetes-aws-cluster"] = &kubernetesAWSCluster.KubernetesAWSCluster{}
	Components["kubernetes-bootstrap"] = &kubernetesBootstrap.KubernetesBootstrap{}
	Components["kubernetes-aws-storage-ebs"] = &kubernetesAWSStorageEBS.KubernetesAWSStorageEBS{}
	Components["kubernetes-aws-storage-efs"] = &kubernetesAWSStorageEFS.KubernetesAWSStorageEFS{}
	Components["kubernetes-aws-network-policies"] = &kubernetesAWSNetworkPolicies.KubernetesAWSNetworkPolicies{}
	Components["kubernetes-aws-loadbalancing"] = &kubernetesAWSLoadbalancing.KubernetesAWSLoadbalancing{}
	Components["kubernetes-azure-pod-identity"] = &kubernetesAzurePodIdentity.KubernetesAzurePodIdentity{}
	Components["kubernetes-azure-cluster"] = &kubernetesAzureCluster.KubernetesAzure{}
	Components["capabilities-dns-aws"] = &capabilitiesDNSAWS.CapabilitiesDNSAWS{}
	Components["capabilities-dns-azure"] = &capabilitiesDNSAzure.CapabilitiesDNSAzure{}
	Components["capabilities-certificates-aws"] = &capabilitiesCertificatesAWS.CapabilitiesCertificatesAWS{}
	Components["capabilities-certificates-azure"] = &capabilitiesCertificatesAzure.CapabilitiesCertificatesAzure{}
	Components["capabilities-loadbalancing"] = &capabilitiesLoadbalancing.CapabilitiesLoadbalacing{}
	Components["docker-base"] = &dockerBase.DockerBase{}
	Components["docker-cicd"] = &dockerCicd.DockerCicd{}
	Components["cli"] = &cli.OpsteadyCli{}
}
