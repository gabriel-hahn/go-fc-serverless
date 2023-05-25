package data

import (
	"github.com/google/uuid"
)

type Product struct {
	ID 		 string  `json:"id"`
	Name 	 string  `json:"name"`
	Price  int     `json:"price"`
}

func (p *Product) GenerateID() {
	p.ID = uuid.New().String()
}
