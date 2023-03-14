package types

import (
	"github.com/ethereum/go-ethereum/common"
	"strconv"
	"strings"
)

type argUint64 uint64

// MarshalText marshals into text
func (b argUint64) MarshalText() ([]byte, error) {
	buf := make([]byte, 2) //nolint:gomnd
	copy(buf, `0x`)
	buf = strconv.AppendUint(buf, uint64(b), 16)
	return buf, nil
}

// UnmarshalText unmarshals from text
func (b *argUint64) UnmarshalText(input []byte) error {
	str := strings.TrimPrefix(string(input), "0x")
	num, err := strconv.ParseUint(str, 16, 64)
	if err != nil {
		return err
	}
	*b = argUint64(num)
	return nil
}

type RpcBatch struct {
	Number              argUint64      `json:"number"`
	Coinbase            common.Address `json:"coinbase"`
	StateRoot           common.Hash    `json:"stateRoot"`
	GlobalExitRoot      common.Hash    `json:"globalExitRoot"`
	AccInputHash        common.Hash    `json:"accInputHash"`
	Timestamp           argUint64      `json:"timestamp"`
	SendSequencesTxHash *common.Hash   `json:"sendSequencesTxHash"`
	VerifyBatchTxHash   *common.Hash   `json:"verifyBatchTxHash"`
	Transactions        []common.Hash  `json:"transactions"`
}
