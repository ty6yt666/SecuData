package main

import (
	"gin_test/internal/ws"
	"gin_test/internal/ws/config"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func main() {
	iniFilePath := filepath.Join("deploy/apiserver", "dev.ini")
	g := gin.Default()
	config.Init(iniFilePath) // os.Getenv("ENVIRONMENT")
	ws.InitRouter(g)
	g.Run()
}
