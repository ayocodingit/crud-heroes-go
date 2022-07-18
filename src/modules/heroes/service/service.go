package service

import (
	"strconv"

	"github.com/ayocodingit/crud-heroes-go/src/modules/heroes/entity"
	"github.com/ayocodingit/crud-heroes-go/src/modules/heroes/repository"
	"github.com/gin-gonic/gin"
)

func Store(c *gin.Context) (hero entity.Hero, err error) {
	err = c.ShouldBindJSON(&hero)
	if err != nil {
		return
	}

	err = repository.Store(&hero)
	if err != nil {
		return
	}

	return
}

func Update(c *gin.Context) (hero entity.Hero, err error) {
	id := c.Param("id")
	err = c.ShouldBindJSON(&hero)

	if err != nil {
		return
	}

	heroId, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	hero.Id = heroId

	err = repository.Update(&hero)
	if err != nil {
		return
	}

	return
}

func FindAll() []entity.Hero {
	return repository.FindAll()
}

func FindById(id int) (data entity.Hero, err error) {
	data, err = repository.FindById(id)

	return
}

func Delete(id int) {
	repository.Delete(id)
}
