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
	"gorm.io/gorm"
)

type ValidationSuite struct {
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

func (suite *ValidationSuite) SetupSuite() {
	cfg := helpers.SetupConfig()
	suite.DB, _ = helpers.SetupDB(&cfg.Database)

	productDB := database.NewProductsDB(suite.DB)
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
	fmt.Println("ValidadionSuite Finished")

}

func (suite *ValidationSuite) TestCreateSameProductTwice() {
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

	respProductCreate = httptest.NewRecorder()
	suite.echoRouter.ServeHTTP(respProductCreate, reqProductCreate)
	suite.T().Log(respProductCreate.Code)
}

/*
func (suite *ValidationSuite) TestValidation() {
	suite.createSameProductTwice()
	assert.Equal(suite.T(), http.StatusCreated, 201)
}
*/

func TestValidadionSuite(t *testing.T) {
	suite.Run(t, new(ValidationSuite))
}
