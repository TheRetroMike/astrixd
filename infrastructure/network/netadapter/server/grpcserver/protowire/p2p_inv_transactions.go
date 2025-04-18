package protowire

import (
	"github.com/astrix-network/astrixd/app/appmessage"
	"github.com/pkg/errors"
)

func (x *AstrixdMessage_InvTransactions) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "AstrixdMessage_InvTransactions is nil")
	}
	return x.InvTransactions.toAppMessage()
}

func (x *InvTransactionsMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "InvTransactionsMessage is nil")
	}
	if len(x.Ids) > appmessage.MaxInvPerTxInvMsg {
		return nil, errors.Errorf("too many hashes for message "+
			"[count %d, max %d]", len(x.Ids), appmessage.MaxInvPerTxInvMsg)
	}

	ids, err := protoTransactionIDsToDomain(x.Ids)
	if err != nil {
		return nil, err
	}
	return &appmessage.MsgInvTransaction{TxIDs: ids}, nil

}

func (x *AstrixdMessage_InvTransactions) fromAppMessage(msgInvTransaction *appmessage.MsgInvTransaction) error {
	if len(msgInvTransaction.TxIDs) > appmessage.MaxInvPerTxInvMsg {
		return errors.Errorf("too many hashes for message "+
			"[count %d, max %d]", len(msgInvTransaction.TxIDs), appmessage.MaxInvPerTxInvMsg)
	}

	x.InvTransactions = &InvTransactionsMessage{
		Ids: wireTransactionIDsToProto(msgInvTransaction.TxIDs),
	}
	return nil
}
