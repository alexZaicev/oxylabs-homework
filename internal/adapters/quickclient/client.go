package quickclient

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/quic-go/quic-go"
)

const (
	defaultBufferSize = 1024
)

type Client struct {
	conn quic.Connection
}

func New(ctx context.Context, port uint16, tlsConf *tls.Config) (*Client, error) {
	conn, err := quic.DialAddr(
		ctx,
		fmt.Sprintf("localhost:%d", port),
		tlsConf,
		&quic.Config{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to dial quic server on port [:%d]: %s", port, err)
	}

	return &Client{
		conn: conn,
	}, nil
}

func (c *Client) Stop() error {
	if err := c.conn.CloseWithError(quic.ApplicationErrorCode(quic.NoError), ""); err != nil {
		return fmt.Errorf("failed to close connection: %s", err)
	}

	return nil
}
