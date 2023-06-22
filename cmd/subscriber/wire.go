//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/google/wire"
	"oxylabs-homework/internal/adapters/quickclient/handlers"
	"oxylabs-homework/internal/drivers/wireproviders"
)

func initialize(ctx context.Context) (application, error) {
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
		// QUIC client
		wireproviders.NewSubscriberQUICClientFromConfig,
		// Handler
		handlers.NewSubscriberHandler,
		// Application
		newApplication,
	)

	return application{}, nil
}
