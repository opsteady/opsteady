// Some of the calls in the Opsteady CLI can take some time
// The data that comes out of it can mostly be cached in memory
// so that it can be used while running.
// Some of the data can also be cached for a longer period
// therefor it is possible to cache the data to a file.
// The caching package is responsible for storing the data
// in-memory or file depending of the needs.
package cache

import (
	"encoding/json"
	"os"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

const (
	creationTimeName = "creationTime"
	readWrite        = 0600
)

// Cache is used to store and retrieve data from memory or file
type Cache interface {
	Store(string, map[string]interface{}, time.Duration)
	Retrieve(string) map[string]interface{}
}

// CacheImpl is the Cache interface implementation
type CacheImpl struct { //nolint
	lock        sync.Mutex
	data        map[string]map[string]interface{}
	filePath    string
	logger      *zerolog.Logger
	storeToFile bool
}

// NewCache returns cache in memory only
func NewCache(logger *zerolog.Logger) (Cache, error) {
	return NewFileCache("", logger)
}

// NewFileCache returns cache which stores data in memory and file
func NewFileCache(filePath string, logger *zerolog.Logger) (Cache, error) {
	logger.Debug().Msg("Initialize Cache")

	cache := &CacheImpl{
		data:        make(map[string]map[string]interface{}),
		logger:      logger,
		filePath:    filePath,
		storeToFile: false,
	}

	if filePath != "" {
		logger.Debug().Str("file", filePath).Msg("Using the cache file")

		if err := cache.initializeFromFile(); err != nil {
			return nil, errors.Wrapf(err, "could not initialize cached file %s", filePath)
		}

		cache.storeToFile = true
	}

	logger.Debug().Msg("Cache initialized")

	return cache, nil
}

// Store data to the cache (memory or file)
// Set the creation time to now
func (c *CacheImpl) Store(dataID string, data map[string]interface{}, ttl time.Duration) {
	c.logger.Debug().Msg("Add creation time to the data")

	data[creationTimeName] = time.Now().UTC().Add(ttl).Format(time.RFC3339)

	c.logger.Trace().Msg("Store the data into memory")
	c.lock.Lock()
	c.data[dataID] = data
	c.lock.Unlock()

	if c.storeToFile {
		c.logger.Trace().Msg("Store to file enabled, saving to cache file")
		c.lock.Lock()
		err := c.saveToFile()
		c.lock.Unlock()

		if err != nil {
			c.logger.Error().Err(err).Msg("continue even though failed to save the cache to file")
		}
	}
}

// Retrieve data from the cache
// Also check if the data will be still valid for 10 minutes
// The creation time should always be available, if not data is considered invalid
func (c *CacheImpl) Retrieve(dataID string) map[string]interface{} {
	c.logger.Debug().Str("id", dataID).Msg("Retrieving cached object")
	c.lock.Lock()
	data, ok := c.data[dataID]
	c.lock.Unlock()

	if !ok {
		c.logger.Trace().Str("id", dataID).Msg("Data not in cache")

		return nil
	}

	creationTimeInterface, containsCreationTime := c.data[dataID][creationTimeName]

	c.logger.Trace().Msg("Check for creation time as it should always be there")

	if !containsCreationTime {
		c.logger.Error().Msg("creation time not available in the cache, continue without cache")

		return nil
	}

	creationTimeString := creationTimeInterface.(string)
	creationTime, err := time.Parse(time.RFC3339, creationTimeString)

	if err != nil {
		c.logger.Error().Msg("could not parse creation time from cache, continue without cache")

		return nil
	}

	c.logger.Trace().Msg("Check if data will still be valid 10 minutes from now")

	now := time.Now().UTC().Add(time.Minute * 10) //nolint

	if creationTime.After(now) {
		c.logger.Trace().Str("id", dataID).Msg("Data is in cache and valid")

		return data
	}

	c.logger.Warn().Str("time", creationTimeString).Str("id", dataID).Msg("Cache data expired, continue without cache")

	return nil
}

func (c *CacheImpl) saveToFile() error {
	c.logger.Trace().Msg("Marshal cache data into JSON")
	data, err := json.MarshalIndent(c.data, "", "\t")

	if err != nil {
		return errors.Wrap(err, "could not marshal cache data into JSON")
	}

	c.logger.Trace().Msg("Save the JSON cache data to file")

	if err := os.WriteFile(c.filePath, data, readWrite); err != nil {
		return errors.Wrap(err, "could not save JSON cache data into file")
	}

	c.logger.Info().Str("file", c.filePath).Msg("Saved cache data into file")

	return nil
}

func (c *CacheImpl) initializeFromFile() error {
	c.logger.Trace().Msg("Read from the cache file and decode into the object")

	if _, err := os.Stat(c.filePath); os.IsNotExist(err) {
		c.logger.Debug().Str("file", c.filePath).Msg("File does not exist, not reading it")

		return nil
	}

	content, err := os.ReadFile(c.filePath)

	if err != nil {
		return errors.Wrapf(err, "could not file %s", c.filePath)
	}

	err = json.Unmarshal(content, &c.data)

	if err != nil {
		return errors.Wrapf(err, "could not decode data from file %s", c.filePath)
	}

	c.logger.Debug().Str("file", c.filePath).Msg("Successfully read cached file")

	return nil
}
