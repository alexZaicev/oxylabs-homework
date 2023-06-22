package handlers

import (
	"go.uber.org/zap"
	"oxylabs-homework/internal/adapters/quickclient"
)

type SubscriberHandler struct {
	logger   *zap.Logger
	client   *quickclient.SubscriberClient
	shutdown chan struct{}
}

func NewSubscriberHandler(logger *zap.Logger, client *quickclient.SubscriberClient) *SubscriberHandler {
	return &SubscriberHandler{
		logger:   logger,
		client:   client,
		shutdown: make(chan struct{}),
	}
}

func (h *SubscriberHandler) Serve() {
	for {
		msg, err := h.client.ReadMessage()
		if err != nil {
			h.logger.With(zap.Any("error", err.Error())).Error("failed to read message")
			return
		}

		// continue with loop cycle on receiving an empty message
		if msg == "" {
			continue
		}

		h.logger.Info("received message: " + msg)
	}
}

func (h *SubscriberHandler) Stop() {
	close(h.shutdown)
}
