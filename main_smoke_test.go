package main_smoke_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/jtonynet/go-products-api/config"
	"github.com/jtonynet/go-products-api/internal/database"
	"github.com/jtonynet/go-products-api/internal/entity"
	"github.com/jtonynet/go-products-api/internal/handlers"
)

type SmokeSuite struct {
	suite.Suite
	router *echo.Echo
	DB     *gorm.DB

	productUUID        uuid.UUID
	productName        string
	productDescription string
	productPrice       int64
}

func (suite *SmokeSuite) SetupSuite() {
	cfg := setupConfig()
	suite.DB, _ = setupDB(&cfg.Database)

	suite.productUUID, _ = uuid.Parse("37ad0486-5ba5-47a5-b352-408b077e12c6")
	suite.productName = "Playstation 7"
	suite.productDescription = "Next generation video game"
	suite.productPrice = 4000
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
	cfg.Database.Port = 3306
	cfg.Database.User = "root"
	cfg.Database.Pass = "root"
	cfg.Database.DB = "go-products-api"

	return &cfg
}

func setupDB(cfg *config.Database) (*gorm.DB, error) {
	strConn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?parseTime=true",
		cfg.User,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.DB,
	)

	db, err := gorm.Open(mysql.Open(strConn))
	if err != nil {
		fmt.Println("Unable to connect to Database")
		return nil, err
	}

	db.AutoMigrate(entity.Product{})
	return db, nil
}

func (suite *SmokeSuite) TestSmoke() {
	suite.createAndRetrieveProductSuccessful()
}

func setupRouter() *echo.Echo {
	router := echo.New()
	return router
}

func (suite *SmokeSuite) createAndRetrieveProductSuccessful() {
	// Create Product
	productDB := database.NewProductsDB(suite.DB)
	productHandler := handlers.NewProductHandler(productDB)

	suite.router = setupRouter()
	suite.router.POST("/products", productHandler.CreateProduct)

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

	suite.router.ServeHTTP(respProductCreate, reqProductCreate)
	assert.Equal(suite.T(), http.StatusCreated, respProductCreate.Code)

	// Retrieve Product By Id
	suite.router = setupRouter()
	suite.router.GET("/products/:product_id", productHandler.RetrieveProductById)

	productsUUIDRoute := fmt.Sprintf("/products/%s", suite.productUUID)
	reqProductRetrieve, err := http.NewRequest("GET", productsUUIDRoute, nil)
	assert.NoError(suite.T(), err)
	respProductRetrieve := httptest.NewRecorder()

	suite.router.ServeHTTP(respProductRetrieve, reqProductRetrieve)
	assert.Equal(suite.T(), http.StatusOK, respProductRetrieve.Code)

	bodyProduct := respProductRetrieve.Body.String()
	assert.Equal(suite.T(), gjson.Get(bodyProduct, "uuid").String(), suite.productUUID.String())
	assert.Equal(suite.T(), gjson.Get(bodyProduct, "name").String(), suite.productName)
	assert.Equal(suite.T(), gjson.Get(bodyProduct, "description").String(), suite.productDescription)
	assert.Equal(suite.T(), gjson.Get(bodyProduct, "price").Int(), suite.productPrice)
}

func TestSmokeSuite(t *testing.T) {
	suite.Run(t, new(SmokeSuite))
}
