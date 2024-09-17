package testapi

import (
	"github.com/astrix-network/astrixd/domain/consensus/model"
	"github.com/astrix-network/astrixd/domain/consensus/utils/txscript"
)

// TestTransactionValidator adds to the main TransactionValidator methods required by tests
type TestTransactionValidator interface {
	model.TransactionValidator
	SigCache() *txscript.SigCache
	SetSigCache(sigCache *txscript.SigCache)
}
