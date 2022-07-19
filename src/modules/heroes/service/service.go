package service

import (
	"github.com/ayocodingit/crud-heroes-go/src/modules/heroes/entity"
	"github.com/ayocodingit/crud-heroes-go/src/modules/heroes/repository"
)

func Store(hero *entity.Hero) (err error) {
	err = repository.Store(hero)
	if err != nil {
		return
	}

	return
}

func Update(hero *entity.Hero) (err error) {
	err = repository.Update(hero)
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
