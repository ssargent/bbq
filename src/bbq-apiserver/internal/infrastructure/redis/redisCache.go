package redis

import (
	"time"

	"github.com/go-redis/cache"
	"github.com/ssargent/bbq/bbq-apiserver/config"
	"github.com/ssargent/bbq/bbq-apiserver/internal/infrastructure"
)

type redisCacheService struct {
	config *config.Config
}

func NewRedisCacheService(config *config.Config) infrastructure.CacheService {
	return &redisCacheService{config: config}
}

func (r redisCacheService) SetItem(key string, object interface{}, expiration time.Duration) error {
	return r.config.Cache.Set(&cache.Item{
		Key:        key,
		Object:     object,
		Expiration: expiration,
	})
}

func (r redisCacheService) GetItem(key string, object interface{}) error {
	err := r.config.Cache.Get(key, &object)

	if err != nil {
		return err
	}

	return nil
}

func (r redisCacheService) RemoveItem(key string) error {
	err := r.config.Cache.Delete(key)

	if err != nil {
		return err
	}

	return nil
}
