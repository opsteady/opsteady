package vault

import (
	"os"
	"testing"
	"time"

	kv "github.com/hashicorp/vault-plugin-secrets-kv"
	"github.com/hashicorp/vault/api"
	vaulthttp "github.com/hashicorp/vault/http"
	"github.com/hashicorp/vault/sdk/helper/logging"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/hashicorp/vault/vault"
	"github.com/opsteady/opsteady/cli/cache"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestWriteReadVault(t *testing.T) {
	cluster := NewTestVault(t)
	logger := zerolog.New(os.Stdout)

	cache, err := cache.NewCache(&logger)
	if err != nil {
		t.Fatal(err)
	}

	// Make sure the token is stored so we don't try to login
	cache.Store("admin", map[string]interface{}{"token": cluster.Cores[0].Client.Token()}, time.Hour*1)

	vaultImpl, err := NewVault(cluster.Cores[0].Client.Address(), "admin", true, cache, &logger)
	if err != nil {
		t.Fatal(err)
	}

	_, err = vaultImpl.Write("kv/data/test", map[string]interface{}{
		"data": map[string]interface{}{
			"hello": "world",
		},
	})
	assert.NoError(t, err, "No error should happen")

	data, err := vaultImpl.Read("kv/data/test", nil)
	assert.NoError(t, err, "No error should happen")

	assert.Len(t, data, 2, "Should contain one key")
	assert.Equal(t, data["data"].(map[string]interface{})["hello"], "world", "Should contain world")
}

// NewTestVault returns a test Vault server and client
func NewTestVault(t *testing.T) *vault.TestCluster {
	t.Helper()

	coreConfig := &vault.CoreConfig{
		LogicalBackends: map[string]logical.Factory{
			"kv": kv.Factory,
		},
		Logger: logging.NewVaultLogger(5), // Log only errors
	}

	cluster := vault.NewTestCluster(t, coreConfig, &vault.TestClusterOptions{
		HandlerFunc: vaulthttp.Handler,
		TempDir:     t.TempDir(),
		Logger:      logging.NewVaultLogger(5), // Log only errors
	})
	cluster.Start()

	// Create KV V2 mount
	err := cluster.Cores[0].Client.Sys().Mount("kv", &api.MountInput{
		Type: "kv",
		Options: map[string]string{
			"version": "2",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	return cluster
}
