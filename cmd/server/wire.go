//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"oxylabs-homework/internal/adapters/quicserver/handlers"
	"oxylabs-homework/internal/drivers/wireproviders"
)

func initialize(outboundCh chan string, inboundCh <-chan string) (application, error) {
	wire.Build(
		// Options
		wireproviders.NewOptions,
		// Config
		wireproviders.NewConfig,
		wire.FieldsOf(new(wireproviders.Config), "Logger", "Server", "TLS"),
		// Logger
		wireproviders.NewLoggerConfig,
		wireproviders.NewLoggerFromConfig,
		// TLS
		wireproviders.NewTLSConfig,
		// Handlers
		handlers.NewPublisherHandler,
		handlers.NewSubscriberHandler,
		// QUIC server
		wireproviders.NewPublisherQUICServerFromConfig,
		wireproviders.NewSubscriberQUICServerFromConfig,
		// Application
		newApplication,
	)

	return application{}, nil
}
