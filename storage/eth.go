package storage

import (
	"fmt"
	"strconv"
)

func (eth ethSingleton) GetLastBlockNumber() int {
	response, err := eth.rpcClient.Call("eth_blockNumber")
	if err != nil {
		panic(err)
	}

	var res string
	response.GetObject(&res)

	fmt.Printf("%s", res)
	result, err := strconv.ParseInt(res, 16, 32)
	if err != nil {
		panic(err)
	}
	return int(result)
}
