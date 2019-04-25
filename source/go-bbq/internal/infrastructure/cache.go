package infrastructure

import (
	"time"

	"github.com/go-redis/cache"
	"github.com/ssargent/go-bbq/config"
)

type CacheService interface {
	SetItem(key string, object interface{}, expiration time.Duration) error
	GetItem(key string, object interface{}) error
}

type redisCacheService struct {
	config *config.Config
}

func NewRedisCacheService(config *config.Config) CacheService {
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
