package http

import (
	"github.com/ayocodingit/crud-heroes-go/src/modules/heroes/handler"
	"github.com/gin-gonic/gin"
)

func Handler(router *gin.Engine) {
	hero := router.Group("v1/heroes")

	hero.GET("/", handler.FindAll)
	hero.GET("/:id", handler.FindById)
	hero.POST("/", handler.Store)
	hero.PUT("/:id", handler.Update)
	hero.DELETE("/:id", handler.Delete)
}
