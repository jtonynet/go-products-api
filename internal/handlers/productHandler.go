package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateProduct(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "Rota POST /products foi acessada mas nao implementada")
}

func RetrieveProductById(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "Rota GET /products/:id foi acessada mas nao implementada")
}
