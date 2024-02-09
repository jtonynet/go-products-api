package main_smoke_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/jtonynet/go-products-api/internal/handlers"
)

type SmokeSuite struct {
	suite.Suite
	router *echo.Echo

	productUUID uuid.UUID
}

func (suite *SmokeSuite) SetupSuite() {
	suite.productUUID, _ = uuid.Parse("37ad0486-5ba5-47a5-b352-408b077e12c6")
}

func (suite *SmokeSuite) TearDownSuite() {
	fmt.Println("Terminei de rodar os smoke tests")
}

func (suite *SmokeSuite) TestSmoke() {
	suite.createAndRetrieveProductSuccessful()
}

func setupRouter() *echo.Echo {
	router := echo.New()
	return router
}

func (suite *SmokeSuite) createAndRetrieveProductSuccessful() {
	suite.router = setupRouter()
	suite.router.POST("/products", handlers.CreateProduct)

	requestProduct := fmt.Sprintf(
		`{
			"id":%s, 
			"name":"Playstation 7", 
			"description": "Videogame de ultima geracao", 
			"price":4000,
		}`,
		suite.productUUID)

	productRequestJSON, err := json.Marshal(requestProduct)
	assert.NoError(suite.T(), err)

	reqProductCreate, err := http.NewRequest("POST", "/products", bytes.NewBuffer(productRequestJSON))
	assert.NoError(suite.T(), err)
	respProductCreate := httptest.NewRecorder()
	suite.router.ServeHTTP(respProductCreate, reqProductCreate)
	assert.Equal(suite.T(), http.StatusCreated, respProductCreate.Code)

	suite.router = setupRouter()
	suite.router.GET("/products/:id", handlers.RetrieveProductById)
	productsUUIDRoute := fmt.Sprintf("/products/%s", suite.productUUID)
	reqProductRetrieve, err := http.NewRequest("GET", productsUUIDRoute, nil)
	assert.NoError(suite.T(), err)
	respProductRetrieve := httptest.NewRecorder()
	suite.router.ServeHTTP(respProductRetrieve, reqProductRetrieve)
	assert.Equal(suite.T(), http.StatusOK, respProductRetrieve.Code)
}

func TestSmokeSuite(t *testing.T) {
	suite.Run(t, new(SmokeSuite))
}
