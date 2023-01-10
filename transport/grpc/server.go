package grpc

import (
	"crypto/tls"
	"github.com/zjllib/go-micro/transport"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	opts       *transport.Options
	listener   net.Listener
	grpcServer *grpc.Server
	secure     bool
	tls        *tls.Config
}

func (s *server) Listen(opts ...transport.Option) error {
	if s.opts == nil {
		s.opts = new(transport.Options)
	}
	for _, o := range opts {
		o(s.opts)
	}

	ln, err := net.Listen(string(transport.TCP), s.opts.Addrs[0])
	if err != nil {
		return err
	}
	s.listener = ln
	s.grpcServer = grpc.NewServer()
	return s.grpcServer.Serve(ln)
}

func (s server) Close() error {
	s.grpcServer.Stop()
	return s.listener.Close()
}

func (s *server) Raw() interface{} {
	return s.grpcServer
}
