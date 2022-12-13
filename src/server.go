package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	server.GET("/test", func(c *gin.Context) {

		income := entity.Income{}
		income.SetAuthor("xxxxx")

		c.JSON(http.StatusOK, gin.H{
			"message": "blabla",
		})
	})

	server.Run()
}
