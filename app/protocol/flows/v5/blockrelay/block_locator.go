package blockrelay

import (
	"github.com/astrix-network/astrixd/app/appmessage"
	"github.com/astrix-network/astrixd/app/protocol/common"
	"github.com/astrix-network/astrixd/app/protocol/protocolerrors"
	"github.com/astrix-network/astrixd/domain/consensus/model/externalapi"
)

func (flow *handleRelayInvsFlow) sendGetBlockLocator(highHash *externalapi.DomainHash, limit uint32) error {
	msgGetBlockLocator := appmessage.NewMsgRequestBlockLocator(highHash, limit)
	return flow.outgoingRoute.Enqueue(msgGetBlockLocator)
}

func (flow *handleRelayInvsFlow) receiveBlockLocator() (blockLocatorHashes []*externalapi.DomainHash, err error) {
	for {
		message, err := flow.incomingRoute.DequeueWithTimeout(common.DefaultTimeout)
		if err != nil {
			return nil, err
		}

		switch message := message.(type) {
		case *appmessage.MsgInvRelayBlock:
			flow.invsQueue = append(flow.invsQueue, invRelayBlock{Hash: message.Hash, IsOrphanRoot: false})
		case *appmessage.MsgBlockLocator:
			return message.BlockLocatorHashes, nil
		default:
			return nil,
				protocolerrors.Errorf(true, "received unexpected message type. "+
					"expected: %s, got: %s", appmessage.CmdBlockLocator, message.Command())
		}
	}
}
