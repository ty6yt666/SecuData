package store

import (
	"context"
	"gin_test/internal/apiserver/service/models"
)

type CalendarStore interface {
	Get(ctx context.Context, calendar *models.Calendar) (*models.Calendar, error)
	List(ctx context.Context, calendar *models.Calendar) (*models.CalendarList, error)
}
