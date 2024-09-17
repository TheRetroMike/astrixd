package protowire

import (
	"github.com/astrix-network/astrixd/app/appmessage"
	"github.com/pkg/errors"
)

func (x *AstrixdMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "AstrixdMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *AstrixdMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
