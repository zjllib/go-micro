// Package cache provides a registry cache
package cache

import (
	"github.com/zjllib/go-micro/registry"
	"github.com/zjllib/go-micro/registry/cache"
)

// New returns a new cache
func New(r registry.IRegistry, opts ...cache.Option) cache.Cache {
	return cache.New(r, opts...)
}
