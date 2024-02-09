package database

import (
	"gorm.io/gorm"
)

type ProductDB struct {
	DB *gorm.DB
}

func NewProductsDB(db *gorm.DB) *ProductDB {
	return &ProductDB{
		DB: db,
	}
}
