package storage

import (
	"errors"
	"fmt"
	"github.com/sda0/eth_scanner/storage/dbconnect"
	"github.com/sda0/eth_scanner/storage/model"
	"log"
	"strconv"
)

type Blockchain struct {
	connect *dbconnect.EthConnection
}

func (bc Blockchain) GetLastBlockNumber() int64 {
	client, err := bc.connect.GetRPCClient()
	if err != nil {
		panic(err)
	}

	response, err := client.Call("eth_blockNumber")
	if err != nil {
		panic(err)
	}

	res, err := response.GetString()
	if err != nil {
		panic(err)
	}

	result, _ := strconv.ParseInt(res[2:], 16, 64)

	GetCache().SetMaxBlockNumber(result)

	return result
}

func (bc Blockchain) GetBlock(blockId int64) (result model.Block) {
	client, err := bc.connect.GetRPCClient()
	if err != nil {
		panic(err)
	}

	blockNumber := "0x" + strconv.FormatInt(blockId, 16)

	response, err := client.Call("eth_getBlockByNumber", blockNumber, true)
	if err != nil {
		panic(err)
	}
	//	fmt.Printf("%#v", response.Result)
	err = response.GetObject(&result)
	if err != nil {
		panic(err)
	}

	return
}

func (bc Blockchain) SendTransaction(t *model.SendEth) (json string, err error) {
	client, err := bc.connect.GetRPCClient()
	if err != nil {
		panic(err)
	}

	params := map[string]interface{}{
		"from":  t.From,
		"to":    t.To,
		"value": t.GetWeiHexed(),
	}

	log.Println(params)
	response, err := client.Call("eth_sendTransaction", params)
	if err != nil {
		panic(err)
	}

	if response.Error != nil {
		err = errors.New(response.Error.Message)
		return
	}

	json = fmt.Sprintf("%#v", response.Result)
	return
}
