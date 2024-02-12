package handlers_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/jtonynet/go-products-api/internal/database"
	"github.com/jtonynet/go-products-api/internal/handlers"
	"github.com/jtonynet/go-products-api/test/helpers"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/tidwall/gjson"
	"gorm.io/gorm"
)

type ValidationSuite struct {
	suite.Suite
	echoRouter *echo.Echo
	db         *gorm.DB

	productHandler *handlers.ProductHandler

	productUUID        uuid.UUID
	productName        string
	productDescription string
	productPrice       int64

	productUpdateName        string
	productUpdateDescription string
	productUpdatePrice       int64
}

func (suite *ValidationSuite) SetupSuite() {
	cfg := helpers.SetupConfig()
	suite.db, _ = helpers.SetupDB(&cfg.Database)

	productDB := database.NewProductsDB(suite.db)
	suite.productHandler = handlers.NewProductHandler(productDB)

	suite.productUUID, _ = uuid.Parse("d16384f2-4640-4e53-88f6-d873dfc0b80c")
	suite.productName = "XBox 720 Series G"
	suite.productDescription = "Most Powerful MicroSoft video game"
	suite.productPrice = 4500

	suite.productUpdateName = "XBox 1080 Series Z"
	suite.productUpdateDescription = "Even more powerful than its predecessor"
	suite.productUpdatePrice = 5500
}

func (suite *ValidationSuite) TearDownSuite() {
	deleteProduct := fmt.Sprintf(
		`delete from products where uuid = "%s";`,
		suite.productUUID.String(),
	)
	suite.db.Exec(deleteProduct)
}

func (suite *ValidationSuite) TearDownTest() {
	deleteProduct := fmt.Sprintf(
		`delete from products where uuid = "%s";`,
		suite.productUUID.String(),
	)
	suite.db.Exec(deleteProduct)
}

func (suite *ValidationSuite) TestCreateSameProductTwiceAndAfterSoftDeletion() {
	// Create Product
	suite.createProductSuccesfull()

	// Create Same Product Again
	suite.echoRouter = helpers.SetupEchoRouter()
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
	assert.Equal(suite.T(), http.StatusConflict, respProductCreate.Code)

	respBody := respProductCreate.Body.String()
	assert.Equal(suite.T(), gjson.Get(respBody, "msg").String(), "resource conflict")

	// Delete Product
	suite.echoRouter = helpers.SetupEchoRouter()
	suite.echoRouter.DELETE("/products/:product_id", suite.productHandler.DeleteProductById)

	productUUIDRoute := fmt.Sprintf("/products/%s", suite.productUUID)
	reqProductDelete, err := http.NewRequest("DELETE", productUUIDRoute, nil)
	assert.NoError(suite.T(), err)
	respProductDelete := httptest.NewRecorder()

	suite.echoRouter.ServeHTTP(respProductDelete, reqProductDelete)
	assert.Equal(suite.T(), http.StatusNoContent, respProductDelete.Code)

	respBody = respProductDelete.Body.String()
	assert.Equal(suite.T(), gjson.Get(respBody, "msg").String(), "resource deleted successfully")

	// Create Same Product Deleted
	suite.echoRouter = helpers.SetupEchoRouter()
	suite.echoRouter.POST("/products", suite.productHandler.CreateProduct)

	reqProductCreate, err = http.NewRequest("POST", "/products", bytes.NewBuffer([]byte(requestProduct)))
	reqProductCreate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	assert.NoError(suite.T(), err)
	respProductCreate = httptest.NewRecorder()

	suite.echoRouter.ServeHTTP(respProductCreate, reqProductCreate)
	assert.Equal(suite.T(), http.StatusConflict, respProductCreate.Code)

	respBody = respProductCreate.Body.String()
	assert.Equal(suite.T(), gjson.Get(respBody, "msg").String(), "resource conflict")
}

func (suite *ValidationSuite) TestCreateProductWithIncorrectUUID() {
	// Create Product
	suite.echoRouter = helpers.SetupEchoRouter()
	suite.echoRouter.POST("/products", suite.productHandler.CreateProduct)

	suite.echoRouter = helpers.SetupEchoRouter()
	suite.echoRouter.POST("/products", suite.productHandler.CreateProduct)

	requestProduct := fmt.Sprintf(
		`{
			"uuid":"%s",
			"name":"%s",
			"description":"%s",
			"price":%v
		}`,
		"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		suite.productName,
		suite.productDescription,
		suite.productPrice,
	)

	reqProductCreate, err := http.NewRequest("POST", "/products", bytes.NewBuffer([]byte(requestProduct)))
	reqProductCreate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	assert.NoError(suite.T(), err)
	respProductCreate := httptest.NewRecorder()

	suite.echoRouter.ServeHTTP(respProductCreate, reqProductCreate)
	assert.Equal(suite.T(), http.StatusBadRequest, respProductCreate.Code)

	respBody := respProductCreate.Body.String()
	assert.Equal(suite.T(), gjson.Get(respBody, "msg").String(), "field UUID is invalid")
}

func (suite *ValidationSuite) TestCreateProductWithNameLengthLessThanThreeChars() {
	// Create Product
	suite.echoRouter = helpers.SetupEchoRouter()
	suite.echoRouter.POST("/products", suite.productHandler.CreateProduct)

	suite.echoRouter = helpers.SetupEchoRouter()
	suite.echoRouter.POST("/products", suite.productHandler.CreateProduct)

	requestProduct := fmt.Sprintf(
		`{
			"uuid":"%s",
			"name":"%s",
			"description":"%s",
			"price":%v
		}`,
		suite.productUUID,
		"ab",
		suite.productDescription,
		suite.productPrice,
	)

	reqProductCreate, err := http.NewRequest("POST", "/products", bytes.NewBuffer([]byte(requestProduct)))
	reqProductCreate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	assert.NoError(suite.T(), err)
	respProductCreate := httptest.NewRecorder()

	suite.echoRouter.ServeHTTP(respProductCreate, reqProductCreate)
	assert.Equal(suite.T(), http.StatusBadRequest, respProductCreate.Code)

	respBody := respProductCreate.Body.String()
	assert.Equal(suite.T(), gjson.Get(respBody, "msg").String(), "field Name is invalid")
}

func (suite *ValidationSuite) TestCreateProductWithDescriptionLengthLessThanThreeChars() {
	// Create Product
	suite.echoRouter = helpers.SetupEchoRouter()
	suite.echoRouter.POST("/products", suite.productHandler.CreateProduct)

	suite.echoRouter = helpers.SetupEchoRouter()
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
		"ab",
		suite.productPrice,
	)

	reqProductCreate, err := http.NewRequest("POST", "/products", bytes.NewBuffer([]byte(requestProduct)))
	reqProductCreate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	assert.NoError(suite.T(), err)
	respProductCreate := httptest.NewRecorder()

	suite.echoRouter.ServeHTTP(respProductCreate, reqProductCreate)
	assert.Equal(suite.T(), http.StatusBadRequest, respProductCreate.Code)

	respBody := respProductCreate.Body.String()
	assert.Equal(suite.T(), gjson.Get(respBody, "msg").String(), "field Description is invalid")
}

func (suite *ValidationSuite) TestCreateProductWithPriceIsZero() {
	// Create Product
	suite.echoRouter = helpers.SetupEchoRouter()
	suite.echoRouter.POST("/products", suite.productHandler.CreateProduct)

	suite.echoRouter = helpers.SetupEchoRouter()
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
		0,
	)

	reqProductCreate, err := http.NewRequest("POST", "/products", bytes.NewBuffer([]byte(requestProduct)))
	reqProductCreate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	assert.NoError(suite.T(), err)
	respProductCreate := httptest.NewRecorder()

	suite.echoRouter.ServeHTTP(respProductCreate, reqProductCreate)
	assert.Equal(suite.T(), http.StatusBadRequest, respProductCreate.Code)

	respBody := respProductCreate.Body.String()
	assert.Equal(suite.T(), gjson.Get(respBody, "msg").String(), "field Price is invalid")
}

func (suite *ValidationSuite) TestUpdateProductWithNameLengthLessThanThreeChars() {
	// Create Product
	suite.createProductSuccesfull()

	// Update Product
	suite.echoRouter = helpers.SetupEchoRouter()
	suite.echoRouter.PATCH("/products/:product_id", suite.productHandler.UpdateProductById)

	requestProduct := `{"name":"ab"}`

	productUUIDRoute := fmt.Sprintf("/products/%s", suite.productUUID)
	reqProductUpdate, err := http.NewRequest("PATCH", productUUIDRoute, bytes.NewBuffer([]byte(requestProduct)))
	reqProductUpdate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	assert.NoError(suite.T(), err)
	respProductUpdate := httptest.NewRecorder()

	suite.echoRouter.ServeHTTP(respProductUpdate, reqProductUpdate)
	assert.Equal(suite.T(), http.StatusBadRequest, respProductUpdate.Code)

	respBody := respProductUpdate.Body.String()
	assert.Equal(suite.T(), gjson.Get(respBody, "msg").String(), "field Name is invalid")
}

func (suite *ValidationSuite) TestUpdateProductWithDescriptionLengthLessThanThreeChars() {
	// Create Product
	suite.createProductSuccesfull()

	// Update Product
	suite.echoRouter = helpers.SetupEchoRouter()
	suite.echoRouter.PATCH("/products/:product_id", suite.productHandler.UpdateProductById)

	requestProduct := `{"description":"ab"}`

	productUUIDRoute := fmt.Sprintf("/products/%s", suite.productUUID)
	reqProductUpdate, err := http.NewRequest("PATCH", productUUIDRoute, bytes.NewBuffer([]byte(requestProduct)))
	reqProductUpdate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	assert.NoError(suite.T(), err)
	respProductUpdate := httptest.NewRecorder()

	suite.echoRouter.ServeHTTP(respProductUpdate, reqProductUpdate)
	assert.Equal(suite.T(), http.StatusBadRequest, respProductUpdate.Code)

	respBody := respProductUpdate.Body.String()
	assert.Equal(suite.T(), gjson.Get(respBody, "msg").String(), "field Description is invalid")
}

func (suite *ValidationSuite) TestUpdateProductWithPriceIsZero() {
	// Create Product
	suite.createProductSuccesfull()

	// Update Product
	suite.echoRouter = helpers.SetupEchoRouter()
	suite.echoRouter.PATCH("/products/:product_id", suite.productHandler.UpdateProductById)

	requestProduct := `{"price":0}`

	productUUIDRoute := fmt.Sprintf("/products/%s", suite.productUUID)
	reqProductUpdate, err := http.NewRequest("PATCH", productUUIDRoute, bytes.NewBuffer([]byte(requestProduct)))
	reqProductUpdate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	assert.NoError(suite.T(), err)
	respProductUpdate := httptest.NewRecorder()

	suite.echoRouter.ServeHTTP(respProductUpdate, reqProductUpdate)
	assert.Equal(suite.T(), http.StatusBadRequest, respProductUpdate.Code)

	respBody := respProductUpdate.Body.String()
	assert.Equal(suite.T(), gjson.Get(respBody, "msg").String(), "field Price is invalid")
}

func (suite *ValidationSuite) createProductSuccesfull() {
	// Create Product
	suite.echoRouter = helpers.SetupEchoRouter()
	suite.echoRouter.POST("/products", suite.productHandler.CreateProduct)

	suite.echoRouter = helpers.SetupEchoRouter()
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

	respBody := respProductCreate.Body.String()
	assert.Equal(suite.T(), gjson.Get(respBody, "msg").String(), "resource created successfully")
}

func TestValidadionSuite(t *testing.T) {
	suite.Run(t, new(ValidationSuite))
}
