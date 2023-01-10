// Package grpc provides a grpc transport
package grpc

import (
	"crypto/tls"
	"net"

	"github.com/zjllib/go-micro/transport"
	maddr "github.com/zjllib/go-micro/util/addr"
	mls "github.com/zjllib/go-micro/util/tls"
)

type grpcTransport struct {
	opts   transport.Options
	server transport.IServer
	client transport.IClient
}

type grpcTransportListener struct {
	listener net.Listener
	secure   bool
	tls      *tls.Config
}

func getTLSConfig(addr string) (*tls.Config, error) {
	hosts := []string{addr}

	// check if its a valid host:port
	if host, _, err := net.SplitHostPort(addr); err == nil {
		if len(host) == 0 {
			hosts = maddr.IPs()
		} else {
			hosts = []string{host}
		}
	}

	// generate a certificate
	cert, err := mls.Certificate(hosts...)
	if err != nil {
		return nil, err
	}

	return &tls.Config{Certificates: []tls.Certificate{cert}}, nil
}

func (t *grpcTransport) init(opts ...transport.Option) error {
	for _, o := range opts {
		o(&t.opts)
	}
	t.server = &server{
		opts: &t.opts,
	}
	t.client = &client{
		opts: &t.opts,
	}
	return nil
}

func (t *grpcTransport) Server() transport.IServer {
	return t.server
}
func (t *grpcTransport) Client() transport.IClient {
	return t.client
}

func NewTransport(opts ...transport.Option) transport.ITransport {
	t := new(grpcTransport)
	t.init(opts...)
	return t
}
