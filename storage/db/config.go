package db

type StoragesConfig struct {
	Postgres  PostgresConfig  `json:"postgres"`
	Eth     EthConfig     `json:"eth" bson:"eth"`
}

func (sc *StoragesConfig) Validate() error {
	if err := sc.Postgres.Validate(); err != nil {
		return err
	}
	return nil
}
