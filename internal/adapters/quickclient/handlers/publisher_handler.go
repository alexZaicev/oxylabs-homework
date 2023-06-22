package handlers

import (
	"time"

	"go.uber.org/zap"
	"oxylabs-homework/internal/adapters/quickclient"
)

type PublisherHandler struct {
	logger   *zap.Logger
	client   *quickclient.PublisherClient
	shutdown chan struct{}
}

func NewPublisherHandler(logger *zap.Logger, client *quickclient.PublisherClient) *PublisherHandler {
	return &PublisherHandler{
		logger:   logger,
		client:   client,
		shutdown: make(chan struct{}),
	}
}

func (h *PublisherHandler) Serve() {
	// Start ticker to send messages
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			h.logger.Info("sending message")
			if err := h.client.SendMessage("Hello world"); err != nil {
				h.logger.With(zap.Any("error", err.Error())).Error("failed to send message")
			}
		case <-h.shutdown:
			return
		}
	}
}

func (h *PublisherHandler) Stop() {
	close(h.shutdown)
}
