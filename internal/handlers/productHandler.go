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
		*reqP.Price,
	)

	if err := ph.productDB.DB.Create(&productEntity).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, `{"msg": "an error occurred, please try again later"}`)
	}

	return c.JSON(http.StatusCreated, `{"msg": "product created successfully"}`)
}

func (ph *ProductHandler) RetrieveProductById(c echo.Context) error {
	productId := c.Param("product_id")
	if !isValidUUID(productId) {
		return c.JSON(http.StatusBadRequest, `{"msg": "bad request"}`)
	}
	productUUID := uuid.MustParse(productId)

	p := entity.Product{UUID: productUUID}
	if err := ph.productDB.DB.Where(&entity.Product{UUID: productUUID}).First(&p).Error; err != nil {
		return c.JSON(http.StatusNotFound, `{"msg": "product not found"}`)
	}

	respP := response.NewProduct(p)
	return c.JSON(http.StatusOK, respP)
}

func (ph *ProductHandler) UpdateProductById(c echo.Context) error {
	productId := c.Param("product_id")
	if !isValidUUID(productId) {
		return c.JSON(http.StatusBadRequest, `{"msg": "bad request"}`)
	}
	productUUID := uuid.MustParse(productId)

	p := entity.Product{UUID: productUUID}
	if err := ph.productDB.DB.Where(&entity.Product{UUID: productUUID}).First(&p).Error; err != nil {
		return c.JSON(http.StatusNotFound, `{"msg": "product not found"}`)
	}

	reqP := new(request.UpdateProduct)
	if err := c.Bind(reqP); err != nil {
		return c.JSON(http.StatusBadRequest, `{"msg": "bad request"}`)
	}

	if reqP.Name != "" {
		p.Name = reqP.Name
	}

	if reqP.Description != "" {
		p.Description = reqP.Description
	}

	if reqP.Price != nil {
		p.Price = *reqP.Price
	}

	if err := ph.productDB.DB.Save(&p).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, `{"msg": "an error occurred, please try again later"}`)
	}

	return c.JSON(http.StatusOK, `{"msg": "product updated successfully"}`)
}

func (ph *ProductHandler) RetriveProductList(c echo.Context) error {
	products := []entity.Product{}
	if err := ph.productDB.DB.Find(&products).Order("id desc").Error; err != nil {
		return c.JSON(http.StatusInternalServerError, `{"msg": "an error occurred, please try again later"}`)
	}

	resp := []response.Product{}
	for _, p := range products {
		resp = append(
			resp,
			*response.NewProduct(p),
		)
	}

	return c.JSON(http.StatusOK, resp)
}

func (ph *ProductHandler) DeleteProductById(c echo.Context) error {
	productId := c.Param("product_id")
	if !isValidUUID(productId) {
		return c.JSON(http.StatusBadRequest, `{"msg": "bad request"}`)
	}
	productUUID := uuid.MustParse(productId)

	p := entity.Product{UUID: productUUID}
	if err := ph.productDB.DB.Where(&entity.Product{UUID: productUUID}).First(&p).Error; err != nil {
		return c.JSON(http.StatusNotFound, `{"msg": "product not found"}`)
	}

	if result := ph.productDB.DB.Delete(&p); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, `{"msg": "an error occurred, please try again later"}`)
	}

	return c.JSON(http.StatusNoContent, `{"msg": "product removed successfully"}`)
}

func isValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
