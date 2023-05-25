package apiserver

import (
	"gin_test/internal/apiserver/controller/v1/instrument"
	"gin_test/internal/apiserver/store"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	"github.com/marmotedu/errors"
)

func InitRouter(g *gin.Engine) {
	installMiddleware(g)
	installController(g)
}

func installMiddleware(g *gin.Engine) {}

func installController(g *gin.Engine) *gin.Engine {
	g.NoRoute(func(context *gin.Context) {
		core.WriteResponse(context, errors.WithCode(500, "Page not found."), nil)
	})

	storeIns := store.Client()
	v1 := g.Group("/v1")
	{
		instrumentv1 := v1.Group("/instrument")
		{
			instrumentController := instrument.NewInstrumentController(storeIns)

			//instrumentv1.GET("", func(context *gin.Context) {
			//	core.WriteResponse(context, nil, "你看我像不像数据")
			//})
			instrumentv1.GET(":instrument_id", instrumentController.Get)
		}
	}
	return g
}
