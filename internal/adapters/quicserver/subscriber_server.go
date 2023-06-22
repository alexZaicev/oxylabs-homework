package quicserver

import (
	"crypto/tls"

	"go.uber.org/zap"
)

type SubscriberServer struct {
	server *Server
}

func NewSubscriberServer(port uint16, handler Handler, tlsConf *tls.Config) (*SubscriberServer, error) {
	server, err := New(port, handler, tlsConf)
	if err != nil {
		return nil, err
	}

	return &SubscriberServer{
		server: server,
	}, nil
}

func (s *SubscriberServer) Start(logger *zap.Logger) {
	s.server.Start(logger)
}

func (s *SubscriberServer) Stop() {
	s.server.Stop()
}
