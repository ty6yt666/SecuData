package store

import (
	"context"
	"gin_test/internal/apiserver/service/models"
)

type InstrumentStore interface {
	Get(ctx context.Context, symbolType string, symbol string) ([]*models.Instrument, error)
	List(ctx context.Context, instrumentIds []int) (*models.InstrumentList, error)
}
