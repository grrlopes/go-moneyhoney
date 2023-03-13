package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grrlopes/go-moneyhoney/src/infra/http/controllers"
)

func AuthCtrl(app gin.IRouter) {
	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "No valid endpoint provided!",
		})
	})

	app.POST("/login", controllers.Login())
}
