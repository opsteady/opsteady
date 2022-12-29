package component

import (
	"fmt"
	"os"
)

// PrepareTerraformBackend prepares the terraform to use remote storage
// in the management subscription
func (c *DefaultComponent) PrepareTerraformBackend() {
	c.Logger.Info().Msg("Preparing Terraform environment...")

	mgmtCreds, err := c.Credentials.Azure("management")
	if err != nil {
		c.Logger.Fatal().Err(err).Msg("could not get management credentials to prepare terraform")
	}

	tfBackendCreds := fmt.Sprintf(`
	subscription_id = "%s"
	tenant_id = "%s"
	client_id = "%s"
	client_secret = "%s"`,
		c.GlobalConfig.ManagementSubscriptionID,
		c.GlobalConfig.TenantID,
		mgmtCreds["client_id"].(string),
		mgmtCreds["client_secret"].(string))

	blobKey := fmt.Sprintf("%s/%s/%s.tfstate", c.CurrentTarget, c.PlatformID, c.Metadata.Name)
	c.Logger.Info().Str("backend", blobKey).Msg("Using backend blob key")
	tfBackendCreds = fmt.Sprintf("key = \"%s\"\n%s", blobKey, tfBackendCreds)

	if err := os.WriteFile(c.TerraformBackendConfigPath, []byte(tfBackendCreds), readWriteAll); err != nil {
		c.Logger.Fatal().Err(err).Str("path", c.TerraformBackendConfigPath).Msg("could not write the backend config file")
	}
}
