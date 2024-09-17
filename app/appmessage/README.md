wire
====

[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://choosealicense.com/licenses/isc/)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/astrix-network/astrixd/wire)
=======

Package wire implements the astrix wire protocol.

## Astrix Message Overview

The astrix protocol consists of exchanging messages between peers. Each message
is preceded by a header which identifies information about it such as which
astrix network it is a part of, its type, how big it is, and a checksum to
verify validity. All encoding and decoding of message headers is handled by this
package.

To accomplish this, there is a generic interface for astrix messages named
`Message` which allows messages of any type to be read, written, or passed
around through channels, functions, etc. In addition, concrete implementations
of most all astrix messages are provided. All of the details of marshalling and 
unmarshalling to and from the wire using astrix encoding are handled so the 
caller doesn't have to concern themselves with the specifics.

## Reading Messages Example

In order to unmarshal astrix messages from the wire, use the `ReadMessage`
function. It accepts any `io.Reader`, but typically this will be a `net.Conn`
to a remote node running a astrix peer. Example syntax is:

```Go
	// Use the most recent protocol version supported by the package and the
	// main astrix network.
	pver := wire.ProtocolVersion
	astrixnet := wire.Mainnet

	// Reads and validates the next astrix message from conn using the
	// protocol version pver and the astrix network astrixnet. The returns
	// are a appmessage.Message, a []byte which contains the unmarshalled
	// raw payload, and a possible error.
	msg, rawPayload, err := wire.ReadMessage(conn, pver, astrixnet)
	if err != nil {
		// Log and handle the error
	}
```

See the package documentation for details on determining the message type.

## Writing Messages Example

In order to marshal astrix messages to the wire, use the `WriteMessage`
function. It accepts any `io.Writer`, but typically this will be a `net.Conn`
to a remote node running a astrix peer. Example syntax to request addresses
from a remote peer is:

```Go
	// Use the most recent protocol version supported by the package and the
	// main bitcoin network.
	pver := wire.ProtocolVersion
	astrixnet := wire.Mainnet

	// Create a new getaddr astrix message.
	msg := wire.NewMsgGetAddr()

	// Writes a astrix message msg to conn using the protocol version
	// pver, and the astrix network astrixnet. The return is a possible
	// error.
	err := wire.WriteMessage(conn, msg, pver, astrixnet)
	if err != nil {
		// Log and handle the error
	}
```
