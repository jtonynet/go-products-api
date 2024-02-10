package request

import (
	"github.com/google/uuid"
)

type Product struct {
	UUID        uuid.UUID `json:"uuid" form:"uuid" binding:"required" example:"2e61ddac-c3cc-46e9-ba88-0e86a790c924"`
	Name        string    `json:"name" form:"name" binding:"required" example:"Playstation 7"`
	Description string    `json:"description" form:"description" binding:"required" example:"Videogame de ultima geracao"`
	Price       *int64    `json:"price" form:"price" binding:"required" example:"4000"`
}

type UpdateProduct struct {
	Name        string `json:"name" binding:"required" form:"name" example:"Playstation 8"`
	Description string `json:"description"  binding:"required" form:"description" example:"The best experience of next generation video game"`
	Price       *int64 `json:"price" binding:"required" form:"price" example:"4500"`
}
