package interfaces

import "github.com/jtonynet/go-products-api/internal/entity"

type ProductDB interface {
	Create(*entity.Product)
}
