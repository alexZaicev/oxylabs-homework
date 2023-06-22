package main

import (
	"crypto/tls"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/lucas-clemente/quic-go"
)

type InboundServer struct {
	Listener       quic.Listener
	MessageChannel chan string
}

type OutboundServer struct {
	Listener quic.Listener
}

func NewInboundServer() (*InboundServer, error) {
	// Load the TLS certificate and private key
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		return nil, err
	}

	// Create a TLS configuration
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	// Create an inbound QUIC listener with the TLS configuration
	listener, err := quic.ListenAddr("localhost:4646", tlsConfig, nil)
	if err != nil {
		return nil, err
	}

	return &InboundServer{
		Listener:       listener,
		MessageChannel: make(chan string),
	}, nil
}

func (s *InboundServer) AcceptConnections() {
	for {
		session, err := s.Listener.Accept()
		if err != nil {
			log.Println("Error accepting session:", err)
			return
		}

		go s.HandleSession(session)
	}
}

func (s *InboundServer) HandleSession(session quic.Session) {
	stream, err := session.AcceptStream()
	if err != nil {
		log.Println("Error accepting stream:", err)
		session.Close()
		return
	}

	log.Println("Inbound session started.")

	for {
		buf := make([]byte, 1024)
		n, err := stream.Read(buf)
		if err != nil {
			log.Println("Error reading from stream:", err)
			stream.Close()
			session.Close()
			return
		}

		message := string(buf[:n])
		s.MessageChannel <- message
	}
}

func NewOutboundServer() (*OutboundServer, error) {
	// Load the TLS certificate and private key
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		return nil, err
	}

	// Create a TLS configuration
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	// Create an outbound QUIC listener with the TLS configuration
	listener, err := quic.ListenAddr("localhost:4747", tlsConfig, nil)
	if err != nil {
		return nil, err
	}

	return &OutboundServer{
		Listener: listener,
	}, nil
}

func (s *OutboundServer) AcceptConnections(messageChannel <-chan string) {
	for {
		session, err := s.Listener.Accept()
		if err != nil {
			log.Println("Error accepting session:", err)
			return
		}

		go s.HandleSession(session, messageChannel)
	}
}

func (s *OutboundServer) HandleSession(session quic.Session, messageChannel <-chan string) {
	stream, err := session.OpenStreamSync()
	if err != nil {
		log.Println("Error opening stream:", err)
		session.Close()
		return
	}

	log.Println("Outbound session started.")

	for {
		message := <-messageChannel

		_, err := stream.Write([]byte(message))
		if err != nil {
			log.Println("Error writing to stream:", err)
			stream.Close()
			session.Close()
			return
		}
	}
}

func main() {
	inboundServer, err := NewInboundServer()
	if err != nil {
		log.Fatal("Error creating inbound server:", err)
	}
	defer inboundServer.Listener.Close()

	outboundServer, err := NewOutboundServer()
	if err != nil {
		log.Fatal("Error creating outbound server:", err)
	}
	defer outboundServer.Listener.Close()

	go inboundServer.AcceptConnections()
	go outboundServer.AcceptConnections(inboundServer.MessageChannel)

	// Handle SIGTERM signal for graceful shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	log.Println("Server started.")

	<-signalChan

	log.Println("Received SIGTERM signal. Gracefully shutting down...")
}