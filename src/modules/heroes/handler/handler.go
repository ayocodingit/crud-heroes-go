package handler

import (
	"net/http"
	"strconv"

	"github.com/ayocodingit/crud-heroes-go/src/modules/heroes/entity"
	"github.com/ayocodingit/crud-heroes-go/src/modules/heroes/service"
	"github.com/gin-gonic/gin"
)

func FindAll(c *gin.Context) {
	result := entity.ResponseFindAll{Message: "success", Data: service.FindAll()}

	c.JSON(http.StatusOK, result)
}

func FindById(c *gin.Context) {
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	}

	data, err := service.FindById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "data not found",
		})
		return
	}

	result := entity.ResponseFindById{Message: "success", Data: data}

	c.JSON(http.StatusOK, result)
}

func Store(c *gin.Context) {
	hero, err := service.Store(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, hero)
}

func Update(c *gin.Context) {
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err = service.FindById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "data not found",
		})
		return
	}

	hero, err := service.Update(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, hero)

}

func Delete(c *gin.Context) {
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	}

	_, err = service.FindById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "data not found",
		})
		return
	}

	service.Delete(id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted!",
	})
}
