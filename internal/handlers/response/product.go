package response

import (
	"time"

	"github.com/google/uuid"
	"github.com/jtonynet/go-products-api/internal/entity"
)

type Result struct {
	Msg string `json:"msg"`
}

type ResultProductList struct {
	Msg   string    `json:"msg"`
	Items []Product `json:"items"`
}

type Product struct {
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int64     `json:"price"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewProduct(p entity.Product) *Product {
	return &Product{
		UUID:        p.UUID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}
