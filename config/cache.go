package config

import (
	"os"

	"github.com/goweb4/services/cache"
	"github.com/goweb4/services/cache/redis"
)

// Cache func
func Cache() cache.Cache {
	return redis.Connect(
		os.Getenv("REDIS_ADDR"),
		os.Getenv("REDIS_PASSWORD"),
		0,
	)
}
