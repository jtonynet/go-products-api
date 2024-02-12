package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	UUID        uuid.UUID `gorm:"unique;not null"`
	Name        string    `gorm:"type:varchar(255)"`
	Description string    `gorm:"type:longtext"`
	Price       int64     `gorm:"type:integer"`
}

func NewProduct(
	uniqueId uuid.UUID,
	name string,
	description string,
	price int64,
) Product {
	return Product{
		UUID:        uniqueId,
		Name:        name,
		Description: description,
		Price:       price,
	}
}
