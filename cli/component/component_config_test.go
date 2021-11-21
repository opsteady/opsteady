package component

import (
	"testing"

	"github.com/opsteady/opsteady/cli/cache"
	"github.com/opsteady/opsteady/cli/vault"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestRetrieveConfig(t *testing.T) {
	conf := setupComponentConfigForTest(t)
	data, err := conf.RetrieveConfig("1", "dev", []string{"one", "two"})
	assert.NoError(t, err, "Shouldn't have an error")
	assert.Len(t, data, 2, "Should be 2")
	assert.Contains(t, data, "one_test", "Should contain one_test")
	assert.Contains(t, data, "two_test", "Should contain two_test")
}

func TestAddAndOverrideConfigManually(t *testing.T) {
	conf := setupComponentConfigForTest(t)
	data, err := conf.RetrieveConfig("1", "dev", []string{"one"})
	assert.NoError(t, err, "Shouldn't have an error")
	assert.Contains(t, data, "one_test", "Should contain one_test")

	// Add new
	conf.GeneralAddOrOverride("abc", "efg")
	data, err = conf.RetrieveConfig("1", "dev", []string{"one"})
	assert.NoError(t, err, "Shouldn't have an error")
	assert.Contains(t, data, "abc", "Should contain abc")

	// Override existing
	conf.GeneralAddOrOverride("one_test", "something else")
	data, err = conf.RetrieveConfig("1", "dev", []string{"one"})
	assert.NoError(t, err, "Shouldn't have an error")
	assert.Equal(t, "something else", data["one_test"], "Should contain two_test")
}

func setupComponentConfigForTest(t *testing.T) ComponentConfig {
	logger := zerolog.Nop()
	cache, err := cache.NewCache(&logger)

	if err != nil {
		t.Fatal(err)
	}

	vault := &vault.FakeVault{
		Requests: []string{},
	}

	return NewComponentConfig(cache, vault, &logger)
}
