package main

import (
	rpc "github.com/ybbus/jsonrpc"
	"fmt"
)

type Result struct{

}

func main() {
	rpcClient := rpc.NewRPCClient("http://172.18.0.3:9000/")

	response, err := rpcClient.Call("eth_protocolVersion", "id:67")
	if err != nil {
		panic(err)
	}

	var res string
	response.GetObject(&res)

	fmt.Printf("%s", res)
}