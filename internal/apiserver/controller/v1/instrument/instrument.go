package instrument

import (
	v1 "gin_test/internal/apiserver/service/v1"
	"gin_test/internal/apiserver/store"
)

type InstrumentController struct {
	srv v1.Service
}

func NewInstrumentController(store store.Factory) *InstrumentController {
	return &InstrumentController{
		srv: v1.NewService(store),
	}
}
