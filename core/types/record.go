package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type PbftRecordHeader struct {
	Number   *big.Int
	Hash     common.Hash
	GasLimit *big.Int
	GasUsed  *big.Int
	Time     *big.Int
}

type PbftRecord struct {
	header       *PbftRecordHeader
	transactions Transactions
	sig          []*string
}


func (r *PbftRecord) Hash() common.Hash {
	return r.header.Hash
}

func (r *PbftRecord) Number() *big.Int {
	return r.header.Number
}


func (r *PbftRecord) Header() *PbftRecordHeader { return r.header }


func CopyRecord(r *PbftRecord) *PbftRecord {
	// TODO: copy all record fields
	record := &PbftRecord{

	}

	return record
}