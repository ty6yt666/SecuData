package v1

import (
	"context"
	"gin_test/internal/apiserver/service/models"
	"gin_test/internal/apiserver/store"
)

type CalendarSrv interface {
	Get(ctx context.Context, calendar *models.Calendar) (*models.Calendar, error)
	List(ctx context.Context, calendar *models.Calendar) (*models.CalendarList, error)
}

type calendarService struct {
	store store.Factory
}

// 绑定
var _ CalendarSrv = (*calendarService)(nil)

func newCalendar(srv *service) *calendarService {
	return &calendarService{store: srv.store}
}

func (i calendarService) Get(ctx context.Context, calendar *models.Calendar) (*models.Calendar, error) {
	calendar, err := i.store.Calendars().Get(ctx, calendar)
	if err != nil {
		return nil, err
	}
	return calendar, nil
}

func (i calendarService) List(ctx context.Context, calendar *models.Calendar) (*models.CalendarList, error) {
	calendars, err := i.store.Calendars().List(ctx, calendar)
	if err != nil {
		return nil, err
	}
	return calendars, nil
}
