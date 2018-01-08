package dbconnect

type StoragesConfig struct {
	Postgres PostgresConfig `json:"postgres"`
	Eth      EthConfig      `json:"geth"`
}

func (sc *StoragesConfig) Validate() error {
	if err := sc.Postgres.Validate(); err != nil {
		return err
	}
	if err := sc.Eth.Validate(); err != nil {
		return err
	}
	return nil
}
