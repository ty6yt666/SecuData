package store

type Factory interface {
	Instruments() InstrumentStore
	Calendars() CalendarStore
	Close() error
}

var client Factory

// Client return the store client instance
func Client() Factory {
	return client
}
