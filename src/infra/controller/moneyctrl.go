package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grrlopes/go-moneyhoney/src/application/usecase/listall"
	"github.com/grrlopes/go-moneyhoney/src/application/usecase/listbyid"
	"github.com/grrlopes/go-moneyhoney/src/application/usecase/login"
	"github.com/grrlopes/go-moneyhoney/src/application/usecase/save"
	"github.com/grrlopes/go-moneyhoney/src/application/usecase/update"
	"github.com/grrlopes/go-moneyhoney/src/application/usecase/usersave"
	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
	"github.com/grrlopes/go-moneyhoney/src/domain/repository"
	_validate "github.com/grrlopes/go-moneyhoney/src/domain/validator"
	"github.com/grrlopes/go-moneyhoney/src/infra/presenters"
	"github.com/grrlopes/go-moneyhoney/src/infra/repositories/couchdb"
	"github.com/grrlopes/go-moneyhoney/src/infra/repositories/mongodb"
)

var (
	repositories       repository.IMoneyRepo     = couchdb.NewMoneyRepository()
	repositorymongo    repository.IMongoRepo     = mongodb.NewMoneyRepository()
	repositoryuser     repository.IMongoUserRepo = mongodb.NewUserRepository()
	usecase_listall    listall.InputBoundary     = listall.NewFindAll(repositorymongo)
	usecase_save       save.InputBoundary        = save.NewSave(repositorymongo)
	usecase_listbyid   listbyid.InputBoundary    = listbyid.NewFindById(repositories)
	usecase_update     update.InputBoundary      = update.NewUpdate(repositories)
	usecase_createUser usersave.InputBoundary    = usersave.NewUserSave(repositoryuser)
	usecase_login      login.InputBoundary       = login.NewLogin(repositoryuser)
)

func MoneyCtrl(app gin.IRouter) {
	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "No valid endpoint provided!",
		})
	})

	app.GET("/findall", func(c *gin.Context) {
		var payload entity.Pagination
		err := c.ShouldBindJSON(&payload)

		checked, validErr := _validate.Validate(&payload)
		if checked {
			fieldErr := presenters.MoneyValidField(validErr)
			c.JSON(http.StatusBadRequest, fieldErr)
			return
		}

		result, count, err := usecase_listall.Execute(payload)

		if err != nil {
			error := presenters.MoneyError(result)
			c.JSON(http.StatusInternalServerError, error)
			return
		}

		data := presenters.MoneySuccess(result, count)

		c.JSON(http.StatusOK, data)
	})

	app.GET("/findbyid", func(c *gin.Context) {
		var payload entity.ById
		err := c.ShouldBind(&payload)

		checked, validErr := _validate.Validate(&payload)
		if checked {
			fieldErr := presenters.MoneyValidFieldResponse(validErr)
			c.JSON(http.StatusBadRequest, fieldErr)
			return
		}

		result, err := usecase_listbyid.Execute(&payload)

		if err != nil {
			error := presenters.MoneyErrorResponse(result)
			c.JSON(http.StatusInternalServerError, error)
			return
		}

		data := presenters.MoneySuccessResponse(result)

		c.JSON(http.StatusOK, data)
	})

	app.POST("/save", func(c *gin.Context) {
		var payload entity.Activity
		err := c.ShouldBindJSON(&payload)

		checked, validErr := _validate.Validate(&payload)
		if checked {
			fieldErr := presenters.MoneyValidFieldResponse(validErr)
			c.JSON(http.StatusBadRequest, fieldErr)
			return
		}

		result, err := usecase_save.Execute(&payload)

		if err != nil {
			error := presenters.MoneyErrorResponse(result)
			c.JSON(http.StatusInternalServerError, error)
			return
		}

		data := presenters.MoneySuccessResponse(result)

		c.JSON(http.StatusOK, data)
	})

	app.PUT("/update", func(c *gin.Context) {
		var payload entity.Value
		err := c.ShouldBindJSON(&payload)

		checked, validErr := _validate.Validate(&payload)
		if checked {
			fieldErr := presenters.MoneyValidFieldResponse(validErr)
			c.JSON(http.StatusBadRequest, fieldErr)
			return
		}

		result, err := usecase_update.Execute(&payload)

		if err != nil {
			error := presenters.MoneyErrorResponse(result)
			c.JSON(http.StatusInternalServerError, error)
			return
		}

		data := presenters.MoneySuccessResponse(result)

		c.JSON(http.StatusOK, data)
	})

	app.POST("/createuser", func(c *gin.Context) {
		var payload entity.Users
		err := c.ShouldBindJSON(&payload)

		checked, validErr := _validate.Validate(&payload)
		if checked {
			fieldErr := presenters.MoneyValidFieldResponse(validErr)
			c.JSON(http.StatusBadRequest, fieldErr)
			return
		}

		result, err := usecase_createUser.Execute(&payload)

		if err != nil {
			error := presenters.MoneyErrorResponse(result)
			c.JSON(http.StatusInternalServerError, error)
			return
		}

		data := presenters.MoneySuccessResponse(result)

		c.JSON(http.StatusOK, data)
	})

	app.POST("/login", func(c *gin.Context) {
		var payload entity.Users
		err := c.ShouldBindJSON(&payload)

		checked, validErr := _validate.Validate(&payload)
		if checked {
			fieldErr := presenters.MoneyValidFieldResponse(validErr)
			c.JSON(http.StatusBadRequest, fieldErr)
			return
		}

		result, err := usecase_login.Execute(&payload)

		if err != nil {
			// error := presenters.MoneyErrorResponse(result)
			// c.JSON(http.StatusInternalServerError, error)
			return
		}

		// data := presenters.MoneySuccessResponse(result)

		c.JSON(http.StatusOK, result)
	})

}
