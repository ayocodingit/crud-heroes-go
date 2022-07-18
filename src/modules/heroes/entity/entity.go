package entity

type Attribute struct {
	Attack *int `json:"attack"`
	Defend *int `json:"defend"`
	Armor  *int `json:"armor"`
}

type Hero struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Ability   string    `json:"ability"`
	Role      string    `json:"role"`
	Attribute Attribute `json:"attribute"`
}

type ResponseFindAll struct {
	Message string `json:"message"`
	Data    []Hero `json:"data"`
}

type ResponseFindById struct {
	Message string `json:"message"`
	Data    Hero   `json:"data"`
}
