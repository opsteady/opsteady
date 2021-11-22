package cache

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

var (
	testDataCaching = make(map[string]interface{})
	logger          = &log.Logger
)

func init() {
	testDataCaching["test"] = "test data"
}
func TestStoreToCache(t *testing.T) {
	cache, err := NewCache(logger)
	assert.Nil(t, err, "Should not have initialization error")
	cache.Store("1", testDataCaching, time.Minute*1)
	assert.Equal(t, 1, len(cache.(*CacheImpl).data), "Expected cache to contain data")
}

func TestRetrieveFromCache(t *testing.T) {
	cache, err := NewCache(logger)
	assert.Nil(t, err, "Should not have initialization error")
	cache.Store("1", testDataCaching, time.Hour*1)
	tmp := cache.Retrieve("1")
	assert.Equal(t, tmp["test"], "test data", "Expected [test data]")
}

func TestRetrieveUnknownIdFromCache(t *testing.T) {
	cache, err := NewCache(logger)
	assert.Nil(t, err, "Should not have initialization error")

	tmp := cache.Retrieve("1")
	assert.Nil(t, tmp, "Expected to be nil")
}

func TestRetrieveFromCacheTTLExpired(t *testing.T) {
	cache, err := NewCache(logger)
	assert.Nil(t, err, "Should not have initialization error")
	cache.Store("1", testDataCaching, time.Minute*1)

	tmp := cache.Retrieve("1")
	assert.Nil(t, tmp, "Expected to be nil")
}

func TestRetrieveFromCachedFile(t *testing.T) {
	cacheFile := fmt.Sprintf("%s/cached_file", t.TempDir())
	cache, err := NewFileCache(cacheFile, logger)
	assert.Nil(t, err, "Should not have initialization error")
	cache.Store("1", testDataCaching, time.Hour*1)

	tmp := cache.Retrieve("1")
	assert.Equal(t, tmp["test"], "test data", "Expected [test data]")

	// Use new cache to see if it reads the file again
	secondCache, err := NewFileCache(cacheFile, logger)
	assert.Nil(t, err, "Should not have initialization error")

	seconTmp := secondCache.Retrieve("1")
	assert.Equal(t, seconTmp["test"], "test data", "Expected [test data]")
}

func TestConcurency(t *testing.T) {
	var wg sync.WaitGroup

	cache, err := NewCache(logger)
	assert.Nil(t, err, "Should not have initialization error")

	for i := 1; i <= 50; i++ {
		wg.Add(1)

		go func() {
			cache.Store(strconv.Itoa(rand.Int()), testDataCaching, time.Minute*1) //nolint
			wg.Done()
		}()
	}

	wg.Wait() // This test should not fail
}
