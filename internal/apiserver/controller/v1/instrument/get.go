package instrument

import (
	"gin_test/internal/apiserver/config"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	"net/http"
	"strconv"
	"strings"
)

func (i *InstrumentController) Get(c *gin.Context) {
	symbolType := c.Query("symbol_type")
	symbols := c.Query("symbols")

	ins, err := i.srv.Instruments().Get(c, symbolType, symbols)
	if err != nil {
		core.WriteResponse(c, err, config.Res{
			Status: http.StatusNotFound,
			Msg:    "error",
		})
		return
	}
	res := config.Res{
		Status: http.StatusOK,
		Msg:    "success",
		Data:   ins,
	}
	core.WriteResponse(c, nil, res)
}

func (i *InstrumentController) GetByParam(c *gin.Context) {
	instrumentIds := strings.Split(c.Query("instrument_ids"), ",")
	var List []int
	for _, i := range instrumentIds {
		intI, err := strconv.Atoi(i)
		if err != nil {
			continue
		}
		List = append(List, intI)
	}
	inss, err := i.srv.Instruments().List(c, List)

	if err != nil {
		core.WriteResponse(c, err, config.Res{
			Status: http.StatusNotFound,
			Msg:    "error",
		})
		return
	}
	res := config.Res{
		Status: http.StatusOK,
		Msg:    "success",
		Data:   inss,
	}

	core.WriteResponse(c, nil, res)
}
