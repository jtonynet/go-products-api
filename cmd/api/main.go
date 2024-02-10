package main

import (
	"log/slog"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/jtonynet/go-products-api/config"
	"github.com/jtonynet/go-products-api/internal/database"
	"github.com/jtonynet/go-products-api/internal/handlers"
	"github.com/jtonynet/go-products-api/internal/router"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic("cannot load enviroment variables")
	}

	// Retry connecting to the database for RetryMaxElapsedTimeInMs Milliseconds
	var dbErr error
	var db *gorm.DB

	retry := backoff.NewExponentialBackOff()
	retry.MaxElapsedTime = time.Duration(cfg.API.RetryMaxElapsedTimeInMs) * time.Millisecond
	backoff.RetryNotify(func() error {
		db, dbErr = database.NewDatabase(&cfg.Database)
		return dbErr
	}, retry, func(err error, t time.Duration) {
		slog.Info("Retrying connect to Database after error: %v", err)
	})

	productDB := database.NewProductsDB(db)
	productHandler := handlers.NewProductHandler(productDB)

	router.Init(&cfg.API, productHandler)
}
