// Split the list of the components into a separate package so there is no cyclic dependency
package components

import (
	"github.com/opsteady/opsteady/cli/component"
	foundationAzure "github.com/opsteady/opsteady/foundation/azure/cicd"
	managementBootstrap "github.com/opsteady/opsteady/management/bootstrap/cicd"
	managementInfra "github.com/opsteady/opsteady/management/infra/cicd"
	managementVaultConfig "github.com/opsteady/opsteady/management/vault/config/cicd"
	managementVaultInfra "github.com/opsteady/opsteady/management/vault/infra/cicd"
	// New component path should be added here
)

// Components contains a list of component initializers
var Components = make(map[string]component.Initialize)

func init() {
	Components["management-bootstrap"] = &managementBootstrap.ManagementBootstrap{}
	Components["management-infra"] = &managementInfra.ManagementInfra{}
	Components["management-vault-infra"] = &managementVaultInfra.ManagementVaultInfra{}
	Components["management-vault-config"] = &managementVaultConfig.ManagementVaultConfig{}
	Components["foundation-azure"] = &foundationAzure.FoundationAzure{}
}
