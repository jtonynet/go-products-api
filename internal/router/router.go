package router

import (
	"fmt"

	"github.com/jtonynet/go-products-api/config"
	"github.com/jtonynet/go-products-api/internal/handlers"
	"github.com/jtonynet/go-products-api/internal/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "github.com/jtonynet/go-products-api/api"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Init(cfg *config.API, productHandler *handlers.ProductHandler) {
	e := echo.New()
	e.Use(echoMiddleware.CORS())
	// e.Use(echoMiddleware.Logger())

	initializeRoutes(cfg, productHandler, e)

	portStr := fmt.Sprintf(":%s", cfg.Port)
	e.Logger.Fatal(e.Start(portStr))
}

func initializeRoutes(
	cfg *config.API,
	productHandler *handlers.ProductHandler,
	e *echo.Echo) {

	if cfg.MetricsEnabled {
		initializeMetricsRoute(e, cfg)
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/products", productHandler.CreateProduct)
	e.GET("/products", productHandler.RetriveProductList)
	e.GET("/products/:product_id", productHandler.RetrieveProductById)
	e.PATCH("/products/:product_id", productHandler.UpdateProductById)
	e.DELETE("/products/:product_id", productHandler.DeleteProductById)
}

func initializeMetricsRoute(e *echo.Echo, cfg *config.API) {
	middleware.InitPrometheus(cfg)

	e.Use(middleware.Prometheus())
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
}
