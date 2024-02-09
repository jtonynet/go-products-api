package services_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jtonynet/go-products-api/internal/entity"
	"github.com/jtonynet/go-products-api/internal/services"
	"github.com/stretchr/testify/assert"
)

type FakeProductDB struct {
	Rows map[uuid.UUID]*entity.Product
}

func (fpDB *FakeProductDB) Create(product *entity.Product) {
	fpDB.Rows[product.ID] = product
}

func NewFakeProductDB() *FakeProductDB {
	return &FakeProductDB{
		Rows: make(map[uuid.UUID]*entity.Product),
	}
}

func TestCreateProduct(t *testing.T) {
	// Arrange
	fakeProductDB := NewFakeProductDB()
	createProduct := services.NewCreateProduct(fakeProductDB)
	productUUID, err := uuid.Parse("37ad0486-5ba5-47a5-b352-408b077e12c6")
	assert.NoError(t, err)
	product := entity.NewProduct(productUUID)

	// Act
	createProduct.With(product)

	// Assert
	assert.Equal(t, fakeProductDB.Rows[product.ID], product)
}
