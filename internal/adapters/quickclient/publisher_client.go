package quickclient

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/quic-go/quic-go"
)

type PublisherClient struct {
	client *Client
	stream quic.Stream
}

func NewPublisherClient(ctx context.Context, port uint16, tlsConf *tls.Config) (*PublisherClient, error) {
	client, err := New(ctx, port, tlsConf)
	if err != nil {
		return nil, err
	}

	stream, err := client.conn.OpenStreamSync(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to open stream: %s", err)
	}

	return &PublisherClient{
		client: client,
		stream: stream,
	}, nil
}

func (c *PublisherClient) SendMessage(msg string) error {
	if _, err := c.stream.Write([]byte(msg)); err != nil {
		return fmt.Errorf("failed to send message: %s", err)
	}

	return nil
}

func (c *PublisherClient) ReadMessage() (string, error) {
	buffer := make([]byte, defaultBufferSize)
	n, err := c.stream.Read(buffer)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %s", err)
	}

	return string(buffer[:n]), nil
}

func (c *PublisherClient) Stop() error {
	if err := c.client.Stop(); err != nil {
		return err
	}

	if err := c.stream.Close(); err != nil {
		return fmt.Errorf("failed to close stream: %s", err)
	}

	return nil
}
