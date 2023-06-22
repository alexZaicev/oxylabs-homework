package quicserver

import (
	"crypto/tls"

	"go.uber.org/zap"
)

type PublisherServer struct {
	server *Server
}

func NewPublisherServer(port uint16, handler Handler, tlsConf *tls.Config) (*PublisherServer, error) {
	server, err := New(port, handler, tlsConf)
	if err != nil {
		return nil, err
	}

	return &PublisherServer{
		server: server,
	}, nil
}

func (s *PublisherServer) Start(logger *zap.Logger) {
	s.server.Start(logger)
}

func (s *PublisherServer) Stop() {
	s.server.Stop()
}
