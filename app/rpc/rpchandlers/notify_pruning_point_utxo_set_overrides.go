package rpchandlers

import (
	"github.com/astrix-network/astrixd/app/appmessage"
	"github.com/astrix-network/astrixd/app/rpc/rpccontext"
	"github.com/astrix-network/astrixd/infrastructure/network/netadapter/router"
)

// HandleNotifyPruningPointUTXOSetOverrideRequest handles the respectively named RPC command
func HandleNotifyPruningPointUTXOSetOverrideRequest(context *rpccontext.Context, router *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	listener, err := context.NotificationManager.Listener(router)
	if err != nil {
		return nil, err
	}
	listener.PropagatePruningPointUTXOSetOverrideNotifications()

	response := appmessage.NewNotifyPruningPointUTXOSetOverrideResponseMessage()
	return response, nil
}
