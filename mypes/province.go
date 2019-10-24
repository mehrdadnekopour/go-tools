package mypes

// Province ....
type Province struct {
	Model
	Name   string `json:"name"`
	Cities Cities `json:"cities"`
}

// Provinces ...
type Provinces []*Province
