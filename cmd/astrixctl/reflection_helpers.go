package main

import (
	"reflect"
	"unicode"

	"github.com/astrix-network/astrixd/infrastructure/network/netadapter/server/grpcserver/protowire"
)

// protobuf generates the command types with two types:
// 1. A concrete type that holds the fields of the command bearing the name of the command with `RequestMessage` as suffix
// 2. A wrapper that implements isAstrixdMessage_Payload, having a single field pointing to the concrete command
//    bearing the name of the command with `AstrixdMessage_` prefix and `Request` suffix

// unwrapCommandType converts a reflect.Type signifying a wrapper type into the concrete request type
func unwrapCommandType(requestTypeWrapped reflect.Type) reflect.Type {
	return requestTypeWrapped.Field(0).Type.Elem()
}

// unwrapCommandValue convertes a reflect.Value of a pointer to a wrapped command into a concrete command
func unwrapCommandValue(commandValueWrapped reflect.Value) reflect.Value {
	return commandValueWrapped.Elem().Field(0)
}

// isFieldExported returns true if the given field is exported.
// Currently the only way to check this is to check if the first rune in the field's name is upper case.
func isFieldExported(field reflect.StructField) bool {
	return unicode.IsUpper(rune(field.Name[0]))
}

// generateAstrixdMessage generates a wrapped AstrixdMessage with the given `commandValue`
func generateAstrixdMessage(commandValue reflect.Value, commandDesc *commandDescription) (*protowire.AstrixdMessage, error) {
	commandWrapper := reflect.New(commandDesc.typeof)
	unwrapCommandValue(commandWrapper).Set(commandValue)

	astrixdMessage := reflect.New(reflect.TypeOf(protowire.AstrixdMessage{}))
	astrixdMessage.Elem().FieldByName("Payload").Set(commandWrapper)
	return astrixdMessage.Interface().(*protowire.AstrixdMessage), nil
}

// pointerToValue returns a reflect.Value that represents a pointer to the given value
func pointerToValue(valuePointedTo reflect.Value) reflect.Value {
	pointer := reflect.New(valuePointedTo.Type())
	pointer.Elem().Set(valuePointedTo)
	return pointer
}
