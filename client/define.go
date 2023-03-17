package client

import "encoding/json"

//JsonRpc version
const JSON_RPC_VERSION = "2.0"

//JsonRpcRequest object in rpc
type JsonRpcRequest struct {
	Version string        `json:"jsonrpc"`
	Id      string        `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

//JsonRpcResponse object response for JsonRpcRequest
type JsonRpcResponse struct {
	Id     string          `json:"id"`
	Error  int64           `json:"error"`
	Desc   string          `json:"desc"`
	Result json.RawMessage `json:"result"`
}

const (
	RPC_CONSOLIDATED_BLOCK_NUMBER    = "zkevm_consolidatedBlockNumber"
	RPC_IS_BLOCK_CONSOLIDATED        = "zkevm_isBlockConsolidated"
	RPC_GET_BATCH_BY_NUMBER          = "zkevm_getBatchByNumber"
	RPC_BATCH_NUMBER_BY_BLOCK_NUMBER = "zkevm_batchNumberByBlockNumber"
)
