package entity

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

type ResponseFindAll struct {
	Message string `json:"message"`
	Data    []Hero `json:"data"`
}

type ResponseFindById struct {
	Message string `json:"message"`
	Data    Hero   `json:"data"`
}
