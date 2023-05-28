package postgres

import (
	"fmt"
	"gin_test/internal/apiserver/store"
	"gin_test/pkg/db"
	pkg_options "gin_test/pkg/options"
	"gorm.io/gorm"
	"sync"
)

type datastore struct {
	db *gorm.DB
}

func (ds *datastore) Calendars() store.CalendarStore {
	//TODO implement me
	panic("implement me")
}

func (ds *datastore) Close() error {
	//TODO implement me
	panic("implement me")
}

func (ds *datastore) Instruments() store.InstrumentStore {
	return newInstruments(ds)
}

var (
	postgresFactory store.Factory
	once            sync.Once
)

func GetPostgresFactorOr(opts *pkg_options.PostgresOptions) (store.Factory, error) {
	if opts == nil && postgresFactory == nil {
		return nil, fmt.Errorf("failed to get postgres store fatory")
	}

	var err error
	var dbIns *gorm.DB
	once.Do(func() {
		options := &db.Options{
			Host:                  opts.Host,
			Port:                  opts.Port,
			Username:              opts.Username,
			Password:              opts.Password,
			Database:              opts.Database,
			MaxIdleConnections:    opts.MaxIdleConnections,
			MaxOpenConnections:    opts.MaxOpenConnections,
			MaxConnectionLifeTime: opts.MaxConnectionLifeTime,
			LogLevel:              opts.LogLevel,
			//Logger:                logger.New(opts.LogLevel),
		}
		dbIns, err = db.New(options)
		postgresFactory = &datastore{dbIns}
	})

	if postgresFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql store fatory, postgresFactory: %+v, error: %w", postgresFactory, err)
	}
	return postgresFactory, nil
}
