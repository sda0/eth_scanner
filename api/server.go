package api

import (
	"encoding/json"
	"errors"
	"github.com/sda0/eth_scanner/storage"
	"github.com/sda0/eth_scanner/storage/model"
	"log"
	"net/http"
)

type Server struct {
	config         Config
	storageManager *storage.Manager
}

func (a *Server) Start() error {
	http.HandleFunc("/sendEth", a.sendEth)
	http.HandleFunc("/getLast", a.getLast)
	return http.ListenAndServe(a.config.ApiAddress, nil)
}

func NewServer(cfg Config) (*Server, error) {
	app := &Server{
		config: cfg,
	}

	sm, err := storage.NewStorageManagerInstance(cfg.Storages)
	if err != nil {
		return app, errors.New("Unable to init storage manager: " + err.Error())
	}
	app.storageManager = sm

	return app, nil
}

func (a *Server) Config() Config {
	return a.config
}

func (a *Server) StorageManager() *storage.Manager {
	return a.storageManager
}

func (a *Server) Stop() {
	log.Println("App terminated")
	a.StorageManager().Close()
}

func (a *Server) sendEth(response http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	sendEth := &model.SendEth{}
	err := decoder.Decode(sendEth)
	if err != nil {
		panic(err)
	}
	resultJson, err := a.storageManager.GetBlockchain().SendTransaction(sendEth)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(err.Error()))
	} else {
		response.WriteHeader(http.StatusOK)
		response.Write([]byte(resultJson))
	}
}

func (a *Server) getLast(response http.ResponseWriter, r *http.Request) {
	cache := storage.GetCache()
	var s string
	/* @todo сделать кеш мидлварью */

	s, err := cache.GetLastTransactions()
	if err != nil {
		s = a.storageManager.GetLocalDB().GetLastSinceBlock(cache.GetLastCursor(), cache.GetMaxBlockNumber())
	}

	response.Write([]byte(s))
}
