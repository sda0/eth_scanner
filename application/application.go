package application

import (
	"errors"
	"../storage"
	"../tracing"
	"log"
	"github.com/sda0/eth_scanner/eth"
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

	//set debug from config
	tracing.SetTracing(cfg.Tracing)

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

func (a *Application) Start() error {
	ethReader := eth.GetEth()
	max := ethReader.GetLastBlockNumber()
	i := pg.GetLastBlockNumber
	for ;i<max; i++ {
		transactions := ethReader.GetBlock(i)
		pg.transactionBegin
		defer pg.transactionRollback
		for transaction := range transactions {
			pg.insert(transaction)
		}
		pg.transactionCommit
	}

	return a.HttpApplication.ListenAndServe()
}

func (a *Application) ReopenLogs() {
}
