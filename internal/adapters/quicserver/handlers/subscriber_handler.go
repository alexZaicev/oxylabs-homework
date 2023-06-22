package handlers

import (
	"context"

	"github.com/quic-go/quic-go"
	"go.uber.org/zap"
)

type SubscriberHandler struct {
	logger   *zap.Logger
	msgCh    <-chan string
	shutdown chan struct{}
}

func NewSubscriberHandler(logger *zap.Logger, msgCh <-chan string) *SubscriberHandler {
	return &SubscriberHandler{
		logger:   logger,
		msgCh:    msgCh,
		shutdown: make(chan struct{}),
	}
}

func (h *SubscriberHandler) HandleSession(conn quic.Connection) {
	defer conn.CloseWithError(quic.ApplicationErrorCode(quic.NoError), "")

	stream, err := conn.OpenStreamSync(context.Background())
	if err != nil {
		h.logger.With(zap.Any("error", err.Error())).Error("failed to accept stream")
		return
	}
	defer stream.Close()

	for {
		select {
		case <-h.shutdown:
			return
		case msg := <-h.msgCh:
			if _, err := stream.Write([]byte(msg)); err != nil {
				h.logger.With(zap.Any("error", err.Error())).Error("failed to write to stream")
				return
				h.logger.With(zap.Any("request", msg)).Info("message send")
			}
		}
	}
}

func (h *SubscriberHandler) Stop() {
	close(h.shutdown)
}
