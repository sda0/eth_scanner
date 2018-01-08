package dbconnect

import (
	"errors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresConnection struct {
	cfg     PostgresConfig
	db      *sqlx.DB
	counter int64
}

type PostgresConfig struct {
	Disabled          bool   `json:"disabled"`
	ConnectionString  string `json:"connectionString"`
	MaxOpenConnection int    `json:"maxOpenConnections"`
	MaxIdleConnection int    `json:"maxIdleConnections"`
}

func (pc *PostgresConfig) Validate() error {
	if pc.Disabled {
		return nil
	}
	if pc.ConnectionString == "" {
		return errors.New("PostgresConfig error: connection string not set")
	}

	return nil
}

func (postgresConnection *PostgresConnection) Connect() (err error) {
	if postgresConnection.cfg.Disabled {
		return errors.New("postgres connection config disabled")
	}
	postgresConnection.db, err = sqlx.Connect("postgres", postgresConnection.cfg.ConnectionString)
	if err != nil {
		panic(err)
	}
	if postgresConnection.cfg.MaxOpenConnection > 0 {
		postgresConnection.db.SetMaxOpenConns(postgresConnection.cfg.MaxOpenConnection)
	}
	if postgresConnection.cfg.MaxIdleConnection > 0 {
		postgresConnection.db.SetMaxIdleConns(postgresConnection.cfg.MaxIdleConnection)
	}
	err = postgresConnection.db.Ping()
	return
}

func (postgresConnection *PostgresConnection) AddCount() {
	postgresConnection.counter++
}

func (postgresConnection *PostgresConnection) GetCount() int64 {
	return postgresConnection.counter
}

func (postgresConnection *PostgresConnection) GetConnection() (*sqlx.DB, error) {
	if postgresConnection.db == nil {
		if err := postgresConnection.Connect(); err != nil {
			return nil, err
		}
	}
	return postgresConnection.db, nil
}

func (postgresConnection *PostgresConnection) SetCfg(cfg PostgresConfig) {
	postgresConnection.cfg = cfg
}

func (postgresConnection *PostgresConnection) Close() {
	if postgresConnection.db != nil {
		postgresConnection.db.Close()
	}
}
