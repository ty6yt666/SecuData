package apiserver

import (
	"github.com/gin-gonic/gin"
)

type apiServer struct {
	APIServer *gin.Engine
}

//func createAPIServer(cfg *config.Config) (*apiServer, error) {
//	return apiServer{}
//}
