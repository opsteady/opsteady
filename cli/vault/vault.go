// Vault package is used to login to Vault using OIDC to retrieve a token
// and to read normal secrets or (write) to get access to other services
// like Azure or AWS
package vault

import (
	"time"

	oidc "github.com/hashicorp/vault-plugin-auth-jwt"
	"github.com/hashicorp/vault/api"
	"github.com/opsteady/opsteady/cli/cache"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

const tokenTTL = time.Hour * 12

// Vault is a simple wrapper for the Vault client
type Vault interface {
	Read(string, map[string][]string) (map[string]interface{}, error)
	Write(string, map[string]interface{}) (map[string]interface{}, error)
	GetToken() string
	GetAddress() string
}

// VaultImpl is the Vault interface implementation
type VaultImpl struct { //nolint
	client *api.Client
	logger *zerolog.Logger
	cache  cache.Cache
}

// NewVault creates Vault
// Token can be empty, if so OIDC login method will be used
func NewVault(address, role, token string, insecure bool, cache cache.Cache, logger *zerolog.Logger) (Vault, error) {
	logger.Debug().Msg("Initialize Vault")

	config := api.DefaultConfig()
	err := config.ConfigureTLS(&api.TLSConfig{Insecure: insecure})

	if err != nil {
		return nil, errors.Wrap(err, "could not configure TLS")
	}

	client, err := api.NewClient(config)

	if err != nil {
		return nil, errors.Wrap(err, "could not create the Vault client")
	}

	if err := client.SetAddress(address); err != nil {
		return nil, errors.Wrap(err, "could not set Vault address")
	}

	if token == "" {
		tokenMap := cache.Retrieve(role)

		if tokenMap == nil {
			logger.Info().Str("role", role).Msg("Token not available in cache, logging in")

			var err error

			if token, err = oidcLogin(role, client, logger); err != nil {
				return nil, err
			}

			cache.Store(role, map[string]interface{}{"token": token}, tokenTTL)
		} else {
			token = tokenMap["token"].(string)
		}
	}

	client.SetToken(token)

	vault := &VaultImpl{
		client: client,
		logger: logger,
		cache:  cache,
	}

	return vault, nil
}

func oidcLogin(role string, client *api.Client, logger *zerolog.Logger) (string, error) {
	logger.Info().Str("role", role).Msg("Login using OIDC method")

	oidcHandler := &oidc.CLIHandler{}
	oidcData := map[string]string{
		"role":          role,
		"listenaddress": "0.0.0.0",
		"port":          "8250",
	}

	secret, err := oidcHandler.Auth(client, oidcData)
	if err != nil {
		return "", errors.Wrap(err, "could not get the Vault token")
	}

	logger.Info().Str("role", role).Msg("Successfully logged in")

	return secret.Auth.ClientToken, nil
}

// Read reads a secret from a specified Vault path
func (v *VaultImpl) Read(path string, data map[string][]string) (map[string]interface{}, error) {
	v.logger.Debug().Str("path", path).Msg("Read from Vault")
	secret, err := v.client.Logical().ReadWithData(path, data)

	if err != nil {
		return nil, errors.Wrapf(err, "could not read from Vault path %s", path)
	}

	if secret == nil {
		v.logger.Warn().Str("path", path).Msg("Secret is empty")

		return make(map[string]interface{}), nil
	}

	return secret.Data, nil
}

// Write to vault to generate a new secret
func (v *VaultImpl) Write(path string, data map[string]interface{}) (map[string]interface{}, error) {
	v.logger.Debug().Str("path", path).Msg("Write to Vault")
	secret, err := v.client.Logical().Write(path, data)

	if err != nil {
		return nil, errors.Wrapf(err, "could not write to Vault path %s", path)
	}

	return secret.Data, nil
}

// GetToken returns the Vault token for others to use, for example Terraform
func (v *VaultImpl) GetToken() string {
	return v.client.Token()
}

// GetAddress returns the Vault address for others to use, for example Terraform
func (v *VaultImpl) GetAddress() string {
	return v.client.Address()
}
