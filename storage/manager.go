package storage
import (
	"./db"
)

type Manager struct {
	Config              db.StoragesConfig
	PostgresConnection  *db.PostgresConnection
	EthConnection  *db.EthConnection
}

func (sm *Manager) GetPgStorage() Pg {
	return Pg{PostgresConnection: sm.PostgresConnection}
}

func (sm *Manager) GetEthStorage() Eth {
	return Eth{EthConnection: sm.EthConnection}
}


func (sm *Manager) Close() {
	sm.PostgresConnection.Close()
}

func NewStorageManagerInstance(cfg db.StoragesConfig) (*Manager, error) {
	sm := &Manager{
		Config: cfg,
		EthConnection: ethConnect(cfg.Eth),
		PostgresConnection: postgresConnect(cfg.Postgres),
	}

	return sm, nil
}

func postgresConnect(cfg db.PostgresConfig) (*db.PostgresConnection) {
	postgresConnection := &db.PostgresConnection{}
	postgresConnection.SetCfg(cfg)
	return postgresConnection
}

func ethConnect(cfg db.EthConfig) (*db.EthConnection) {
	ethConnection := &db.EthConnection{}
	ethConnection.SetCfg(cfg)
	return ethConnection
}