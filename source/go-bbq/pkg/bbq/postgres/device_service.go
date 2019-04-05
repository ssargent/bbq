package postgres

import "github.com/ssargent/go-bbq/pkg/bbq"

//	"database/sql"

// Ensure DialService implements wtf.DialService.
//var _ bbq.DeviceService = &DeviceService{}

// DialService represents a service for managing dials.
type DeviceService struct {
}

// Dial returns a dial by ID.
func (s *DeviceService) GetDevices(tenantName string) (*bbq.Device, error) {
	return nil, nil
}
