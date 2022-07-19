package http

import (
	"database/sql"

	"github.com/ayocodingit/crud-heroes-go/src/modules/heroes/handler"
	"github.com/ayocodingit/crud-heroes-go/src/modules/heroes/repository"
	"github.com/ayocodingit/crud-heroes-go/src/modules/heroes/service"
	"github.com/gin-gonic/gin"
)

func Handler(router *gin.Engine, db *sql.DB) {
	repository := repository.New(db)
	service := service.New(repository)
	handler := handler.New(service)

	r := router.Group("v1/heroes")

	r.GET("/", handler.FindAll)
	r.GET("/:id", handler.FindById)
	r.POST("/", handler.Store)
	r.PUT("/:id", handler.Update)
	r.DELETE("/:id", handler.Delete)
}
