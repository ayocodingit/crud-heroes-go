package service

import (
	"github.com/ayocodingit/crud-heroes-go/src/modules/heroes/entity"
)

type service struct {
	repository entity.Repository
}

func New(repository entity.Repository) *service {
	return &service{repository}
}

func (s *service) Store(hero *entity.Hero) (err error) {
	err = s.repository.Store(hero)
	if err != nil {
		return
	}

	return
}

func (s *service) Update(hero entity.Hero) (err error) {
	err = s.repository.Update(hero)
	if err != nil {
		return
	}

	return
}

func (s *service) FindAll(req entity.QueryFindAll) []entity.Hero {
	return s.repository.FindAll(req)
}

func (s *service) FindById(id int) (hero entity.Hero, err error) {
	hero, err = s.repository.FindById(id)

	return
}

func (s *service) Delete(id int) {
	s.repository.Delete(id)
}
