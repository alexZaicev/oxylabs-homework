package main

import (
	"context"

	"go.uber.org/zap"
	"oxylabs-homework/internal/adapters/quickclient/handlers"
)

type application struct {
	logger            *zap.Logger
	subscriberHandler *handlers.SubscriberHandler
}

func newApplication(
	logger *zap.Logger,
	subscriberHandler *handlers.SubscriberHandler,
) application {
	return application{
		logger:            logger,
		subscriberHandler: subscriberHandler,
	}
}

func (a *application) Shutdown(_ context.Context) {
	a.subscriberHandler.Stop()
}
