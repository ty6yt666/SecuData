package apiserver

import (
	"gin_test/internal/apiserver/controller/v1/instrument"
	"gin_test/internal/apiserver/store/postgres"
	pkg_options "gin_test/pkg/options"
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

	// 没检查错误
	storeIns, _ := postgres.GetPostgresFactorOr(pkg_options.NewPostgresOptions())
	v1 := g.Group("/v1")
	{
		instrumentv1 := v1.Group("/instrument")
		{
			instrumentController := instrument.NewInstrumentController(storeIns)
			instrumentv1.GET(":instrument_id", instrumentController.Get)
			instrumentv1.GET("", instrumentController.GetByParam)
		}
	}
	return g
}
