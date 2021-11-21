package component

import (
	"fmt"
	"io/ioutil"
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

	if blobKey := c.determinBlobKey(); blobKey != "" {
		c.Logger.Info().Str("backend", blobKey).Msg("Using backend blob key")
		tfBackendCreds = fmt.Sprintf("key = \"%s\"\n%s", blobKey, tfBackendCreds)
	}

	if err := ioutil.WriteFile(c.TerraformBackendConfigPath, []byte(tfBackendCreds), 0644); err != nil {
		c.Logger.Fatal().Err(err).Str("path", c.TerraformBackendConfigPath).Msg("could not write the backend config file")
	}
}

func (c *DefaultComponent) determinBlobKey() string {
	// Try to determine which blob key to use for Terraform state
	var blobKey string
	if c.AwsID != "" && c.AzureID == "" { //nolint
		blobKey = fmt.Sprintf("%s/%s/%s.tfstate", "aws", c.AwsID, c.ComponentName)
	} else if c.AwsID == "" && c.AzureID != "" {
		blobKey = fmt.Sprintf("%s/%s/%s.tfstate", "azure", c.AzureID, c.ComponentName)
	} else if c.AwsID == "" && c.AzureID == "" {
		c.Logger.Fatal().Msg("Please specify a target AWS/Azure ID")
	} else if c.AwsID != "" && c.AzureID != "" {
		c.Logger.Info().Msg("You specified both an Azure and AWS ID, using the backend blob key in the Terraform provider")
	}
	return blobKey
}
