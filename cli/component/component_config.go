package component

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/opsteady/opsteady/cli/cache"
	"github.com/opsteady/opsteady/cli/vault"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// ComponentConfig interface to retrieve component config from Vault.
type ComponentConfig interface {
	RetrieveConfig(string, string, []string) (map[string]interface{}, error)
	GeneralAddOrOverride(string, string)
}

// ComponentConfigImpl is the implementation of the ComponentConfig interface.
type ComponentConfigImpl struct {
	TTL       time.Duration
	cache     cache.Cache
	logger    *zerolog.Logger
	vault     vault.Vault
	overrides map[string]string
}

// NewComponentConfig returns an implementation of ComponentConfig.
func NewComponentConfig(cache cache.Cache, vault vault.Vault, logger *zerolog.Logger) ComponentConfig {
	return &ComponentConfigImpl{
		cache:     cache,
		logger:    logger,
		vault:     vault,
		TTL:       time.Hour * 1,
		overrides: make(map[string]string),
	}
}

// GeneralAddOrOverride adds key values which will be always added regardless of which component you are using.
func (c *ComponentConfigImpl) GeneralAddOrOverride(key, value string) {
	c.overrides[key] = value
}

// RetrieveConfig retrieves the component config from Vault.
func (c *ComponentConfigImpl) RetrieveConfig(version, environment string, components []string) (map[string]interface{}, error) {
	componentID := fmt.Sprintf("%s-%s-%s", version, environment, strings.Join(components[:], "-"))
	settings := c.cache.Retrieve(componentID)
	if settings != nil {
		c.logger.Debug().Str("id", componentID).Msg("Using cached settings")
		return settings, nil
	}

	values := make(map[string]interface{})
	chanComponents := make(chan map[string]interface{}, len(components)+1)
	chanPlatform := make(chan map[string]interface{}, len(components)+1)
	chanPlatformTerraform := make(chan map[string]interface{}, len(components)+1)
	chanErrors := make(chan error, len(components)+1)

	var wg sync.WaitGroup

	c.logger.Debug().Msg("Fetch default settings for components")
	for _, component := range components {
		wg.Add(1)
		path := fmt.Sprintf("config/data/%s/component/%s", version, component)
		go c.fetchConfig(path, component, chanComponents, chanErrors, &wg)
	}

	c.logger.Debug().Msg("Fetch settings per environment per component")
	for _, component := range components {
		wg.Add(1)
		env := environment
		if strings.HasPrefix(component, "management-") {
			// Don't look at the platform env because it is the management env
			env = "management"
		}
		path := fmt.Sprintf("config/data/%s/platform/%s/%s", version, env, component)
		go c.fetchConfig(path, component, chanPlatform, chanErrors, &wg)
	}

	c.logger.Debug().Msg("Fetch Terraform output per environment per component")
	for _, component := range components {
		wg.Add(1)
		env := environment
		if strings.HasPrefix(component, "management-") {
			// Don't look at the platform env because it is the management env
			env = "management"
		}
		path := fmt.Sprintf("config/data/%s/platform/%s/%s-tf", version, env, component)
		go c.fetchConfig(path, component, chanPlatformTerraform, chanErrors, &wg)
	}

	wg.Wait()
	close(chanErrors)
	close(chanComponents)
	close(chanPlatform)
	close(chanPlatformTerraform)

	for err := range chanErrors {
		if err != nil {
			return nil, err
		}
	}

	c.logger.Debug().Msg("First get the components")
	for value := range chanComponents {
		for k, v := range value {
			values[k] = v
		}
	}

	c.logger.Debug().Msg("Now override with platform specific info")
	for value := range chanPlatform {
		for k, v := range value {
			values[k] = v
		}
	}

	c.logger.Debug().Msg("Now add Terraform output platform specific")
	for value := range chanPlatformTerraform {
		for k, v := range value {
			values[k] = v
		}
	}

	c.cache.Store(componentID, values, c.TTL)

	for k, v := range c.overrides {
		values[k] = v
	}
	return values, nil
}

func (c *ComponentConfigImpl) fetchConfig(path, component string, chanValues chan map[string]interface{}, chanErrors chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	values := make(map[string]interface{})

	secret, err := c.vault.Read(path, nil)
	if err != nil {
		chanErrors <- err
		return
	}

	if secret == nil {
		chanErrors <- errors.Errorf("Data is empty in path [%s]", path)
		return
	}
	if data, ok := secret["data"]; ok {
		// Data can be nil if it is a '-tf' secret which is automatically created and deleted
		// when deleted the secret is still their but without any data
		if data != nil {
			for key, value := range data.(map[string]interface{}) {
				values[key] = value
			}
		}
	}

	chanValues <- values
}
