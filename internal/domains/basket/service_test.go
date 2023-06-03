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

}

// TestBasketService_CreateBasket functionality
func TestBasketService_CreateBasket(t *testing.T) {

}

// TestBasketService_AddProductsToBasket functionality
func TestBasketService_AddProductsToBasket(t *testing.T) {

}

// TestBasketService_DisableActiveBasket functionality
func TestBasketService_DisableActiveBasket(t *testing.T) {

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
