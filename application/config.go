package application

import (
	"../storage/dbconnect"
	"encoding/json"
)

type Config struct {
	Storages dbconnect.StoragesConfig `json:"storages"`
	Tracing  bool                     `json:"trace,omitempty"`
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
