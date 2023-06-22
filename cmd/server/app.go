package main

import (
	"context"

	"go.uber.org/zap"
	"oxylabs-homework/internal/adapters/quicserver"
)

type application struct {
	logger           *zap.Logger
	publisherServer  *quicserver.PublisherServer
	subscriberServer *quicserver.SubscriberServer
}

func newApplication(
	logger *zap.Logger,
	publisherServer *quicserver.PublisherServer,
	subscriberServer *quicserver.SubscriberServer,
) application {
	return application{
		logger:           logger,
		publisherServer:  publisherServer,
		subscriberServer: subscriberServer,
	}
}

func (a *application) Shutdown(_ context.Context) {
	a.logger.Info("shutting down publisher server")
	a.publisherServer.Stop()
	a.logger.Info("shutting down subscriber server")
	a.subscriberServer.Stop()
}
