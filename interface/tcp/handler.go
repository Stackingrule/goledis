package tcp

import (
	"context"
	"net"
)

type handler interface {
	handler(ctx context.Context, conn net.Conn)
	Close() error
}
