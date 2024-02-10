package database

import (
	"fmt"

	"github.com/jtonynet/go-products-api/config"
	"github.com/jtonynet/go-products-api/internal/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(cfg *config.Database) (*gorm.DB, error) {
	strConn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?parseTime=true",
		cfg.User,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.DB,
	)

	db, err := gorm.Open(mysql.Open(strConn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(entity.Product{})

	return db, nil
}
