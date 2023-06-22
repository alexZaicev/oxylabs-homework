package quicserver

import (
	"github.com/quic-go/quic-go"
)

type Handler interface {
	HandleSession(connection quic.Connection)
	Stop()
}
