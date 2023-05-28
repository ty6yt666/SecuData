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
	Port                  int
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
	// 目标样式：dsn = "host=127.0.0.1 port=5432 user=root password=root dbname=datamaster"
	dsn := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s`, // ?charset=utf8&parseTime=%t&loc=%s
		opts.Host,
		opts.Port,
		opts.Username,
		opts.Password,
		opts.Database,
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	configDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	configDB.SetMaxOpenConns(100)
	configDB.SetConnMaxLifetime(opts.MaxConnectionLifeTime)
	configDB.SetMaxIdleConns(opts.MaxIdleConnections)

	return db, nil
}
