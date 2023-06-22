package quicserver

import (
	"context"
	"crypto/tls"
	"fmt"
	"sync"

	"github.com/quic-go/quic-go"
	"go.uber.org/zap"
)

type Server struct {
	listener *quic.Listener
	handler  Handler
	shutdown chan struct{}
	wg       sync.WaitGroup
}

func New(port uint16, handler Handler, tlsConf *tls.Config) (*Server, error) {
	if handler == nil {
		return nil, fmt.Errorf("invalid parameter handler: must not be nil")
	}

	if tlsConf == nil {
		return nil, fmt.Errorf("invalid parameter tlsConf: must not be nil")
	}

	ln, err := quic.ListenAddr(
		fmt.Sprintf(":%d", port),
		tlsConf,
		&quic.Config{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to start listening on port [:%d]: %s", port, err)
	}
	return &Server{
		listener: ln,
		handler:  handler,
		shutdown: make(chan struct{}),
	}, nil
}

func (s *Server) Start(logger *zap.Logger) {
	s.wg.Add(1)
	go s.acceptConn(logger)

	// Wait for shutdown signal
	<-s.shutdown

	// Close the listener to stop accepting new connections
	if err := s.listener.Close(); err != nil {
		logger.With(zap.Any("error", err.Error())).Error("failed to close quic listener")
	}

	// Wait for all connections to be processed
	s.wg.Wait()
}

func (s *Server) acceptConn(logger *zap.Logger) {
	defer s.wg.Done()

	for {
		select {
		case <-s.shutdown:
			return
		default:
			ctx := context.Background()

			conn, err := s.listener.Accept(ctx)
			if err != nil {
				logger.With(zap.Any("error", err.Error())).Error("failed to accept incoming quic connection")
				continue
			}

			s.wg.Add(1)
			go s.handleSession(conn)
		}
	}
}

func (s *Server) handleSession(conn quic.Connection) {
	defer s.wg.Done()
	s.handler.HandleSession(conn)
}

func (s *Server) Stop() {
	close(s.shutdown)
	s.handler.Stop()
	s.wg.Wait()
}
