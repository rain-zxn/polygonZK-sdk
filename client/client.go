package client

import (
	"context"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/polynetwork/polygonZK-sdk/types"
)

type Client struct {
	*rpc.Client
}

func NewPolygonZkClient(url string) (*Client, error) {
	client, err := rpc.Dial(url)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func (p *Client) ConsolidatedBlockNumber() (uint64, error) {
	var result types.ArgUint64
	err := p.CallContext(context.Background(), &result, RPC_CONSOLIDATED_BLOCK_NUMBER)
	if err != nil {
		return 0, err
	}
	return uint64(result), nil
}

func (p *Client) IsBlockConsolidated(height uint64) (bool, error) {
	block := types.ArgUint64(height)
	var result bool
	err := p.CallContext(context.Background(), &result, RPC_IS_BLOCK_CONSOLIDATED, block)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (p *Client) GetBatchByNumber(batch uint64) (*types.RpcBatch, error) {
	batchNumber := types.ArgUint64(batch)
	result := new(types.RpcBatch)
	err := p.CallContext(context.Background(), result, RPC_GET_BATCH_BY_NUMBER, batchNumber)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (p *Client) BatchNumberByBlockNumber(height uint64) (uint64, error) {
	block := types.ArgUint64(height)
	var result types.ArgUint64
	err := p.CallContext(context.Background(), &result, RPC_BATCH_NUMBER_BY_BLOCK_NUMBER, block)
	if err != nil {
		return 0, err
	}
	return uint64(result), nil
}
