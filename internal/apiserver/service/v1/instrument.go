package v1

import (
	"context"
	"gin_test/internal/apiserver/service/models"
	"gin_test/internal/apiserver/store"
)

type InstrumentSrv interface {
	Get(ctx context.Context, instrument int) (*models.Instrument, error)
	List(ctx context.Context, instrument []int) (*models.InstrumentList, error)
}

type instrumentService struct {
	store store.Factory
}

// 绑定
var _ InstrumentSrv = (*instrumentService)(nil)

func newInstruments(srv *service) *instrumentService {
	return &instrumentService{store: srv.store}
}

func (i instrumentService) Get(ctx context.Context, instrumentId int) (*models.Instrument, error) {
	ins, err := i.store.Instruments().Get(ctx, instrumentId)
	if err != nil {
		return nil, err
	}
	return ins, nil
}

func (i instrumentService) List(ctx context.Context, instrumentIds []int) (*models.InstrumentList, error) {
	inses, err := i.store.Instruments().List(ctx, instrumentIds)
	if err != nil {
		return nil, err
	}
	return inses, nil
}
