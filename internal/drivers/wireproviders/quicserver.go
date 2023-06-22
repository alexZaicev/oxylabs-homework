package wireproviders

import (
	"crypto/tls"

	"oxylabs-homework/internal/adapters/quicserver"
	"oxylabs-homework/internal/adapters/quicserver/handlers"
)

func NewPublisherQUICServerFromConfig(conf ServerConfig, handler *handlers.PublisherHandler, tlsConf *tls.Config) (*quicserver.PublisherServer, error) {
	return quicserver.NewPublisherServer(conf.PublisherPort, handler, tlsConf)
}

func NewSubscriberQUICServerFromConfig(conf ServerConfig, handler *handlers.SubscriberHandler, tlsConf *tls.Config) (*quicserver.SubscriberServer, error) {
	return quicserver.NewSubscriberServer(conf.SubscriberPort, handler, tlsConf)
}
