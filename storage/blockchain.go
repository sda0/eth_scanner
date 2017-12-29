package storage

import (
	"fmt"
	"github.com/sda0/eth_scanner/storage/dbconnect"
	"strconv"
)

type Blockchain struct {
	connect *dbconnect.EthConnection
}

func (bc Blockchain) GetLastBlockNumber() int {
	client, err := bc.connect.GetRPCClient()
	if err != nil {
		panic(err)
	}

	response, err := client.Call("eth_blockNumber")
	if err != nil {
		panic(err)
	}

	var res string
	err = response.GetObject(&res)
	if err == nil {
		panic(err)
	}

	fmt.Printf("%s", res)
	result, err := strconv.ParseInt(res, 16, 32)
	if err != nil {
		panic(err)
	}
	return int(result)
}
