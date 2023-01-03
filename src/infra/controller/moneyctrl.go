package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grrlopes/go-moneyhoney/src/application/usecase/listall"
	"github.com/grrlopes/go-moneyhoney/src/domain/repository"
	"github.com/grrlopes/go-moneyhoney/src/infra/presenters"
	"github.com/grrlopes/go-moneyhoney/src/infra/repositories/couchdb"
)

var (
	repositories    repository.IMoneyRepo = couchdb.NewMoneyRepository()
	usecase_listall listall.InputBoundary = listall.NewFindAll(repositories)
)

func MoneyCtrl(app gin.IRouter) {
	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "No valid endpoint provided!",
		})
	})

	app.GET("/findall", func(c *gin.Context) {
		result, err := usecase_listall.Execute()

		if err != nil {
			error := presenters.MoneyErrorResponse(result)
			c.JSON(http.StatusInternalServerError, error)
			return
		}

		data := presenters.MoneySuccessResponse(result)

		c.JSON(http.StatusOK, data)
	})

}
