package router

import (
	"fmt"

	"github.com/jtonynet/go-products-api/config"
	"github.com/jtonynet/go-products-api/internal/handlers"
	"github.com/labstack/echo/v4"
)

func Init(cfg *config.API, productHandler *handlers.ProductHandler) {
	r := echo.New()

	initializeRoutes(productHandler, r)

	portStr := fmt.Sprintf(":%s", cfg.Port)
	r.Logger.Fatal(r.Start(portStr))
}

func initializeRoutes(productHandler *handlers.ProductHandler, r *echo.Echo) {
	r.POST("/products", productHandler.CreateProduct)
	r.GET("/products", productHandler.RetriveProductList)
	r.GET("/products/:product_id", productHandler.RetrieveProductById)
	r.PATCH("/products/:product_id", productHandler.UpdateProductById)
	r.DELETE("/products/:product_id", productHandler.DeleteProductById)
}
