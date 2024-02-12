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

// @Summary Create Product
// @Description Create One Product
// @Tags Products
// @Accept json
// @Produce json
// @Param request body request.Product true "Request body for create"
// @Success 201 {object} response.Result
// @Failure 400 {object} response.Result
// @Failure 500 {object} response.Result
// @Router /products [post]
func (ph *ProductHandler) CreateProduct(c echo.Context) error {
	reqP := new(request.Product)
	if err := c.Bind(reqP); err != nil {
		return c.JSON(http.StatusBadRequest, response.Result{Msg: "bad request"})
	}

	productEntity := entity.NewProduct(
		reqP.UUID,
		reqP.Name,
		reqP.Description,
		*reqP.Price,
	)

	if err := ph.productDB.DB.Create(&productEntity).Error; err != nil {
		if ph.productDB.ErrIsDuplicateKey(err) {
			return c.JSON(http.StatusConflict, response.Result{Msg: "resource conflict"})
		}

		return c.JSON(http.StatusInternalServerError, response.Result{Msg: "an error occurred, please try again later"})
	}

	return c.JSON(http.StatusCreated, response.Result{Msg: "resource created successfully"})
}

// @Summary Retrieve Product By UUID
// @Description Retrieve one Product
// @Tags Products
// @Produce json
// @Param product_id path string true "Product UUID"
// @Success 200 {object} response.Product
// @Failure 400 {object} response.Result
// @Failure 404 {object} response.Result
// @Router /products/{product_id} [get]
func (ph *ProductHandler) RetrieveProductById(c echo.Context) error {
	productId := c.Param("product_id")
	if !isValidUUID(productId) {
		return c.JSON(http.StatusBadRequest, response.Result{Msg: "bad request"})
	}
	productUUID := uuid.MustParse(productId)

	p := entity.Product{UUID: productUUID}
	if err := ph.productDB.DB.Where(&entity.Product{UUID: productUUID}).First(&p).Error; err != nil {
		return c.JSON(http.StatusNotFound, response.Result{Msg: "product not found"})
	}

	respP := response.NewProduct(p)
	return c.JSON(http.StatusOK, respP)
}

// @Summary Update Product By UUID
// @Description Update One Product
// @Tags Products
// @Accept json
// @Produce json
// @Param product_id path string true "Product UUID"
// @Param request body request.UpdateProduct true "Request body for update"
// @Success 200 {object} response.Result
// @Failure 400 {object} response.Result
// @Failure 404 {object} response.Result
// @Failure 500 {object} response.Result
// @Router /products/{product_id} [patch]
func (ph *ProductHandler) UpdateProductById(c echo.Context) error {
	productId := c.Param("product_id")
	if !isValidUUID(productId) {
		return c.JSON(http.StatusBadRequest, response.Result{Msg: "bad request"})
	}
	productUUID := uuid.MustParse(productId)

	p := entity.Product{UUID: productUUID}
	if err := ph.productDB.DB.Where(&entity.Product{UUID: productUUID}).First(&p).Error; err != nil {
		return c.JSON(http.StatusNotFound, response.Result{Msg: "product not found"})
	}

	reqP := new(request.UpdateProduct)
	if err := c.Bind(reqP); err != nil {
		return c.JSON(http.StatusBadRequest, response.Result{Msg: "bad request"})
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
		return c.JSON(http.StatusInternalServerError, response.Result{Msg: "an error occurred, please try again later"})
	}

	return c.JSON(http.StatusOK, response.Result{Msg: "product updated successfully"})
}

// @Summary Retrieve Product List
// @Description Retrieve List of Products
// @Tags Products
// @Produce json
// @Success 200 {object} response.ResultProductList
// @Failure 500 {object} response.Result
// @Router /products [get]
func (ph *ProductHandler) RetriveProductList(c echo.Context) error {
	products := []entity.Product{}
	if err := ph.productDB.DB.Find(&products).Order("id desc").Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.Result{Msg: "an error occurred, please try again later"})
	}

	productList := []response.Product{}
	for _, p := range products {
		productList = append(
			productList,
			*response.NewProduct(p),
		)
	}

	return c.JSON(http.StatusOK, response.ResultProductList{
		Msg:   "all products list",
		Items: productList,
	})
}

// @Summary Delete Product By UUID
// @Description Delete one Product
// @Tags Products
// @Produce json
// @Param product_id path string true "Product UUID"
// @Success 204
// @Failure 400 {object} response.Result
// @Failure 404 {object} response.Result
// @Failure 500 {object} response.Result
// @Router /products/{product_id} [delete]
func (ph *ProductHandler) DeleteProductById(c echo.Context) error {
	productId := c.Param("product_id")
	if !isValidUUID(productId) {
		return c.JSON(http.StatusBadRequest, response.Result{Msg: "bad request"})
	}
	productUUID := uuid.MustParse(productId)

	p := entity.Product{UUID: productUUID}
	if err := ph.productDB.DB.Where(&entity.Product{UUID: productUUID}).First(&p).Error; err != nil {
		return c.JSON(http.StatusNotFound, response.Result{Msg: "resource not found"})
	}

	if delP := ph.productDB.DB.Delete(&p); delP.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.Result{Msg: "an error occurred, please try again later"})
	}

	return c.JSON(http.StatusNoContent, response.Result{Msg: "resource deleted successfully"})
}

func isValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
