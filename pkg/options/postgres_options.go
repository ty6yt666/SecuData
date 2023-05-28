package options

import (
	"gin_test/pkg/db"
	"gorm.io/gorm"
	"time"
)

type PostgresOptions struct {
	Host                  string        `json:"host,omitempty"`
	Port                  int           `json:"port,omitempty"`
	Username              string        `json:"username,omitempty"`
	Password              string        `json:"-"`
	Database              string        `json:"database"`
	MaxIdleConnections    int           `json:"max-idle-connections,omitempty"`
	MaxOpenConnections    int           `json:"max-open-connections,omitempty"`
	MaxConnectionLifeTime time.Duration `json:"max-connection-life-time,omitempty"`
	LogLevel              int           `json:"log-level"`
}

// NewPostgresOptions create default instance
func NewPostgresOptions() *PostgresOptions {
	return &PostgresOptions{
		Host:                  "127.0.0.1",
		Port:                  5432,
		Username:              "root",
		Password:              "root",
		Database:              "public",
		MaxIdleConnections:    100,
		MaxOpenConnections:    100,
		MaxConnectionLifeTime: time.Duration(10) * time.Second,
		LogLevel:              1,
	}
}

func (o *PostgresOptions) NewClient() (*gorm.DB, error) {
	opts := &db.Options{
		Host:                  o.Host,
		Port:                  o.Port,
		Username:              o.Username,
		Password:              o.Password,
		Database:              o.Database,
		MaxIdleConnections:    o.MaxIdleConnections,
		MaxOpenConnections:    o.MaxOpenConnections,
		MaxConnectionLifeTime: o.MaxConnectionLifeTime,
		LogLevel:              o.LogLevel,
	}
	return db.New(opts)
}
