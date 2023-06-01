package postgres

import (
	"context"
	"errors"
	"gin_test/internal/apiserver/service/models"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type instruments struct {
	db *gorm.DB
}

// newInstruments 包内调用
func newInstruments(ds *datastore) *instruments {
	return &instruments{ds.db}
}

func (ins *instruments) Get(ctx context.Context, symbolType string, symbol string) ([]*models.Instrument, error) {
	var instrumentSlice []*models.Instrument
	var querySymbols []interface{}
	if symbolType == "instrument_id" {
		symbolSlice := strings.Split(symbol, ",")
		for _, s := range symbolSlice {
			sInt, err := strconv.Atoi(s)
			if err != nil {
				continue
			}
			querySymbols = append(querySymbols, sInt)
		}
	} else {
		querySymbol := strings.Split(symbol, ",")
		for _, perSymbol := range querySymbol {
			querySymbols = append(querySymbols, perSymbol)
		}
	}
	err := ins.db.Where(symbolType+" in ?", querySymbols).Find(&instrumentSlice).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// todo: error type
			return nil, err
		}
		return nil, err
	}
	return instrumentSlice, nil
}

func (ins *instruments) List(ctx context.Context, InstrumentIds []int) (*models.InstrumentList, error) {
	instrumentList := &models.InstrumentList{}

	d := ins.db.Where("instrument_id in ?", InstrumentIds).Order("instrument_id").Find(&instrumentList.Items)
	return instrumentList, d.Error
}
