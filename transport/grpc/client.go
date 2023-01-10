package grpc

import (
	"crypto/tls"
	"github.com/zjllib/go-micro/transport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type client struct {
	opts   *transport.Options
	conn   *grpc.ClientConn
	secure bool
	tls    *tls.Config
}

func (c *client) Raw() interface{} {
	return c.conn
}

func (c *client) Dial(opts ...transport.Option) error {
	if c.opts == nil {
		c.opts = new(transport.Options)
	}
	for _, o := range opts {
		o(c.opts)
	}
	grpcDialoptions := []grpc.DialOption{
		grpc.WithTimeout(c.opts.Timeout),
	}

	if c.opts.Secure || c.opts.TLSConfig != nil {
		config := c.opts.TLSConfig
		if config == nil {
			config = &tls.Config{
				InsecureSkipVerify: true,
			}
		}
		creds := credentials.NewTLS(config)
		grpcDialoptions = append(grpcDialoptions, grpc.WithTransportCredentials(creds))
	} else {
		grpcDialoptions = append(grpcDialoptions, grpc.WithInsecure())
	}
	conn, err := grpc.Dial(c.opts.Addrs[0], grpcDialoptions...)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *client) Close() error {
	return c.conn.Close()
}
