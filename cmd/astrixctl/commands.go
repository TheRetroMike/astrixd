package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/astrix-network/astrixd/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.AstrixdMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.AstrixdMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.AstrixdMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.AstrixdMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.AstrixdMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.AstrixdMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.AstrixdMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.AstrixdMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.AstrixdMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.AstrixdMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.AstrixdMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.AstrixdMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.AstrixdMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.AstrixdMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.AstrixdMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.AstrixdMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.AstrixdMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.AstrixdMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.AstrixdMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.AstrixdMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.AstrixdMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.AstrixdMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.AstrixdMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.AstrixdMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.AstrixdMessage_BanRequest{}),
	reflect.TypeOf(protowire.AstrixdMessage_UnbanRequest{}),
}

type commandDescription struct {
	name       string
	parameters []*parameterDescription
	typeof     reflect.Type
}

type parameterDescription struct {
	name   string
	typeof reflect.Type
}

func commandDescriptions() []*commandDescription {
	commandDescriptions := make([]*commandDescription, len(commandTypes))

	for i, commandTypeWrapped := range commandTypes {
		commandType := unwrapCommandType(commandTypeWrapped)

		name := strings.TrimSuffix(commandType.Name(), "RequestMessage")
		numFields := commandType.NumField()

		var parameters []*parameterDescription
		for i := 0; i < numFields; i++ {
			field := commandType.Field(i)

			if !isFieldExported(field) {
				continue
			}

			parameters = append(parameters, &parameterDescription{
				name:   field.Name,
				typeof: field.Type,
			})
		}
		commandDescriptions[i] = &commandDescription{
			name:       name,
			parameters: parameters,
			typeof:     commandTypeWrapped,
		}
	}

	return commandDescriptions
}

func (cd *commandDescription) help() string {
	sb := &strings.Builder{}
	sb.WriteString(cd.name)
	for _, parameter := range cd.parameters {
		_, _ = fmt.Fprintf(sb, " [%s]", parameter.name)
	}
	return sb.String()
}
