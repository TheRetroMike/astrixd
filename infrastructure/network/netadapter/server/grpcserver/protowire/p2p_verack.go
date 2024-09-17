package protowire

import (
	"github.com/astrix-network/astrixd/app/appmessage"
	"github.com/pkg/errors"
)

func (x *AstrixdMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "AstrixdMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *AstrixdMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
