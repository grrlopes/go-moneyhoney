package controller

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
	usecase      listall.InputBoundary = listall.NewFindAll(repositories)
)

func serverr(srv gin.Engine) gin.Engine {

	srv.GET("/test", func(c *gin.Context) {

		income := entity.Income{}
		income.SetAuthor("xxxxx")

		result, _ := usecase.Execute(income)

		c.JSON(http.StatusOK, gin.H{
			"message": "blabla",
			"data":    result,
		})
	})

  return srv
}
