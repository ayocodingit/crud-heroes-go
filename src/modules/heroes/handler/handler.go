package handler

import (
	"net/http"
	"strconv"

	"github.com/ayocodingit/crud-heroes-go/src/helpers"
	"github.com/ayocodingit/crud-heroes-go/src/modules/heroes/entity"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service entity.Service
}

func New(service entity.Service) *handler {
	return &handler{service}
}

func (h handler) FindAll(c *gin.Context) {
	var req entity.QueryFindAll

	result := h.service.FindAll(req)

	res := entity.ResponseFindAll{Message: "success", Data: result}

	c.JSON(http.StatusOK, res)
}

func (h handler) FindById(c *gin.Context) {
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	}

	data, err := h.service.FindById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "data not found",
		})
		return
	}

	result := entity.ResponseFindById{Message: "success", Data: data}

	c.JSON(http.StatusOK, result)
}

func (h handler) Store(c *gin.Context) {
	var hero entity.Hero
	err := c.ShouldBindJSON(&hero)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = helpers.Validate(hero)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.service.Store(&hero)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, hero)
}

func (h handler) Update(c *gin.Context) {
	var hero entity.Hero

	err := c.ShouldBindJSON(&hero)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = helpers.Validate(hero)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err = h.service.FindById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "data not found",
		})
		return
	}

	hero.Id = id

	err = h.service.Update(hero)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, hero)

}

func (h handler) Delete(c *gin.Context) {
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	_, err = h.service.FindById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "data not found",
		})
		return
	}

	h.service.Delete(id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted!",
	})
}
