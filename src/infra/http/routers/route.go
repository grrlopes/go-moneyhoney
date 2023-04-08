package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grrlopes/go-moneyhoney/src/infra/http/controllers"
	"github.com/grrlopes/go-moneyhoney/src/middleware"
)

func AuthCtrl(app gin.IRouter) {
	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "No valid endpoint provided!",
		})
	})

	app.POST("/login", controllers.Login())
}

func UserCtrl(app gin.IRouter) {
	app.POST("/createuser", controllers.CreateUser())
}

func MoneyCtrl(app gin.IRouter) {
	app.GET("/findall", middleware.AuthUserToken(), controllers.FindAll())
	app.GET("/findbyid", controllers.FindById())
	app.PUT("/update", controllers.Update())
	app.POST("/save", controllers.Save())
}
