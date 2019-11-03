package bbq

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

//go:generate mockgen  -destination=./mocks/bbq.go -package=mock_bbq github.com/ssargent/bbq/bbq-apiserver/bbq SessionRepository,DeviceRepository,MonitorRepository,SubjectRepository,SessionService,DeviceService,MonitorService,SubjectService

//Device is
type Device struct {
	ID          int       `json:"id"`
	Uid         uuid.UUID `json:"uid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	TenantID    uuid.UUID `json:"tenantid"`
}

//Subject is the thing thats being cooked.
type Subject struct {
	ID          int       `json:"id"`
	Uid         uuid.UUID `json:"uid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	TenantID    uuid.UUID `json:"tenantid"`
}

//Monitor is
type Monitor struct {
	ID          int       `json:"id"`
	Uid         uuid.UUID `json:"uid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Address     string    `json:"address"`
	TenantID    uuid.UUID `json:"tenantid"`
}

// Session is
type Session struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Subject     string      `json:"subject"`
	Type        string      `json:"type"`
	Weight      float64     `json:"weight"`
	Device      string      `json:"device"`
	Monitor     string      `json:"monitor"`
	StartTime   time.Time   `json:"starttime"`
	EndTime     pq.NullTime `json:"endtime"`
	TenantID    uuid.UUID   `json:"tenantid"`
	UID         uuid.UUID   `json:"uid"`
}

type SessionRecord struct {
	ID          int         `json:"id"`
	DeviceID    int         `json:"deviceid"`
	DeviceUID   uuid.UUID   `json:"deviceuid"`
	MonitorID   int         `json:"monitorid"`
	MonitorUID  uuid.UUID   `json:"monitoruid"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	StartTime   time.Time   `json:"starttime"`
	SubjectID   int         `json:"subjectid"`
	SubjectUID  uuid.UUID   `json:"subjectuid"`
	Weight      float64     `json:"weight"`
	TenantID    uuid.UUID   `json:"tenantid"`
	UID         uuid.UUID   `json:"uid"`
	EndTime     pq.NullTime `json:"endtime"`
}

// DeviceService is the service for devices
type DeviceService interface {
	GetDevices(tenantID uuid.UUID) ([]Device, error)
	GetDeviceByName(tenantID uuid.UUID, deviceName string) (Device, error)
	GetDeviceByID(tenantId uuid.UUID, id uuid.UUID) (Device, error)
	CreateDevice(tenantID uuid.UUID, newDevice Device) (Device, error)
	UpdateDevice(tenantID uuid.UUID, existingDevice Device) (Device, error)
	DeleteDevice(tenantID uuid.UUID, existingDevice Device) error
}

// DeviceRepository is the repo for Devices
type DeviceRepository interface {
	GetByTenantID(tenantID uuid.UUID) ([]Device, error)
	GetByID(tenantID uuid.UUID, id uuid.UUID) (Device, error)
	GetByName(tenantID uuid.UUID, deviceName string) (Device, error)
	Create(newDevice Device) (Device, error)
	Update(device Device) (Device, error)
	Delete(device Device) error
}

// SessionService is the service for Sessions
type SessionService interface {
	GetSessions(tenantID uuid.UUID) ([]Session, error)
	GetSessionByID(tenantID uuid.UUID, id uuid.UUID) (Session, error)
	GetSessionByMonitorAddress(tenantID uuid.UUID, address string) (Session, error)
	CreateSession(tenantID uuid.UUID, entity Session) (Session, error)
	UpdateSession(tenantID uuid.UUID, entity Session) (Session, error)
	DeleteSession(tenantID uuid.UUID, entity Session) error
}

// SessionRepository is the repo for Sessions
type SessionRepository interface {
	GetByTenantID(tenantID uuid.UUID) ([]Session, error)
	GetByID(tenantID uuid.UUID, id uuid.UUID) (Session, error)
	GetByMonitorAddress(tenantID uuid.UUID, address string) (Session, error)
	Create(tenantID uuid.UUID, entity SessionRecord) (SessionRecord, error)
	Update(tenantID uuid.UUID, entity SessionRecord) (SessionRecord, error)
	Delete(tenantID uuid.UUID, entity Session) error
}

// MonitorService is the service for monitors
type MonitorService interface {
	GetMonitors(tenantID uuid.UUID) ([]Monitor, error)
	GetMonitorByID(tenantID uuid.UUID, monitorId uuid.UUID) (Monitor, error)
	GetMonitorByName(tenantID uuid.UUID, name string) (Monitor, error)
	GetMonitorByAddress(tenantID uuid.UUID, address string) (Monitor, error)
	CreateMonitor(tenantID uuid.UUID, entity Monitor) (Monitor, error)
	UpdateMonitor(tenantID uuid.UUID, entity Monitor) (Monitor, error)
	DeleteMonitor(tenantID uuid.UUID, entity Monitor) error
}

// MonitorRepository is the repo for Devices
type MonitorRepository interface {
	GetByTenantID(tenantID uuid.UUID) ([]Monitor, error)
	GetByID(tenantID uuid.UUID, monitorId uuid.UUID) (Monitor, error)
	GetByName(tenantID uuid.UUID, name string) (Monitor, error)
	GetByAddress(tenantID uuid.UUID, address string) (Monitor, error)
	Create(entity Monitor) (Monitor, error)
	Update(entity Monitor) (Monitor, error)
	Delete(entity Monitor) error
}

type SubjectService interface {
	GetOrCreateSubject(tenantID uuid.UUID, name string, description string) (Subject, error)
	GetSubjectByID(tenantID uuid.UUID, subjectId uuid.UUID) (Subject, error)
}

type SubjectRepository interface {
	GetByID(subjectID uuid.UUID) (Subject, error)
	GetByName(tenantID uuid.UUID, name string) (Subject, error)
	Create(entity Subject) (Subject, error)
	Update(entity Subject) (Subject, error)
	Delete(entity Subject) error
}

type BBQUnitOfWork struct {
	Monitor MonitorRepository
	Device  DeviceRepository
	Session SessionRepository
	Subject SubjectRepository
}
