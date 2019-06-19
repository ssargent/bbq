package infrastructure

import (
	"time"
)

//go:generate mockgen  -destination=./mocks/infrastructure.go  -self_package github.com/ssargent/bbq/bbq-apiserver/internal/infrastructure -package=mock_infrastructure github.com/ssargent/bbq/bbq-apiserver/internal/infrastructure CacheService

type CacheService interface {
	SetItem(key string, object interface{}, expiration time.Duration) error
	GetItem(key string, object interface{}) error
	RemoveItem(key string) error
}
