// Split the list of the components into a separate package so there is no cyclic dependency
package components

import (
	"github.com/opsteady/opsteady/cli/component"
	managementBootstrap "github.com/opsteady/opsteady/management/bootstrap/cicd"
	// New component path should be added here
)

// Components contains a list of component initializers
var Components = make(map[string]component.Initialize)

func init() {
	Components["management-bootstrap"] = &managementBootstrap.ManagementBootstrap{}
	// New component should be added here
}
