package basket

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"math/rand"
	"shop/internal/domains/products"
	"testing"
)

// TestBasketService_GetUserActiveBasket functionality
func TestBasketService_GetUserActiveBasket(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up temporary database connection failed")

	sv := createBasketService(conn)
	randUserId := rand.Int()
	testingCount := 1

	mockedBasked := mockAndInsertBasket(conn, testingCount, randUserId, false)
	defer destructBasket(conn, mockedBasked)
	assert.Equal(t, len(mockedBasked), testingCount, "Created basket and required basket are not equal in Testing basket service")

	_, err = sv.GetUserActiveBasket(randUserId)
	assert.Error(t, err, "Fetching data from basket repository failed, no user active basket should exists")

	mockedBasked = mockAndInsertBasket(conn, testingCount, randUserId, true)
	assert.Equal(t, len(mockedBasked), testingCount, "Created basket and required basket are not equal in Testing basket service")

	result, err := sv.GetUserActiveBasket(randUserId)
	assert.NoError(t, err, "Fetching data from basket service failed, User should have active basket")
	assertBasketsEqual(t, mockedBasked, []Basket{*result})
}

// TestBasketService_GetUserBaskets functionality
func TestBasketService_GetUserBaskets(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up temporary database connection failed")

	service := createBasketService(conn)
	randUserId := rand.Int()
	testingCount := 5

	baskets, err := service.GetUserBaskets(randUserId)
	assert.Equal(t, len(baskets), 0, "Fetching data from basket service failed, User doesnt have any basket")

	mockedBasked := mockAndInsertBasket(conn, testingCount, randUserId, true)
	defer destructBasket(conn, mockedBasked)
	assert.Equal(t, len(mockedBasked), testingCount, "Created basket and required basket are not equal in Testing basket service")

	result, err := service.GetUserBaskets(randUserId)
	assert.NoError(t, err, "Fetching data from basket service failed, User should have baskets")
	assertBasketsEqual(t, mockedBasked, result)
}

// TestBasketService_CreateBasket functionality
func TestBasketService_CreateBasket(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up temporary database connection failed")

	service := createBasketService(conn)
	randUserId := rand.Int()
	testingCount := 1

	mockedBasked := mockAndInsertBasket(conn, testingCount, randUserId, true)
	defer destructBasket(conn, mockedBasked)
	assert.Equal(t, len(mockedBasked), testingCount, "Created basket and required basket are not equal in Testing basket service")

	createdBasket, err := service.CreateBasket(randUserId)
	defer destructBasket(conn, []Basket{*createdBasket})
	assert.NoError(t, err, "User basket Creation failed in user service")
	assert.NotEqual(t, createdBasket.Id, mockedBasked[0].Id, "Previous basket is not disabled in basket service basket creation")

	disabledBasket := new(Basket)
	result := conn.Where("id = ?", mockedBasked[0].Id).First(disabledBasket)
	assert.NoError(t, result.Error, "Fetching basket in basket service failed")

	assert.False(t, disabledBasket.Status, "Previous basket is not disabled in basket service  new basket creation")
}

// TestBasketService_AddProductsToBasket functionality
func TestBasketService_AddProductsToBasket(t *testing.T) {

}

// TestBasketService_DisableActiveBasket functionality
func TestBasketService_DisableActiveBasket(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up temporary database connection failed")

	service := createBasketService(conn)
	randUserId := rand.Int()

	mockedBasket := mockAndInsertBasket(conn, 1, randUserId, true)
	defer destructBasket(conn, mockedBasket)

	err = service.DisableUserActiveBasket(randUserId)
	assert.NoError(t, err, "Disabling basket failed")

	disabledBasket := new(Basket)
	result := conn.Where("id = ?", mockedBasket[0].Id).First(disabledBasket)
	assert.NoError(t, result.Error, "Fetching basket in basket service failed")
	assert.False(t, disabledBasket.Status, "Basket service disabling basket functionality failed")

}

// TestBasketService_UpdateBasketProductsAmount functionality
func TestBasketService_UpdateBasketProductsAmount(t *testing.T) {

}

// createBasketService and return it for testing purpose
func createBasketService(db *gorm.DB) BasketServiceInterface {
	pSv := createProductService(db)
	return NewBasketService(NewBasketRepository(db), pSv)
}

// createProductService and return it
// BasketService service is tightly coupled with ProductService
func createProductService(db *gorm.DB) products.ProductServiceInterface {
	return products.NewProductsService(products.NewProductRepository(db))
}
