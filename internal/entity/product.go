package entity

import "github.com/google/uuid"

type Product struct {
	ID uuid.UUID
}

func NewProduct(id uuid.UUID) *Product {
	return &Product{
		ID: id,
	}
}
