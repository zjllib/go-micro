// Package mucp provides an mucp server
package mucp

import (
	"github.com/zjllib/go-micro/cmd"
	"github.com/zjllib/go-micro/server"
)

func init() {
	cmd.DefaultServers["mucp"] = NewServer
}

// NewServer returns a micro server interface
func NewServer(opts ...server.Option) server.Server {
	return server.NewServer(opts...)
}
