package caching

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func SensorCacheKey(id uuid.UUID) string {
	return fmt.Sprintf("sensors_%s", id.String())
}

func DeviceCacheKey(id uuid.UUID) string {
	return fmt.Sprintf("devices_%s", id.String())
}

func SessionCacheKey(id uuid.UUID) string {
	return fmt.Sprintf("sessions_%s", id.String())
}

func SessionCacheDuration() time.Duration {
	return 8 * time.Hour
}
