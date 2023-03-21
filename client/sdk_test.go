package client

import (
	"fmt"
	"testing"
)

func initClient() *PolygonZkClient {
	url := "https://rpc.public.zkevm-test.net"
	client, err := NewPolygonZkClient(url)
	if err != nil {
		panic(err)
	}
	return client
}

func Test_ConsolidatedBlockNumber(t *testing.T) {
	fmt.Println(initClient().ConsolidatedBlockNumber())
}

func Test_IsBlockConsolidated(t *testing.T) {
	fmt.Println(initClient().IsBlockConsolidated(6666))
}

func Test_GetBatchByNumber(t *testing.T) {
	x, _ := initClient().GetBatchByNumber(16103)
	fmt.Println(x.Transactions[len(x.Transactions)-1].String())
}

func Test_BatchNumberByBlockNumber(t *testing.T) {
	fmt.Println(initClient().BatchNumberByBlockNumber(6666))
}
