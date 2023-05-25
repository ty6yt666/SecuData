package store

import (
	"context"
	"gin_test/internal/apiserver/service/models"
)

type InstrumentStore interface {
	Get(ctx context.Context, instrument int) (*models.Instrument, error)
	List(ctx context.Context, instrument []int) (*models.InstrumentList, error)
}
