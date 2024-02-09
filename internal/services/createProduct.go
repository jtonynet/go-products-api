package services

import (
	"github.com/jtonynet/go-products-api/internal/entity"
	"github.com/jtonynet/go-products-api/internal/interfaces"
)

type CreateProduct struct {
	ProductDB interfaces.ProductDB
}

func NewCreateProduct(productDB interfaces.ProductDB) *CreateProduct {
	return &CreateProduct{
		ProductDB: productDB,
	}
}

func (cp *CreateProduct) With(product entity.Product) {
	cp.ProductDB.Create(&product)
}
