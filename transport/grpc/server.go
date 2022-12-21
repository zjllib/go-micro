package grpc

import (
	"crypto/tls"
	"github.com/zjllib/go-micro/transport"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	listener   net.Listener
	grpcServer *grpc.Server
	secure     bool
	tls        *tls.Config
}

func (t *server) Listen(opts ...transport.Option) error {
	var options transport.Options
	for _, o := range opts {
		o(&options)
	}

	ln, err := net.Listen(string(transport.TCP), options.Addrs[0])
	if err != nil {
		return err
	}
	t.listener = ln
	t.grpcServer = grpc.NewServer()
	return t.grpcServer.Serve(ln)
}

func (s server) Close() error {
	s.grpcServer.Stop()
	return s.listener.Close()
}

func (s *server) Raw() interface{} {
	return s.grpcServer
}
