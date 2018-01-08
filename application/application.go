package application

import (
	"errors"
	"github.com/sda0/eth_scanner/storage"
	"github.com/sda0/eth_scanner/storage/model"
	"log"
	"math"
	"time"
)

type Application struct {
	config         Config
	storageManager *storage.Manager
}

func NewApplication(cfg Config) (*Application, error) {
	app := &Application{
		config: cfg,
	}

	sm, err := storage.NewStorageManagerInstance(cfg.Storages)
	if err != nil {
		return app, errors.New("Unable to init storage manager: " + err.Error())
	}
	app.storageManager = sm

	return app, nil
}

func (a *Application) Config() Config {
	return a.config
}

func (a *Application) StorageManager() *storage.Manager {
	return a.storageManager
}

func (a *Application) Stop() {
	log.Println("App terminated")

	a.StorageManager().Close()
}

func (a *Application) Start() (err error) {
	var nextBlock, maxBlock int64
	var affected int
	var block model.Block
	for {
		nextBlock = a.storageManager.GetLocalDB().GetLastBlockNumber() + 1
		maxBlock = a.storageManager.GetBlockchain().GetLastBlockNumber()

		if a.config.Debug {
			nextBlock = int64(math.Max(float64(maxBlock-10000), float64(nextBlock))) //only last 10 000 blocks
		}

		log.Printf("Eth last block %d, next block to parse %d", maxBlock, nextBlock)
		for ; nextBlock <= maxBlock; nextBlock++ {
			block = a.storageManager.GetBlockchain().GetBlock(nextBlock)
			affected, err = a.storageManager.GetLocalDB().Save(block)
			if err != nil {
				return
			}
			log.Printf("Block %d (transaction count %d) imported to local db ", nextBlock, affected)
		}

		time.Sleep(2 * time.Second)
	}
}
