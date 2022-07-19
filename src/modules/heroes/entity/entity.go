package entity

import "github.com/gin-gonic/gin"

type Attribute struct {
	Attack *int `json:"attack" validate:"required,numeric,min=0,max=100"`
	Defend *int `json:"defend" validate:"required,numeric,min=0,max=100"`
	Armor  *int `json:"armor" validate:"required,numeric,min=0,max=100"`
}

type Hero struct {
	Id        int       `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Ability   string    `json:"ability" validate:"required"`
	Role      string    `json:"role" validate:"required"`
	Attribute Attribute `json:"attribute" validate:"required"`
}

type QueryFindAll struct {
	Page    string `json:"page"`
	PerPage string `json:"per_page"`
}

type ResponseFindAll struct {
	Message string `json:"message"`
	Data    []Hero `json:"data"`
}

type ResponseFindById struct {
	Message string `json:"message"`
	Data    Hero   `json:"data"`
}

type Repository interface {
	FindAll(req QueryFindAll) []Hero
	FindById(id int) (Hero, error)
	Update(hero Hero) error
	Delete(id int)
	Store(hero *Hero) error
}

type Service interface {
	FindAll(req QueryFindAll) []Hero
	FindById(id int) (Hero, error)
	Update(hero Hero) error
	Delete(id int)
	Store(hero *Hero) error
}

type Handler interface {
	FindAll(c *gin.Context)
	FindById(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Store(c *gin.Context)
}
