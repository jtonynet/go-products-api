package request

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator"
)

type Product struct {
	UUID        string `json:"uuid" validate:"required,uuid" binding:"required" example:"7829cc30-1d6e-4a5d-bcc1-ec65c8c338ab"`
	Name        string `json:"name" validate:"required,min=3,max=255" binding:"required" example:"XBox 720 Series G"`
	Description string `json:"description" validate:"required,min=3" binding:"required" example:"Most Powerful MicroSoft video game"`
	Price       *int64 `json:"price" validate:"gt=0" binding:"required" example:"5000"`
}

func (p *Product) IsValid() (string, bool) {
	return isValid(p)
}

type UpdateProduct struct {
	Name        string `json:"name" validate:"omitempty,min=3,max=255" example:"XBox 1080 Series Z"`
	Description string `json:"description" validate:"omitempty,min=3" example:"Even more powerful than its predecessor"`
	Price       *int64 `json:"price" validate:"omitempty,gt=0" example:"6500"`
}

func (up *UpdateProduct) IsValid() (string, bool) {
	return isValid(up)
}

func isValid(a any) (string, bool) {
	validate := validator.New()

	err := validate.Struct(a)
	if err != nil {
		var errMsgs []string
		for _, err := range err.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is invalid", err.Field()))
		}
		return strings.Join(errMsgs, ", "), false
	}

	return "", true
}
