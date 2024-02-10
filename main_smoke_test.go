package main_smoke_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/jtonynet/go-products-api/config"
	"github.com/jtonynet/go-products-api/internal/database"
	"github.com/jtonynet/go-products-api/internal/handlers"
)

type SmokeSuite struct {
	suite.Suite
	echoRouter *echo.Echo
	DB         *gorm.DB

	productHandler *handlers.ProductHandler

	productUUID        uuid.UUID
	productName        string
	productDescription string
	productPrice       int64

	productUpdateName        string
	productUpdateDescription string
	productUpdatePrice       int64
}

func (suite *SmokeSuite) SetupSuite() {
	cfg := setupConfig()
	suite.DB, _ = setupDB(&cfg.Database)

	productDB := database.NewProductsDB(suite.DB)
	suite.productHandler = handlers.NewProductHandler(productDB)

	suite.productUUID, _ = uuid.Parse("37ad0486-5ba5-47a5-b352-408b077e12c6")
	suite.productName = "Playstation 7"
	suite.productDescription = "Next generation video game"
	suite.productPrice = 4000

	suite.productUpdateName = "Playstation 8"
	suite.productUpdateDescription = "The best experience of next generation video game"
	suite.productUpdatePrice = 5500
}

func (suite *SmokeSuite) TearDownSuite() {
	deleteProduct := fmt.Sprintf(
		`delete from products where uuid = "%s";`,
		suite.productUUID.String(),
	)
	suite.DB.Exec(deleteProduct)
}

func setupConfig() *config.Config {
	cfg := config.Config{}

	cfg.Database.Host = "localhost"
	cfg.Database.Port = "3306"
	cfg.Database.User = "root"
	cfg.Database.Pass = "root"
	cfg.Database.DB = "go-products-api"
	cfg.Database.RetryMaxElapsedTimeInMs = 1000

	return &cfg
}

func setupDB(cfg *config.Database) (*gorm.DB, error) {
	db, err := database.NewDatabase(cfg)
	return db, err
}

func (suite *SmokeSuite) TestSmoke() {
	suite.createAndRetrieveProductSuccessful()
	suite.updateAndRetrieveProductListSuccessful()
	suite.deleteProductSuccessful()
}

func setupEchoRouter() *echo.Echo {
	echoRouter := echo.New()
	return echoRouter
}

func (suite *SmokeSuite) createAndRetrieveProductSuccessful() {
	// Create Product
	suite.echoRouter = setupEchoRouter()
	suite.echoRouter.POST("/products", suite.productHandler.CreateProduct)

	requestProduct := fmt.Sprintf(
		`{
			"uuid":"%s",
			"name":"%s",
			"description":"%s",
			"price":%v
		}`,
		suite.productUUID,
		suite.productName,
		suite.productDescription,
		suite.productPrice,
	)

	reqProductCreate, err := http.NewRequest("POST", "/products", bytes.NewBuffer([]byte(requestProduct)))
	reqProductCreate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	assert.NoError(suite.T(), err)
	respProductCreate := httptest.NewRecorder()

	suite.echoRouter.ServeHTTP(respProductCreate, reqProductCreate)
	assert.Equal(suite.T(), http.StatusCreated, respProductCreate.Code)

	// Retrieve Product By Id
	suite.echoRouter = setupEchoRouter()
	suite.echoRouter.GET("/products/:product_id", suite.productHandler.RetrieveProductById)

	productUUIDRoute := fmt.Sprintf("/products/%s", suite.productUUID)
	reqProductRetrieve, err := http.NewRequest("GET", productUUIDRoute, nil)
	assert.NoError(suite.T(), err)
	respProductRetrieve := httptest.NewRecorder()

	suite.echoRouter.ServeHTTP(respProductRetrieve, reqProductRetrieve)
	assert.Equal(suite.T(), http.StatusOK, respProductRetrieve.Code)

	bodyProduct := respProductRetrieve.Body.String()
	assert.Equal(suite.T(), gjson.Get(bodyProduct, "uuid").String(), suite.productUUID.String())
	assert.Equal(suite.T(), gjson.Get(bodyProduct, "name").String(), suite.productName)
	assert.Equal(suite.T(), gjson.Get(bodyProduct, "description").String(), suite.productDescription)
	assert.Equal(suite.T(), gjson.Get(bodyProduct, "price").Int(), suite.productPrice)
}

func (suite *SmokeSuite) updateAndRetrieveProductListSuccessful() {
	// Update Product
	suite.echoRouter = setupEchoRouter()
	suite.echoRouter.PATCH("/products/:product_id", suite.productHandler.UpdateProductById)

	requestProduct := fmt.Sprintf(
		`{
			"name":"%s",
			"description":"%s",
			"price":%v
		}`,
		suite.productUpdateName,
		suite.productUpdateDescription,
		suite.productUpdatePrice,
	)

	productUUIDRoute := fmt.Sprintf("/products/%s", suite.productUUID)
	reqProductUpdate, err := http.NewRequest("PATCH", productUUIDRoute, bytes.NewBuffer([]byte(requestProduct)))
	reqProductUpdate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	assert.NoError(suite.T(), err)
	respProductUpdate := httptest.NewRecorder()

	suite.echoRouter.ServeHTTP(respProductUpdate, reqProductUpdate)
	assert.Equal(suite.T(), http.StatusOK, respProductUpdate.Code)

	// Retrieve Products List
	suite.echoRouter = setupEchoRouter()
	suite.echoRouter.GET("/products", suite.productHandler.RetriveProductList)

	reqProductsRetrieve, err := http.NewRequest("GET", "/products", nil)
	assert.NoError(suite.T(), err)
	respProductsRetrieve := httptest.NewRecorder()

	suite.echoRouter.ServeHTTP(respProductsRetrieve, reqProductsRetrieve)
	assert.Equal(suite.T(), http.StatusOK, respProductsRetrieve.Code)

	bodyProducts := respProductsRetrieve.Body.String()
	productInList, err := getProductJSONByUUID(bodyProducts, suite.productUUID.String())
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), gjson.Get(productInList, "name").String(), suite.productUpdateName)
	assert.Equal(suite.T(), gjson.Get(productInList, "description").String(), suite.productUpdateDescription)
	assert.Equal(suite.T(), gjson.Get(productInList, "price").Int(), suite.productUpdatePrice)
}

func (suite *SmokeSuite) deleteProductSuccessful() {
	// Delete Product
	suite.echoRouter = setupEchoRouter()
	suite.echoRouter.DELETE("/products/:product_id", suite.productHandler.DeleteProductById)

	productUUIDRoute := fmt.Sprintf("/products/%s", suite.productUUID)
	reqProductDelete, err := http.NewRequest("DELETE", productUUIDRoute, nil)
	assert.NoError(suite.T(), err)
	respProductDelete := httptest.NewRecorder()
	suite.echoRouter.ServeHTTP(respProductDelete, reqProductDelete)
	assert.Equal(suite.T(), http.StatusNoContent, respProductDelete.Code)

	// Retrieve Product By Id
	suite.echoRouter = setupEchoRouter()
	suite.echoRouter.GET("/products/:product_id", suite.productHandler.RetrieveProductById)

	reqProductRetrieve, err := http.NewRequest("GET", productUUIDRoute, nil)
	assert.NoError(suite.T(), err)
	respProductRetrieve := httptest.NewRecorder()

	suite.echoRouter.ServeHTTP(respProductRetrieve, reqProductRetrieve)
	assert.Equal(suite.T(), http.StatusNotFound, respProductRetrieve.Code)
}

func getProductJSONByUUID(productsJSON string, uuidToFind string) (string, error) {
	var products []map[string]interface{}
	err := json.Unmarshal([]byte(productsJSON), &products)
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

func TestSmokeSuite(t *testing.T) {
	suite.Run(t, new(SmokeSuite))
}
