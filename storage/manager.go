package storage

import (
	"github.com/sda0/eth_scanner/storage/dbconnect"
)

type Manager struct {
	Config             dbconnect.StoragesConfig
	PostgresConnection *dbconnect.PostgresConnection
	EthConnection      *dbconnect.EthConnection
}

func (sm *Manager) GetLocalDB() LocalDB {
	return LocalDB{connect: sm.PostgresConnection}
}

func (sm *Manager) GetBlockchain() Blockchain {
	return Blockchain{connect: sm.EthConnection}
}

func (sm *Manager) Close() {
	sm.PostgresConnection.Close()
}

func NewStorageManagerInstance(cfg dbconnect.StoragesConfig) (*Manager, error) {
	sm := &Manager{
		Config:             cfg,
		EthConnection:      ethConnect(cfg.Eth),
		PostgresConnection: postgresConnect(cfg.Postgres),
	}

	return sm, nil
}

func postgresConnect(cfg dbconnect.PostgresConfig) *dbconnect.PostgresConnection {
	postgresConnection := &dbconnect.PostgresConnection{}
	postgresConnection.SetCfg(cfg)
	return postgresConnection
}

func ethConnect(cfg dbconnect.EthConfig) *dbconnect.EthConnection {
	ethConnection := &dbconnect.EthConnection{}
	ethConnection.SetCfg(cfg)
	return ethConnection
}
