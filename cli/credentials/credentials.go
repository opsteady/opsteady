// credentials package simplifies fetching of all kinds of credentials from Vault
package credentials

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/preview/subscription/mgmt/2019-10-01-preview/subscription"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/opsteady/opsteady/cli/cache"
	"github.com/opsteady/opsteady/cli/vault"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// DefaultTTL is used as default TTL when sending the information to Vault
const DefaultTTL = time.Hour * 1

// VaultCredentials provides authentication mechanisms
type VaultCredentials struct {
	vault  vault.Vault
	cache  cache.Cache
	logger *zerolog.Logger
}

// Credentials defines an interface for getting credentials
type Credentials interface {
	AWS(accountID string, ttl string) (map[string]interface{}, error)
	Azure(subscriptionID string) (map[string]interface{}, error)
	AzureAD() (map[string]interface{}, error)
	AKS(subscriptionID string) (map[string]interface{}, error)
}

// NewCredentials creates a new credentials struct
func NewCredentials(vault vault.Vault, cache cache.Cache, logger *zerolog.Logger) Credentials {
	logger.Debug().Msg("Initialize Credentials")

	return &VaultCredentials{
		vault:  vault,
		logger: logger,
		cache:  cache,
	}
}

// AWS retrieves AWS credentials
func (vc *VaultCredentials) AWS(accountID string, ttl string) (map[string]interface{}, error) {
	vc.logger.Info().Str("ttl", ttl).Str("account", accountID).Msg("Retrieve AWS credentials")
	fullID := fmt.Sprintf("AWS/%s", accountID)
	secret := vc.cache.Retrieve(fullID)

	if secret != nil {
		vc.logger.Info().Str("id", fullID).Msg("Using AWS credentials from cache")

		return secret, nil
	}

	ttlDuration, err := time.ParseDuration(ttl)

	if err != nil {
		return nil, errors.Wrapf(err, "could not parse the TTL %s as duration", ttl)
	}

	data := map[string]interface{}{
		"ttl": ttl,
	}

	vc.logger.Debug().Str("id", fullID).Msg("Requesting AWS credentials from Vault")
	secret, err = vc.vault.Write(fmt.Sprintf("aws/%s/sts/vault-aws-access", accountID), data)

	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve credentials from Vault for %s", fullID)
	}

	vc.cache.Store(fullID, secret, ttlDuration)
	vc.logger.Debug().Str("id", fullID).Msg("Returning retrieved credentials")

	return secret, nil
}

// Azure retrieves Azure credentials
func (vc *VaultCredentials) Azure(subscriptionID string) (map[string]interface{}, error) {
	path := fmt.Sprintf("azure/creds/%s", subscriptionID)
	cacheIndex := fmt.Sprintf("Azure/%s", subscriptionID)

	return vc.getAzureCreds(path, cacheIndex, subscriptionID)
}

// AKS retrieves service principal credentials for AKS
func (vc *VaultCredentials) AKS(subscriptionID string) (map[string]interface{}, error) {
	path := fmt.Sprintf("azure/creds/%s-k8s", subscriptionID)
	cacheIndex := fmt.Sprintf("AKS/%s", subscriptionID)

	return vc.getAzureCreds(path, cacheIndex, subscriptionID)
}

func (vc *VaultCredentials) getAzureCreds(path string, cacheID string, subscriptionID string) (map[string]interface{}, error) {
	secret := vc.cache.Retrieve(cacheID)

	if secret != nil {
		vc.logger.Info().Str("id", cacheID).Msg("Using Azure credentials from cache")

		return secret, nil
	}

	vc.logger.Debug().Str("id", cacheID).Msg("Requesting Azure credentials from Vault")
	secret, err := vc.vault.Read(path, nil)

	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve credentials from Vault for %s", path)
	}

	vc.logger.Info().Msg("Waiting for Azure credentials propagation.")

	clientID := secret["client_id"].(string)
	clientSecret := secret["client_secret"].(string)

	configPath := "azure/config"
	azureConfig, err := vc.vault.Read(configPath, nil)

	if err != nil {
		return nil, errors.Wrapf(err, "could not read the Azure secret backend configuration")
	}

	tenantID := ""
	if _, ok := azureConfig["tenant_id"]; ok {
		tenantID = azureConfig["tenant_id"].(string)
	}

	if tenantID == "" {
		return nil, errors.Wrapf(err, "could not get tenant ID from Azure secret backend configuration: %+v", azureConfig)
	}

	if err := vc.verifyCredentials(clientID, clientSecret, tenantID, subscriptionID); err != nil {
		return nil, err
	}

	vc.cache.Store(cacheID, secret, DefaultTTL)

	vc.logger.Info().Msg("Azure credentials are propagated.")
	vc.logger.Debug().Str("id", cacheID).Msg("Returning retrieved credentials")

	return secret, nil
}

func (vc *VaultCredentials) verifyCredentials(clientID, clientSecret, tenantID, subscriptionID string) error {
	config := auth.NewClientCredentialsConfig(clientID, clientSecret, tenantID)
	authorizer, err := config.Authorizer()

	if err != nil {
		return errors.Wrapf(err, "could not authorize the Azure Go client")
	}

	subscriptionClient := subscription.NewSubscriptionsClient()
	subscriptionClient.Authorizer = authorizer

	var subFound bool

	tries := 0

OUTER:
	// We try to get the subscription list as a heuristic to know if the credential permissions have
	// been propagated. The amount of tries is arbitrary but should succeed most of the time within
	// five tries.
	for tries < 5 {
		for iter, err := subscriptionClient.ListComplete(context.Background()); iter.NotDone(); err = iter.Next() {
			if err != nil {
				return err
			}

			vc.logger.Info().Msg(fmt.Sprintf("Client has permissions for subscription '%s'", *iter.Value().DisplayName))

			if strings.EqualFold(*iter.Value().DisplayName, subscriptionID) {
				subFound = true

				break OUTER
			}
		}
	}

	if !subFound {
		return errors.Wrapf(err, "client does not have permissions for subscription '%s'", subscriptionID)
	}

	return nil
}

// AzureAD retrieves credentials used for reading from Azure AD.
func (vc *VaultCredentials) AzureAD() (map[string]interface{}, error) {
	fixedID := "AzureAD"

	secret := vc.cache.Retrieve(fixedID)

	if secret != nil {
		vc.logger.Info().Str("id", fixedID).Msg("Using AzureAD credentials from cache")

		return secret, nil
	}

	vc.logger.Debug().Str("id", fixedID).Msg("Requesting AzureAD credentials")
	secret, err := vc.vault.Read("azuread/creds/management", nil)

	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve credentials from Vault for %s", fixedID)
	}

	vc.cache.Store(fixedID, secret, DefaultTTL)

	vc.logger.Debug().Str("id", fixedID).Msg("Returning retrieved credentials")

	return secret, nil
}
