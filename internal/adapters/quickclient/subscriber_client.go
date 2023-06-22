package quickclient

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/quic-go/quic-go"
)

type SubscriberClient struct {
	client *Client
	stream quic.Stream
}

func NewSubscriberClient(ctx context.Context, port uint16, tlsConf *tls.Config) (*SubscriberClient, error) {
	client, err := New(ctx, port, tlsConf)
	if err != nil {
		return nil, err
	}

	stream, err := client.conn.AcceptStream(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to accept stream: %s", err)
	}

	return &SubscriberClient{
		client: client,
		stream: stream,
	}, nil
}

func (c *SubscriberClient) ReadMessage() (string, error) {
	buffer := make([]byte, defaultBufferSize)
	n, err := c.stream.Read(buffer)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %s", err)
	}

	return string(buffer[:n]), nil
}

func (c *SubscriberClient) Stop() error {
	if err := c.client.Stop(); err != nil {
		return err
	}

	if err := c.stream.Close(); err != nil {
		return fmt.Errorf("failed to close stream: %s", err)
	}

	return nil
}
