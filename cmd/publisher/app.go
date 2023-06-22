package main

import (
	"context"

	"go.uber.org/zap"
	"oxylabs-homework/internal/adapters/quickclient/handlers"
)

type application struct {
	logger           *zap.Logger
	publisherHandler *handlers.PublisherHandler
}

func newApplication(
	logger *zap.Logger,
	publisherHandler *handlers.PublisherHandler,
) application {
	return application{
		logger:           logger,
		publisherHandler: publisherHandler,
	}
}

func (a *application) Shutdown(_ context.Context) {
	a.publisherHandler.Stop()
}
