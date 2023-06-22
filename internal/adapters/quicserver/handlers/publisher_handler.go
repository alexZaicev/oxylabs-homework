package handlers

import (
	"context"

	"github.com/quic-go/quic-go"
	"go.uber.org/zap"
)

type PublisherHandler struct {
	logger *zap.Logger
	msgCh  chan string
}

func NewPublisherHandler(logger *zap.Logger, msgCh chan string) *PublisherHandler {
	return &PublisherHandler{
		logger: logger,
		msgCh:  msgCh,
	}
}

func (h *PublisherHandler) HandleSession(conn quic.Connection) {
	defer conn.CloseWithError(quic.ApplicationErrorCode(quic.NoError), "")

	stream, err := conn.AcceptStream(context.Background())
	if err != nil {
		h.logger.With(zap.Any("error", err.Error())).Error("failed to accept stream")
		return
	}
	defer stream.Close()

	for {
		buffer := make([]byte, defaultBufferSize)
		n, err := stream.Read(buffer)
		if err != nil {
			h.logger.With(zap.Any("error", err.Error())).Error("failed to read from stream")
			return
		}

		request := string(buffer[:n])
		h.logger.With(zap.Any("request", request)).Info("received message")

		h.msgCh <- request
	}
}

func (h *PublisherHandler) Stop() {}
