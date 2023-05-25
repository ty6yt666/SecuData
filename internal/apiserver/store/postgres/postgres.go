package postgres

import (
	"gin_test/internal/apiserver/store"
	"gorm.io/gorm"
)

type datastore struct {
	db *gorm.DB
}

func (ds *datastore) Users() store.InstrumentStore {
	return newInstruments(ds)
}
