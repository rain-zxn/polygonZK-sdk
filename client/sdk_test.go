package client

import (
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"testing"
)

func initClient() *rpc.Client {
	client, err := rpc.Dial("https://rpc.public.zkevm-test.net")
	if err != nil {
		panic(err)
	}
	return client
}

func Test_ConsolidatedBlockNumber(t *testing.T) {
	fmt.Println(ConsolidatedBlockNumber(initClient()))
}

func Test_IsBlockConsolidated(t *testing.T) {
	fmt.Println(IsBlockConsolidated(initClient(), 6666))
}

func Test_GetBatchByNumber(t *testing.T) {
	x, _ := GetBatchByNumber(initClient(), 16103)
	fmt.Println(x.Transactions[len(x.Transactions)-1].String())
}

func Test_BatchNumberByBlockNumber(t *testing.T) {
	fmt.Println(BatchNumberByBlockNumber(initClient(), 6666))
}
