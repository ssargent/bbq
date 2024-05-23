package config

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

type Config struct {
	Port     int
	Grpc     GrpcConfig
	Database DatabaseConfig
	Cache    CacheConfig
	Services ServiceConfig
}

type GrpcConfig struct {
	Port int `split_words:"true" default:"41337"`
}
type DatabaseConfig struct {
	Driver   string `split_words:"true" default:"postgres"`
	Username string `split_words:"true" default:"bbq"`
	Password string `required:"true" split_words:"true" `
	Server   string `split_words:"true" default:"localhost"`
	Name     string `split_words:"true" default:"bbq"`
	SSLMode  string `split_words:"true" default:"disable"`
	Port     int    `split_words:"true" default:"5432"`
}

type CacheConfig struct {
	DefaultExpiration time.Duration `split_words:"true" default:"5m"`
	DefaultCleanup    time.Duration `split_words:"true" default:"10m"`
}

type ServiceConfig struct {
	CollectorEnabled  bool `split_words:"true" default:"false"`
	ConsoleEnabled    bool `split_words:"true" default:"false"`
	SimulatorEnabled  bool `split_words:"true" default:"false"`
	PublicApiEnabled  bool `split_words:"true" default:"false"`
	PrivateApiEnabled bool `split_words:"true" default:"false"`
}

func (d *DatabaseConfig) Uri() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		d.Username,
		d.Password,
		net.JoinHostPort(d.Server, strconv.Itoa(d.Port)),
		d.Name,
		d.SSLMode)
}
