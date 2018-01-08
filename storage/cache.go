package storage

import (
	"sync"
)

type cache struct {
	getLastRegistry int64
	getLastResult   string
	maxBlockNumber  int64
}

var cacheInstance *cache
var once sync.Once

func GetCache() *cache {
	once.Do(func() {
		cacheInstance = &cache{}
	})
	return cacheInstance
}

func (c *cache) SetMaxBlockNumber(i int64) *cache {
	if c.maxBlockNumber < i {
		c.maxBlockNumber = i
	}
	return c
}

func (c *cache) SetGetLastRegistry(i int64) *cache {
	c.getLastRegistry = i
	return c
}

func (c *cache) SetGetLastResult(s string) *cache {
	c.getLastResult = s
	return c
}

func (c *cache) GetMaxBlockNumber() int64 {
	return c.maxBlockNumber
}

func (c *cache) GetLastRegistry() int64 {
	return c.getLastRegistry
}

func (c *cache) GetLastResult() string {
	return c.getLastResult
}

