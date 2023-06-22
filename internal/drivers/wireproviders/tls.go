package wireproviders

import (
	"crypto/tls"
	"fmt"
)

func NewTLSConfig(conf TLSConfig) (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(conf.X509, conf.Key)
	if err != nil {
		return nil, fmt.Errorf("failed to load x509 certificate: %s", err)
	}

	return &tls.Config{
		Certificates:       []tls.Certificate{cert},
		NextProtos:         []string{"oxylab-homework"},
		InsecureSkipVerify: true,
	}, nil
}
