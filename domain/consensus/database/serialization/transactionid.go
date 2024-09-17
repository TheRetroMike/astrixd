package serialization

import (
	"github.com/astrix-network/astrixd/domain/consensus/model/externalapi"
	"github.com/astrix-network/astrixd/domain/consensus/utils/transactionid"
)

// DbTransactionIDToDomainTransactionID converts DbTransactionId to DomainTransactionID
func DbTransactionIDToDomainTransactionID(dbTransactionID *DbTransactionId) (*externalapi.DomainTransactionID, error) {
	return transactionid.FromBytes(dbTransactionID.TransactionId)
}

// DomainTransactionIDToDbTransactionID converts DomainTransactionID to DbTransactionId
func DomainTransactionIDToDbTransactionID(domainTransactionID *externalapi.DomainTransactionID) *DbTransactionId {
	return &DbTransactionId{TransactionId: domainTransactionID.ByteSlice()}
}
