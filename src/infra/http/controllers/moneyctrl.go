package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grrlopes/go-moneyhoney/src/application/usecase/listall"
	"github.com/grrlopes/go-moneyhoney/src/application/usecase/listbyid"
	"github.com/grrlopes/go-moneyhoney/src/application/usecase/save"
	"github.com/grrlopes/go-moneyhoney/src/application/usecase/update"
	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
	"github.com/grrlopes/go-moneyhoney/src/domain/repository"
	_validate "github.com/grrlopes/go-moneyhoney/src/domain/validator"
	"github.com/grrlopes/go-moneyhoney/src/infra/presenters"
	"github.com/grrlopes/go-moneyhoney/src/infra/repositories/couchdb"
	"github.com/grrlopes/go-moneyhoney/src/infra/repositories/mongodb"
)

var (
	repositories    repository.IMoneyRepo  = couchdb.NewMoneyRepository()
	repositorymongo repository.IMongoRepo  = mongodb.NewMoneyRepository()
	usecaseListall  listall.InputBoundary  = listall.NewFindAll(repositorymongo)
	usecaseSave     save.InputBoundary     = save.NewSave(repositorymongo)
	usecaseListbyid listbyid.InputBoundary = listbyid.NewFindById(repositories)
	usecaseUpdate   update.InputBoundary   = update.NewUpdate(repositories)
)

func FindAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload entity.Pagination
		err := c.ShouldBindJSON(&payload)

		checked, validErr := _validate.Validate(&payload)
		if checked {
			fieldErr := presenters.MoneyValidField(validErr)
			c.JSON(http.StatusBadRequest, fieldErr)
			return
		}

		result, count, err := usecaseListall.Execute(payload)

		if err != nil {
			error := presenters.MoneyError(result)
			c.JSON(http.StatusInternalServerError, error)
			return
		}

		data := presenters.MoneySuccess(result, count)

		c.JSON(http.StatusOK, data)
	}
}

func FindById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload entity.ById
		err := c.ShouldBind(&payload)

		checked, validErr := _validate.Validate(&payload)
		if checked {
			fieldErr := presenters.MoneyValidFieldResponse(validErr)
			c.JSON(http.StatusBadRequest, fieldErr)
			return
		}

		result, err := usecaseListbyid.Execute(&payload)

		if err != nil {
			error := presenters.MoneyErrorResponse(result)
			c.JSON(http.StatusInternalServerError, error)
			return
		}

		data := presenters.MoneySuccessResponse(result)

		c.JSON(http.StatusOK, data)
	}
}

func Save() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload entity.Activity
		err := c.ShouldBindJSON(&payload)

		checked, validErr := _validate.Validate(&payload)
		if checked {
			fieldErr := presenters.MoneyValidFieldResponse(validErr)
			c.JSON(http.StatusBadRequest, fieldErr)
			return
		}

		result, err := usecaseSave.Execute(&payload)

		if err != nil {
			error := presenters.MoneyErrorResponse(result)
			c.JSON(http.StatusInternalServerError, error)
			return
		}

		data := presenters.MoneySuccessResponse(result)

		c.JSON(http.StatusOK, data)
	}
}

func Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload entity.Value
		err := c.ShouldBindJSON(&payload)

		checked, validErr := _validate.Validate(&payload)
		if checked {
			fieldErr := presenters.MoneyValidFieldResponse(validErr)
			c.JSON(http.StatusBadRequest, fieldErr)
			return
		}

		result, err := usecaseUpdate.Execute(&payload)

		if err != nil {
			error := presenters.MoneyErrorResponse(result)
			c.JSON(http.StatusInternalServerError, error)
			return
		}

		data := presenters.MoneySuccessResponse(result)

		c.JSON(http.StatusOK, data)
	}
}
