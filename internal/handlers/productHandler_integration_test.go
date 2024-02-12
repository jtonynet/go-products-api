package handlers_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
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

func (suite *ValidationSuite) TestCreateSameProductTwice() {
	// Create Product
	respProductCreate, err := suite.createProduct()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusCreated, respProductCreate.Code)

	respBody := respProductCreate.Body.String()
	assert.Equal(suite.T(), gjson.Get(respBody, "msg").String(), "resource created successfully")

	// Create Same Product Again
	respProductCreate, err = suite.createProduct()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusConflict, respProductCreate.Code)

	respBody = respProductCreate.Body.String()
	assert.Equal(suite.T(), gjson.Get(respBody, "msg").String(), "resource conflict")
}

func (suite *ValidationSuite) TestCreateSameProductAfterSoftDeletion() {
	// Create Product
	respProductCreate, err := suite.createProduct()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusCreated, respProductCreate.Code)

	respBody := respProductCreate.Body.String()
	assert.Equal(suite.T(), gjson.Get(respBody, "msg").String(), "resource created successfully")

	// Delete Product
	suite.echoRouter = helpers.SetupEchoRouter()
	suite.echoRouter.DELETE("/products/:product_id", suite.productHandler.DeleteProductById)

	productUUIDRoute := fmt.Sprintf("/products/%s", suite.productUUID)
	reqProductDelete, err := http.NewRequest("DELETE", productUUIDRoute, nil)
	assert.NoError(suite.T(), err)
	respProductDelete := httptest.NewRecorder()

	suite.echoRouter.ServeHTTP(respProductDelete, reqProductDelete)
	assert.Equal(suite.T(), http.StatusNoContent, respProductDelete.Code)

	// Create Same Product After Soft Deleted
	respProductCreate, err = suite.createProduct()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusConflict, respProductCreate.Code)

	respBody = respProductCreate.Body.String()
	assert.Equal(suite.T(), gjson.Get(respBody, "msg").String(), "resource conflict")
}

func (suite *ValidationSuite) TestCreateProductWithIncorrectUUID() {
	// Create Product
	requestProductJSON := fmt.Sprintf(
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

	respProductCreate, err := suite.createProductWithJSONParam(requestProductJSON)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusBadRequest, respProductCreate.Code)

	respBody := respProductCreate.Body.String()
	assert.Equal(suite.T(), gjson.Get(respBody, "msg").String(), "field UUID is invalid")
}

func (suite *ValidationSuite) TestCreateProductWithNameLengthLessThanThreeChars() {
	// Create Product
	requestProductJSON := fmt.Sprintf(
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

	respProductCreate, err := suite.createProductWithJSONParam(requestProductJSON)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusBadRequest, respProductCreate.Code)

	respBody := respProductCreate.Body.String()
	assert.Equal(suite.T(), gjson.Get(respBody, "msg").String(), "field Name is invalid")
}

func (suite *ValidationSuite) TestCreateProductWithDescriptionLengthLessThanThreeChars() {
	// Create Product
	requestProductJSON := fmt.Sprintf(
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

	respProductCreate, err := suite.createProductWithJSONParam(requestProductJSON)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusBadRequest, respProductCreate.Code)

	respBody := respProductCreate.Body.String()
	assert.Equal(suite.T(), gjson.Get(respBody, "msg").String(), "field Description is invalid")
}

func (suite *ValidationSuite) TestCreateProductWithPriceIsZero() {
	// Create Product
	requestProductJSON := fmt.Sprintf(
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

	respProductCreate, err := suite.createProductWithJSONParam(requestProductJSON)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusBadRequest, respProductCreate.Code)

	respBody := respProductCreate.Body.String()
	assert.Equal(suite.T(), gjson.Get(respBody, "msg").String(), "field Price is invalid")
}

func (suite *ValidationSuite) TestUpdateProductWithNameLengthLessThanThreeChars() {
	// Create Product
	respProductCreate, err := suite.createProduct()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusCreated, respProductCreate.Code)

	// Update Product
	requestProductJSON := `{"name":"ab"}`
	respProductUpdate, err := suite.updateProductWithParams(suite.productUUID, requestProductJSON)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusBadRequest, respProductUpdate.Code)

	respBody := respProductUpdate.Body.String()
	assert.Equal(suite.T(), gjson.Get(respBody, "msg").String(), "field Name is invalid")
}

func (suite *ValidationSuite) TestUpdateProductWithNameLengthGreaterThanTwoHundredAndFiftyFiveChars() {
	// Create Product
	respProductCreate, err := suite.createProduct()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusCreated, respProductCreate.Code)

	// Update Product
	requestProductJSON := `{"name":"` + strings.Repeat("x", 300) + `"}`
	respProductUpdate, err := suite.updateProductWithParams(suite.productUUID, requestProductJSON)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusBadRequest, respProductUpdate.Code)

	respBody := respProductUpdate.Body.String()
	assert.Equal(suite.T(), gjson.Get(respBody, "msg").String(), "field Name is invalid")
}

func (suite *ValidationSuite) TestUpdateProductWithDescriptionLengthLessThanThreeChars() {
	// Create Product
	respProductCreate, err := suite.createProduct()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusCreated, respProductCreate.Code)

	// Update Product
	requestProductJSON := `{"description":"ab"}`
	respProductUpdate, err := suite.updateProductWithParams(suite.productUUID, requestProductJSON)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusBadRequest, respProductUpdate.Code)

	respBody := respProductUpdate.Body.String()
	assert.Equal(suite.T(), gjson.Get(respBody, "msg").String(), "field Description is invalid")
}

func (suite *ValidationSuite) TestUpdateProductWithPriceIsZero() {
	// Create Product
	respProductCreate, err := suite.createProduct()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusCreated, respProductCreate.Code)

	// Update Product
	requestProductJSON := `{"price":0}`
	respProductUpdate, err := suite.updateProductWithParams(suite.productUUID, requestProductJSON)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusBadRequest, respProductUpdate.Code)

	respBody := respProductUpdate.Body.String()
	assert.Equal(suite.T(), gjson.Get(respBody, "msg").String(), "field Price is invalid")
}

func (suite *ValidationSuite) TestUpdateProductWithPriceIsNegative() {
	// Create Product
	respProductCreate, err := suite.createProduct()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusCreated, respProductCreate.Code)

	// Update Product
	requestProductJSON := `{"price":-1}`
	respProductUpdate, err := suite.updateProductWithParams(suite.productUUID, requestProductJSON)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusBadRequest, respProductUpdate.Code)

	respBody := respProductUpdate.Body.String()
	assert.Equal(suite.T(), gjson.Get(respBody, "msg").String(), "field Price is invalid")
}

func (suite *ValidationSuite) createProduct() (*httptest.ResponseRecorder, error) {
	requestProductJSON := fmt.Sprintf(
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

	respProductCreate, err := suite.createProductWithJSONParam(requestProductJSON)

	return respProductCreate, err
}

func (suite *ValidationSuite) createProductWithJSONParam(requestProductJSON string) (*httptest.ResponseRecorder, error) {
	suite.echoRouter = helpers.SetupEchoRouter()
	suite.echoRouter.POST("/products", suite.productHandler.CreateProduct)

	reqProductCreate, err := http.NewRequest("POST", "/products", bytes.NewBuffer([]byte(requestProductJSON)))

	if err != nil {
		return nil, err
	}

	reqProductCreate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	respProductCreate := httptest.NewRecorder()
	suite.echoRouter.ServeHTTP(respProductCreate, reqProductCreate)

	return respProductCreate, nil
}

func (suite *ValidationSuite) updateProductWithParams(productUUID uuid.UUID, requestProductJSON string) (*httptest.ResponseRecorder, error) {
	suite.echoRouter = helpers.SetupEchoRouter()
	suite.echoRouter.PATCH("/products/:product_id", suite.productHandler.UpdateProductById)

	productUUIDRoute := fmt.Sprintf("/products/%s", productUUID)
	reqProductUpdate, err := http.NewRequest("PATCH", productUUIDRoute, bytes.NewBuffer([]byte(requestProductJSON)))

	if err != nil {
		return nil, err
	}

	reqProductUpdate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	respProductUpdate := httptest.NewRecorder()
	suite.echoRouter.ServeHTTP(respProductUpdate, reqProductUpdate)

	return respProductUpdate, nil
}

func TestValidadionSuite(t *testing.T) {
	suite.Run(t, new(ValidationSuite))
}
