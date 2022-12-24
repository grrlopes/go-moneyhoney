package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grrlopes/go-moneyhoney/src/application/usecase/listall"
	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
	"github.com/grrlopes/go-moneyhoney/src/domain/repository"
	"github.com/grrlopes/go-moneyhoney/src/infra/repositories/couchdb"
)

var (
	repositories repository.IMoneyRepo = couchdb.NewMoneyRepository()
	usecase      listall.Input         = listall.NewFindAll(repositories)
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	server.GET("/test", func(c *gin.Context) {

		income := entity.Income{}
		income.SetAuthor("xxxxx")

		usecase.Execute(income)

		c.JSON(http.StatusOK, gin.H{
			"message": "blabla",
		})
	})

	server.Run()
}
