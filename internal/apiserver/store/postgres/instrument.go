package postgres

import (
	"context"
	"errors"
	"gin_test/internal/apiserver/service/models"
	"gorm.io/gorm"
)

type instruments struct {
	db *gorm.DB
}

// newInstruments 包内调用
func newInstruments(ds *datastore) *instruments {
	return &instruments{ds.db}
}

func (ins *instruments) Get(ctx context.Context, insObj *models.Instrument) (*models.Instrument, error) {
	instrument := &models.Instrument{}
	err := ins.db.Where("instrument_id = ?", insObj.InstrumentId).First(&instrument).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// todo: error type
			return nil, err
		}
		return nil, err
	}
	return instrument, nil
}

func (ins *instruments) List(ctx context.Context, insIdList []int) (*models.InstrumentList, error) {
	instrumentList := &models.InstrumentList{}

	d := ins.db.Where("instrument_id in ?", insIdList).Order("instrument_id").Find(&instrumentList.Items)
	return instrumentList, d.Error
}
