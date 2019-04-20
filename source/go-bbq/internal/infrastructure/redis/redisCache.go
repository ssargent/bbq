package redis

import (
	"time"

	"github.com/ssargent/go-bbq/internal/infrastructure"
)

type redisCache struct {
}

func NewRedisCache() infrastructure.CacheService {
	return &redisCache{}
}

func (r *redisCache) GetItem(key string, object interface{}) error {
	return nil
}

func (r *redisCache) SetItem(key string, object interface{}, expiration time.Duration) error {
	return nil
}
