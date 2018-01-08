package dbconnect

import (
	"errors"
	"github.com/ybbus/jsonrpc"
)

type EthConnection struct {
	cfg    EthConfig
	client *jsonrpc.RPCClient
}

type EthConfig struct {
	Disabled bool   `json:"disabled"`
	Endpoint string `json:"endpoint"`
}

func (ec *EthConfig) Validate() error {
	if ec.Disabled {
		return nil
	}
	if ec.Endpoint == "" {
		return errors.New("ethConfig error: endpoint not set")
	}

	return nil
}

func (rpc *EthConnection) Connect() (err error) {
	if rpc.cfg.Disabled {
		return errors.New("eth connection config disabled")
	}
	rpc.client = jsonrpc.NewRPCClient(rpc.cfg.Endpoint)
	return
}

func (rpc *EthConnection) GetRPCClient() (*jsonrpc.RPCClient, error) {
	if rpc.client == nil {
		if err := rpc.Connect(); err != nil {
			return nil, err
		}
	}
	return rpc.client, nil
}

func (rpc *EthConnection) SetCfg(cfg EthConfig) {
	rpc.cfg = cfg
}

func (rpc *EthConnection) Close() {
	if rpc.client != nil {
		rpc.client = nil
	}
}
