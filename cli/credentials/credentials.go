// credentials package simplifies fetching of all kinds of credentials from Vault
package credentials

import (
	"fmt"
	"time"

	"github.com/opsteady/opsteady/cli/cache"
	"github.com/opsteady/opsteady/cli/vault"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// DefaultTTL is used as default TTL when sending the information to Vault
const DefaultTTL = time.Hour * 1

// VaultCredentials provides authentication mechanisms
type VaultCredentials struct {
	vault      vault.Vault
	cache      cache.Cache
	logger     *zerolog.Logger
	AzureSleep time.Duration
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
		vault:      vault,
		logger:     logger,
		cache:      cache,
		AzureSleep: 20 * time.Second,
	}
}

// AWS retrieves AWS credentials
func (vc *VaultCredentials) AWS(accountID string, ttl string) (map[string]interface{}, error) {
	vc.logger.Info().Str("ttl", ttl).Str("account", accountID).Msg("Retrieve AWS credentials")
	id := fmt.Sprintf("AWS/%s", accountID)
	secret := vc.cache.Retrieve(id)
	if secret != nil {
		vc.logger.Info().Str("id", id).Msg("Using AWS credentials from cache")
		return secret, nil
	}

	ttlDuration, err := time.ParseDuration(ttl)
	if err != nil {
		return nil, errors.Wrapf(err, "could not parse the TTL %s as duration", ttl)
	}

	data := map[string]interface{}{
		"ttl": ttl,
	}

	vc.logger.Debug().Str("id", id).Msg("Requesting AWS credentials from Vault")
	secret, err = vc.vault.Write(fmt.Sprintf("aws/%s/sts/vault-aws-access", accountID), data)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve credentials from Vault for %s", id)
	}

	vc.cache.Store(id, secret, ttlDuration)
	vc.logger.Debug().Str("id", id).Msg("Returning retrieved credentials")
	return secret, nil
}

// Azure retrieves Azure credentials
func (vc *VaultCredentials) Azure(subscriptionID string) (map[string]interface{}, error) {
	path := fmt.Sprintf("azure/creds/%s", subscriptionID)
	cacheIndex := fmt.Sprintf("Azure/%s", subscriptionID)

	return vc.getAzureCreds(path, cacheIndex)
}

// AKS retrieves service principal credentials for AKS
func (vc *VaultCredentials) AKS(subscriptionID string) (map[string]interface{}, error) {
	path := fmt.Sprintf("azure/creds/%s-k8s", subscriptionID)
	cacheIndex := fmt.Sprintf("AKS/%s", subscriptionID)

	return vc.getAzureCreds(path, cacheIndex)
}

func (vc *VaultCredentials) getAzureCreds(path string, id string) (map[string]interface{}, error) {
	secret := vc.cache.Retrieve(id)
	if secret != nil {
		vc.logger.Info().Str("id", id).Msg("Using Azure credentials from cache")
		return secret, nil
	}

	vc.logger.Debug().Str("id", id).Msg("Requesting Azure credentials from Vault")
	secret, err := vc.vault.Read(path, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve credentials from Vault for %s", id)
	}

	// Azure API is eventually consistent with permissions, therefore we wait an arbitrary amount of time.
	// This way we make sure that the new service principal has permissions on the subscription.
	vc.logger.Info().Dur("wait", vc.AzureSleep).Msg("Waiting for credentials to be processed by Azure")
	time.Sleep(vc.AzureSleep)

	vc.cache.Store(id, secret, DefaultTTL)

	vc.logger.Debug().Str("id", id).Msg("Returning retrieved credentials")
	return secret, nil
}

// AzureAD retrieves credentials used for reading from Azure AD.
func (vc *VaultCredentials) AzureAD() (map[string]interface{}, error) {
	id := "AzureAD"

	secret := vc.cache.Retrieve(id)
	if secret != nil {
		vc.logger.Info().Str("id", id).Msg("Using AzureAD credentials from cache")
		return secret, nil
	}

	vc.logger.Debug().Str("id", id).Msg("Requesting AzureAD credentials")
	secret, err := vc.vault.Read("azuread/creds/management", nil)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve credentials from Vault for %s", id)
	}

	vc.cache.Store(id, secret, DefaultTTL)

	vc.logger.Debug().Str("id", id).Msg("Returning retrieved credentials")
	return secret, nil
}
