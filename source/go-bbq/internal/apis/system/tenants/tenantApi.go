package tenants

import (
	"github.com/google/uuid"
)

// Tenant is the representation of a sys tenant
type Tenant struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	URLKey string `json:"urlkey"`
	IsEnabled bool `json:"isenabled"` 
}
