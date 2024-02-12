package helpers

import (
	"encoding/json"
	"errors"

	"github.com/jtonynet/go-products-api/config"
	"github.com/jtonynet/go-products-api/internal/database"
	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
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
	cfg.Database.MetricsEnabled = false

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

func GetProductFromItemsListByUUID(itemsList string, uuidToFind string) (string, error) {
	productItems := gjson.Get(itemsList, "items").String()
	var products []map[string]interface{}
	err := json.Unmarshal([]byte(productItems), &products)
	if err != nil {
		return "", err
	}

	for _, p := range products {
		if p["uuid"].(string) == uuidToFind {
			productJSON, err := json.Marshal(p)
			if err != nil {
				return "", err
			}
			return string(productJSON), nil
		}
	}
	return "", errors.New("json document not found")
}
