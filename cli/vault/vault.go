// Vault package is used to login to Vault using OIDC to retrieve a token
// and to read normal secrets or (write) to get access to other services
// like Azure or AWS
package vault

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
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
}

// VaultImpl is the Vault interface implementation
type VaultImpl struct {
	client *api.Client
	logger *zerolog.Logger
	cache  cache.Cache
}

// NewVault creates Vault
func NewVault(address, role string, insecure bool, cache cache.Cache, logger *zerolog.Logger) (Vault, error) {
	logger.Debug().Msg("Initialize Vault")
	client, err := createClient(address, insecure, logger)
	if err != nil {
		return nil, err
	}

	vault := &VaultImpl{
		client: client,
		logger: logger,
		cache:  cache,
	}

	tokenMap := cache.Retrieve(role)
	if tokenMap != nil {
		vault.client.SetToken(tokenMap["token"].(string))
		return vault, nil
	}

	logger.Info().Str("role", role).Msg("Token not available in cache, logging in")
	if err := vault.oidcLogin(role); err != nil {
		return nil, err
	}

	cache.Store(role, map[string]interface{}{"token": vault.client.Token()}, tokenTTL)

	return vault, nil
}

func createClient(address string, insecure bool, logger *zerolog.Logger) (*api.Client, error) {
	logger.Debug().Msg("Creating Vault client")
	retryClient := retryablehttp.NewClient()

	transport := &http.Transport{
		Proxy:           http.ProxyFromEnvironment,
		MaxIdleConns:    10,
		IdleConnTimeout: 5 * time.Second,
	}

	// Enable insecure Vault
	if insecure {
		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	httpClient := retryClient.StandardClient()
	httpClient.Timeout = 30 * time.Second
	httpClient.Transport = transport

	client, err := api.NewClient(&api.Config{
		Address:    address,
		HttpClient: httpClient,
	})
	if err != nil {
		return nil, errors.Wrap(err, "could not create the Vault client")
	}

	return client, nil
}

func (v *VaultImpl) oidcLogin(role string) error {
	v.logger.Info().Str("role", role).Msg("Login using OIDC method")
	oidcHandler := &oidc.CLIHandler{}
	oidcData := map[string]string{
		"role":          role,
		"listenaddress": "0.0.0.0",
		"port":          "8250",
	}

	secret, err := oidcHandler.Auth(v.client, oidcData)
	if err != nil {
		return errors.Wrap(err, "could not get the Vault token")
	}

	v.logger.Info().Str("role", role).Msg("Successfully logged in")
	v.client.SetToken(secret.Auth.ClientToken)
	return nil
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
