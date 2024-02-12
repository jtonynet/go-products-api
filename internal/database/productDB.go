package database

import (
	"errors"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type ProductDB struct {
	DB *gorm.DB
}

func NewProductsDB(db *gorm.DB) *ProductDB {
	return &ProductDB{
		DB: db,
	}
}

func (pDB ProductDB) ErrIsDuplicateKey(err error) bool {
	// I haven't found a better way than this:
	// https://github.com/go-gorm/gorm/issues/4037
	var mysqlErr *mysql.MySQLError
	return errors.As(err, &mysqlErr) && mysqlErr.Number == 1062
}
