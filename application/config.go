package application

import (
	"encoding/json"
	"github.com/sda0/eth_scanner/storage/dbconnect"
)

type Config struct {
	Storages dbconnect.StoragesConfig `json:"storages"`
	Debug    bool                     `json:"debug,omitempty"`
}

func (c *Config) Validate() error {
	if err := c.Storages.Validate(); err != nil {
		return err
	}

	return nil
}

func (c *Config) LoadFromJson(rawConfig json.RawMessage) (err error) {
	err = json.Unmarshal(rawConfig, &c)
	return
}
