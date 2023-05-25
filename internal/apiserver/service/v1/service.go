package v1

import "gin_test/internal/apiserver/store"

type Service interface {
	Instruments() InstrumentSrv
	Calendars() CalendarSrv
}

type service struct {
	store store.Factory
}

func NewService(store store.Factory) Service {
	return &service{
		store: store,
	}
}

func (s *service) Instruments() InstrumentSrv {
	return newInstruments(s)
}

func (s *service) Calendars() CalendarSrv {
	return newCalendar(s)
}
