package options

import (
	"gin_test/pkg/db"
	"gorm.io/gorm"
	"time"
)

// PostgresOptions defines options for postgres database.
type PostgresOptions struct {
	Host                  string        `json:"host,omitempty"`
	Username              string        `json:"username,omitempty"`
	Password              string        `json:"-"`
	Database              string        `json:"database"`
	MaxIdleConnections    int           `json:"max-idle-connections,omitempty"`
	MaxOpenConnections    int           `json:"max-open-connections,omitempty"`
	MaxConnectionLifeTime time.Duration `json:"max-connection-life-time,omitempty"`
	LogLevel              int           `json:"log-level"`
}

// NewPostgresOptions create a default value instance.
func NewPostgresOptions() *PostgresOptions {
	return &PostgresOptions{
		Host:                  "localhost:3306",
		Username:              "",
		Password:              "",
		Database:              "",
		MaxIdleConnections:    100,
		MaxOpenConnections:    100,
		MaxConnectionLifeTime: time.Duration(10) * time.Second,
		LogLevel:              1,
	}
}

// NewClient create postgres obj with the config
func (o *PostgresOptions) NewClient() (*gorm.DB, error) {
	opts := &db.Options{
		Host:                  o.Host,
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
