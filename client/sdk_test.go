package client

import (
	"fmt"
	"testing"
)

func initClient() *ClientMgr {
	client := new(ClientMgr)
	client.NewRpcClient().SetAddress("https://rpc.public.zkevm-test.net")
	return client
}

func Test_GetCurrentBlockHeight(t *testing.T) {
	fmt.Println(initClient().GetConsolidatedBlockNumber())
}

func Test_IsBlockConsolidated(t *testing.T) {
	height := uint64(6666)
	fmt.Println(initClient().IsBlockConsolidated(height))
}

func Test_GetBatchByNumber(t *testing.T) {
	batch := uint64(6666)
	fmt.Println(initClient().GetBatchByNumber(batch))
}
