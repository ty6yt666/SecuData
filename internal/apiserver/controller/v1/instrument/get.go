package instrument

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	"strconv"
)

func (i *InstrumentController) Get(c *gin.Context) {
	instrumentId, err := strconv.Atoi(c.Param("instrument_id"))
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	ins, err := i.srv.Instruments().Get(c, instrumentId)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, ins)
}
