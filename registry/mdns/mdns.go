// Package mdns provides a multicast dns registry
package mdns

import (
	"github.com/zjllib/go-micro/cmd"
	"github.com/zjllib/go-micro/registry"
)

func init() {
	cmd.DefaultRegistries["mdns"] = NewRegistry
}

// NewRegistry returns a new mdns registry
func NewRegistry(opts ...registry.Option) registry.IRegistry {
	return registry.NewRegistry(opts...)
}
