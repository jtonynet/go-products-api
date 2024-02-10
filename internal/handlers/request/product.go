package request

import (
	"github.com/google/uuid"
)

type Product struct {
	UUID        uuid.UUID `json:"uuid" form:"uuid" binding:"required" example:"7829cc30-1d6e-4a5d-bcc1-ec65c8c338ab"`
	Name        string    `json:"name" form:"name" binding:"required" example:"Exis Boxis Series G"`
	Description string    `json:"description" form:"description" binding:"required" example:"Next Gen powerful videogame"`
	Price       *int64    `json:"price" form:"price" binding:"required" example:"5000"`
}

type UpdateProduct struct {
	Name        string `json:"name" binding:"required" form:"name" example:"Exis Boxis Series S"`
	Description string `json:"description"  binding:"required" form:"description" example:"The best experience of Powerfull"`
	Price       *int64 `json:"price" binding:"required" form:"price" example:"6500"`
}
