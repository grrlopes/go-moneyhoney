package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/grrlopes/go-moneyhoney/src/infra/http/routers"
)

func main() {
	if os.Getenv("MODE") == "Prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.Default()
	app := server.Group("/")

	routers.AuthCtrl(app)

	server.Run()
}
