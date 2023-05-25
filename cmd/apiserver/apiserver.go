package main

import (
	"gin_test/internal/apiserver"
	"github.com/gin-gonic/gin"
)

func main() {
	//apiserver.NewApp("datamaster")
	//g := &gin.Engine{}
	g := gin.Default()
	apiserver.InitRouter(g)
	g.Run()
}
