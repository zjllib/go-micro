package cache

import (
	"time"

	"github.com/zjllib/go-micro/registry/cache"
)

// WithTTL sets the cache TTL
func WithTTL(t time.Duration) cache.Option {
	return cache.WithTTL(t)
}
