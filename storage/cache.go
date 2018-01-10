package storage

import (
	"sync"
	"github.com/sda0/eth_scanner/storage/model"
)

type cache struct {
	getLastRegistry model.BlockNumber
	maxBlockNumber  model.BlockNumber
}

var cacheInstance *cache
var once sync.Once

func GetCache() *cache {
	once.Do(func() {
		cacheInstance = &cache{}
	})
	return cacheInstance
}

func (c *cache) SetMaxBlockNumber(i model.BlockNumber) *cache {
	if c.maxBlockNumber < i {
		c.maxBlockNumber = i
	}
	return c
}

func (c *cache) GetMaxBlockNumber() model.BlockNumber {
	return c.maxBlockNumber
}

func (c *cache) GetLastRegistry() model.BlockNumber {
	return c.getLastRegistry
}


