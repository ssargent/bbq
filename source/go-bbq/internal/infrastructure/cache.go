package infrastructure

import (
	"time"
)

type CacheService interface {
	SetItem(key string, object interface{}, expiration time.Duration) error
	GetItem(key string, object interface{}) error
}
