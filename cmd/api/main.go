package main

import (
	"fmt"

	"github.com/jtonynet/go-products-api/config"
	"github.com/jtonynet/go-products-api/internal/database"
	"github.com/jtonynet/go-products-api/internal/entity"
	"github.com/jtonynet/go-products-api/internal/handlers"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	cfg := config.Config{}

	cfg.API.Port = 8080

	cfg.Database.Host = "localhost"
	cfg.Database.Port = 3306
	cfg.Database.User = "root"
	cfg.Database.Pass = "root"
	cfg.Database.DB = "go-products-api"

	strConn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?parseTime=true",
		cfg.Database.User,
		cfg.Database.Pass,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DB,
	)

	db, err := gorm.Open(mysql.Open(strConn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("cannot connect to database")
	}
	db.AutoMigrate(entity.Product{})

	productDB := database.NewProductsDB(db)
	productHandler := handlers.NewProductHandler(productDB)

	router := echo.New()

	router.POST("/products", productHandler.CreateProduct)
	router.GET("/products", productHandler.RetriveProductList)
	router.GET("/products/:product_id", productHandler.RetrieveProductById)
	router.PATCH("/products/:product_id", productHandler.UpdateProductById)
	router.DELETE("/products/:product_id", productHandler.DeleteProductById)

	portStr := fmt.Sprintf(":%v", cfg.API.Port)
	router.Logger.Fatal(router.Start(portStr))

}
