package http

import (
	"github.com/zjllib/go-micro/registry"
)

type Options struct {
	registry.IRegistry.Registry
}

type Option func(*Options)

func WithRegistry(r registry.IRegistry) Option {
	return func(o *Options) {
		o.Registry = r
	}
}
