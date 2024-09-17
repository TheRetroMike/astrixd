package client

import (
	"context"
	"github.com/astrix-network/astrixd/cmd/astrixwallet/daemon/server"
	"time"

	"github.com/pkg/errors"

	"github.com/astrix-network/astrixd/cmd/astrixwallet/daemon/pb"
	"google.golang.org/grpc"
)

// Connect connects to the astrixwalletd server, and returns the client instance
func Connect(address string) (pb.AstrixwalletdClient, func(), error) {
	// Connection is local, so 1 second timeout is sufficient
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(server.MaxDaemonSendMsgSize)))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, nil, errors.New("astrixwallet daemon is not running, start it with `astrixwallet start-daemon`")
		}
		return nil, nil, err
	}

	return pb.NewAstrixwalletdClient(conn), func() {
		conn.Close()
	}, nil
}
