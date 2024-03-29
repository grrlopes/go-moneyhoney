package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/grrlopes/go-moneyhoney/src/infra/http/routers"
)

func main() {
	if os.Getenv("MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.Default()
	app := server.Group("/")

	routers.AuthCtrl(app)
	routers.UserCtrl(app)
	routers.MoneyCtrl(app)

	server.Run()
}
