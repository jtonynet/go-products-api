package helpers

import (
	"github.com/jtonynet/go-products-api/config"
	"github.com/jtonynet/go-products-api/internal/database"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupConfig() *config.Config {
	cfg := config.Config{}

	cfg.Database.Host = "localhost"
	cfg.Database.Port = "3306"
	cfg.Database.User = "root"
	cfg.Database.Pass = "root"
	cfg.Database.DB = "go-products-api"
	cfg.Database.RetryMaxElapsedTimeInMs = 1000

	return &cfg
}

func SetupDB(cfg *config.Database) (*gorm.DB, error) {
	db, err := database.NewDatabase(cfg)
	return db, err
}

func SetupEchoRouter() *echo.Echo {
	echoRouter := echo.New()
	return echoRouter
}
