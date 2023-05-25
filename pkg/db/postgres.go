package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type Options struct {
	Host                  string
	Username              string
	Password              string
	Database              string
	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxConnectionLifeTime time.Duration
	LogLevel              int
	Logger                logger.Interface
}

// New create a new gorm db instance with the options
func New(opts *Options) (*gorm.DB, error) {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		opts.Username,
		opts.Password,
		opts.Host,
		opts.Database,
		true,
		"Local",
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: opts.Logger,
	})
	if err != nil {
		return nil, err
	}

	configDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	configDB.SetMaxOpenConns(opts.MaxOpenConnections)
	configDB.SetConnMaxLifetime(opts.MaxConnectionLifeTime)
	configDB.SetMaxIdleConns(opts.MaxIdleConnections)

	return db, nil
}
