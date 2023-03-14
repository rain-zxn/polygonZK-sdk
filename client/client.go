package client

import (
	"encoding/json"
	"fmt"
	"github.com/polynetwork/polygonZK-sdk/types"
	"github.com/polynetwork/polygonZK-sdk/utils"
	"strings"
	"sync/atomic"
)

type ClientMgr struct {
	rpc *RpcClient //Rpc client used the rpc api of ontology
	qid uint64
}

func (this *ClientMgr) NewRpcClient() *RpcClient {
	this.rpc = NewRpcClient()
	return this.rpc
}

func (this *ClientMgr) GetRpcClient() *RpcClient {
	return this.rpc
}

func (this *ClientMgr) getClient() *RpcClient {
	if this.rpc != nil {
		return this.rpc
	}
	return nil
}

func (this *ClientMgr) getNextQid() string {
	return fmt.Sprintf("%d", atomic.AddUint64(&this.qid, 1))
}

func (this *ClientMgr) GetConsolidatedBlockNumber() (uint64, error) {
	client := this.getClient()
	if client == nil {
		return 0, fmt.Errorf("don't have available client")
	}
	data, err := client.getConsolidatedBlockNumber(this.getNextQid())
	if err != nil {
		return 0, err
	}
	return utils.GetHexToUint64(data)
}

func (this *ClientMgr) IsBlockConsolidated(height uint64) (bool, error) {
	client := this.getClient()
	if client == nil {
		return false, fmt.Errorf("don't have available client")
	}
	block := fmt.Sprintf("%#x", height)
	data, err := client.isBlockConsolidated(this.getNextQid(), block)
	if err != nil {
		return false, err
	}
	return strings.EqualFold(string(data), "true"), nil
}

func (this *ClientMgr) GetBatchByNumber(batch uint64) (*types.RpcBatch, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client")
	}
	batchNumber := fmt.Sprintf("%#x", batch)
	data, err := client.getBatchByNumber(this.getNextQid(), batchNumber)
	if err != nil {
		return nil, err
	}
	rpcBatch := new(types.RpcBatch)
	err = json.Unmarshal(data, rpcBatch)
	if err != nil {
		return nil, fmt.Errorf("Unmarshal rpcBatch err")
	}
	return rpcBatch, nil
}
