package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/jtonynet/go-products-api/internal/database"
	"github.com/jtonynet/go-products-api/internal/entity"
	"github.com/jtonynet/go-products-api/internal/handlers/request"
	"github.com/jtonynet/go-products-api/internal/handlers/response"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productDB *database.ProductDB
}

func NewProductHandler(productDB *database.ProductDB) *ProductHandler {
	return &ProductHandler{
		productDB: productDB,
	}
}

func (ph *ProductHandler) CreateProduct(c echo.Context) error {
	reqP := new(request.Product)
	if err := c.Bind(reqP); err != nil {
		return c.JSON(http.StatusBadRequest, `{"msg": "bad request"}`)
	}

	productEntity := entity.NewProduct(
		reqP.UUID,
		reqP.Name,
		reqP.Description,
		reqP.Price,
	)

	if err := ph.productDB.DB.Create(&productEntity).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, `{"msg": "An error occurred, please try again later"}`)
	}

	return c.JSON(http.StatusCreated, `{"msg": "Product created successfully"}`)
}

func (ph *ProductHandler) RetrieveProductById(c echo.Context) error {
	productId := c.Param("product_id")
	if !isValidUUID(productId) {
		return c.JSON(http.StatusBadRequest, `{"msg": "bad request"}`)
	}
	productUUID := uuid.MustParse(productId)

	p := entity.Product{UUID: productUUID}
	if err := ph.productDB.DB.Where(&entity.Product{UUID: productUUID}).First(&p).Error; err != nil {
		return c.JSON(http.StatusNotFound, `{"msg": "Product not found"}`)
	}

	respP := response.NewProduct(p)
	return c.JSON(http.StatusOK, respP)
}

func isValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
