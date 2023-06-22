package wireproviders

import (
	"context"
	"crypto/tls"

	"oxylabs-homework/internal/adapters/quickclient"
)

func NewPublisherQUICClientFromConfig(ctx context.Context, conf ServerConfig, tlsConf *tls.Config) (*quickclient.PublisherClient, error) {
	return quickclient.NewPublisherClient(ctx, conf.PublisherPort, tlsConf)
}

func NewSubscriberQUICClientFromConfig(ctx context.Context, conf ServerConfig, tlsConf *tls.Config) (*quickclient.SubscriberClient, error) {
	return quickclient.NewSubscriberClient(ctx, conf.SubscriberPort, tlsConf)
}
