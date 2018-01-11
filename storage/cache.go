package storage

import (
	"container/heap"
	json2 "encoding/json"
	"errors"
	"github.com/sda0/eth_scanner/storage/model"
	"strings"
	"sync"
)

const limitConfirm = 3

type cache struct {
	mutex            sync.RWMutex
	getLastCursor    model.BlockNumber
	maxBlockNumber   model.BlockNumber
	lastTransactions model.SortedTransactions
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
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.maxBlockNumber < i {
		c.maxBlockNumber = i
	}
	return c
}

func (c *cache) GetMaxBlockNumber() model.BlockNumber {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.maxBlockNumber
}

func (c *cache) GetLastCursor() model.BlockNumber {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.getLastCursor
}

func (c *cache) SetLastCursor(blockNumber model.BlockNumber) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.getLastCursor = blockNumber
}

func (c *cache) GetLastTransactions() (json string, err error) {
	var str []byte

	c.mutex.Lock()
	defer c.mutex.Unlock()

	for i := range c.lastTransactions {
		c.lastTransactions[i].Confirmations = int(c.maxBlockNumber - c.lastTransactions[i].GetBlockNumber())
		str, _ = json2.Marshal(c.lastTransactions[i])
		json += string(str) + ","
	}

	if json == "" {
		return "", errors.New("getLast cache is empty")
	}

	c.getLastCursor = c.maxBlockNumber

	//все у кого более limitConfirm - попаем
	go func() {
		max := c.maxBlockNumber
		for {
			if c.lastTransactions.PopLimited(max-limitConfirm) == nil {
				break
			}
		}
	}()

	json = "[" + strings.Trim(json, ",") + "]"
	return
}

func (c *cache) SaveLastTransactions(block model.Block) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	date := block.GetDate()
	blockNumber := block.GetNumber()

	for _, t := range block.Transactions {
		item := &model.LastTransaction{
			Date:          date,
			To:            t.To,
			Amount:        t.GetValueEth(),
			Confirmations: 0,
		}
		item.SetBlock(blockNumber)
		heap.Push(&c.lastTransactions, item)
	}

	return
}
