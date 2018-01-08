package api

import (
	"encoding/json"
	"errors"
	"github.com/sda0/eth_scanner/storage/dbconnect"
)

type Config struct {
	Storages   dbconnect.StoragesConfig `json:"storages"`
	ApiAddress string                   `json:"apiAddress"`
	Debug      bool                     `json:"debug,omitempty"`
}

func (c *Config) Validate() error {
	if err := c.Storages.Validate(); err != nil {
		return err
	}

	if c.ApiAddress == "" {
		return errors.New("apiAddress is not set in config file")
	}

	return nil
}

func (c *Config) LoadFromJson(rawConfig json.RawMessage) (err error) {
	err = json.Unmarshal(rawConfig, &c)
	return
}
