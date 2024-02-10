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
	r := echo.New()
	r.Use(echoMiddleware.CORS())

	initializeRoutes(cfg, productHandler, r)

	portStr := fmt.Sprintf(":%s", cfg.Port)
	r.Logger.Fatal(r.Start(portStr))
}

func initializeRoutes(
	cfg *config.API,
	productHandler *handlers.ProductHandler,
	r *echo.Echo) {

	if cfg.MetricsEnabled {
		initializeMetricsRoute(r, cfg)
	}

	r.GET("/swagger/*", echoSwagger.WrapHandler)

	r.POST("/products", productHandler.CreateProduct)
	r.GET("/products", productHandler.RetriveProductList)
	r.GET("/products/:product_id", productHandler.RetrieveProductById)
	r.PATCH("/products/:product_id", productHandler.UpdateProductById)
	r.DELETE("/products/:product_id", productHandler.DeleteProductById)
}

func initializeMetricsRoute(r *echo.Echo, cfg *config.API) {
	middleware.InitPrometheus(cfg)

	r.Use(middleware.Prometheus())
	r.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
}
