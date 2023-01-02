package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/grrlopes/go-moneyhoney/src/infra/controller"
)

func main() {
	if os.Getenv("MODE") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.Default()
	app := server.Group("/")

	controller.MoneyCtrl(app)

	server.Run()
}
