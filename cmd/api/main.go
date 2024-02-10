package main

import (
	"github.com/jtonynet/go-products-api/config"
	"github.com/jtonynet/go-products-api/internal/database"
	"github.com/jtonynet/go-products-api/internal/handlers"
	"github.com/jtonynet/go-products-api/internal/router"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic("cannot load enviroment variables")
	}

	db, err := database.NewDatabase(&cfg.Database)
	if err != nil {
		panic("cannot connect to database")
	}

	productDB := database.NewProductsDB(db)
	productHandler := handlers.NewProductHandler(productDB)

	router.Init(&cfg.API, productHandler)
}
