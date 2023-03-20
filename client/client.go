package client

import (
	"context"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/polynetwork/polygonZK-sdk/types"
)

func ConsolidatedBlockNumber(rpcClient *rpc.Client) (uint64, error) {
	var result types.ArgUint64
	err := rpcClient.CallContext(context.Background(), &result, RPC_CONSOLIDATED_BLOCK_NUMBER)
	if err != nil {
		return 0, err
	}
	return uint64(result), nil
}

func IsBlockConsolidated(rpcClient *rpc.Client, height uint64) (bool, error) {
	block := types.ArgUint64(height)
	var result bool
	err := rpcClient.CallContext(context.Background(), &result, RPC_IS_BLOCK_CONSOLIDATED, block)
	if err != nil {
		return false, err
	}
	return result, nil
}

func GetBatchByNumber(rpcClient *rpc.Client, batch uint64) (*types.RpcBatch, error) {
	batchNumber := types.ArgUint64(batch)
	result := new(types.RpcBatch)
	err := rpcClient.CallContext(context.Background(), result, RPC_GET_BATCH_BY_NUMBER, batchNumber)
	if err != nil {
		return result, err
	}
	return result, nil
}

func BatchNumberByBlockNumber(rpcClient *rpc.Client, height uint64) (uint64, error) {
	block := types.ArgUint64(height)
	var result types.ArgUint64
	err := rpcClient.CallContext(context.Background(), &result, RPC_BATCH_NUMBER_BY_BLOCK_NUMBER, block)
	if err != nil {
		return 0, err
	}
	return uint64(result), nil
}
