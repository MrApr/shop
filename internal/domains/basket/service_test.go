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

// TestBasketService_DisableUserActiveBasket functionality
func TestBasketService_DisableUserActiveBasket(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up temporary database connection failed")

	sv := createBasketService(conn)
	randUserId := rand.Int()

	mockedBasket := mockAndInsertBasket(conn, 1, randUserId, true)
	defer destructBasket(conn, mockedBasket)

	err = sv.DisableUserActiveBasket(randUserId)
	assert.NoError(t, err, "Disabling basket failed")

	disabledBasket := new(Basket)
	result := conn.Where("id = ?", mockedBasket[0].Id).First(disabledBasket)
	assert.NoError(t, result.Error, "Fetching basket in basket service failed")
	assert.False(t, disabledBasket.Status, "Previous basket is not disabled in basket service  new basket creation")

}

// TestBasketService_AddProductsToBasket functionality
func TestBasketService_AddProductsToBasket(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up temporary database connection failed")

	service := createBasketService(conn)
	randUserId := rand.Int()

	insufficientProduct := mockAndInsertProducts(conn, 0)
	defer destructCreatedProduct(conn, insufficientProduct)
	_, err = service.AddProductsToBasket(randUserId, insufficientProduct.Id, 1)
	assert.Error(t, err, "Product inventory is insufficient, expected error")
	assert.ErrorIs(t, err, ProductIsFinished, "product inventory is empty, expected ProductIsFinished error")

	okProduct := mockAndInsertProducts(conn, 2)
	defer destructCreatedProduct(conn, okProduct)

	_, err = service.AddProductsToBasket(randUserId, okProduct.Id, 5)
	assert.Error(t, err, "Product inventory is lower than expected, expected error")
	assert.ErrorIs(t, err, InsufficientProductAmount, "product amount is insufficient, expected insufficient error")

	tmpBasket, err := service.AddProductsToBasket(randUserId, okProduct.Id, 2)
	assert.NoError(t, err, "Cannot add product to basket")
	assert.NotNil(t, tmpBasket, "Expected basket, but got nil")
	assert.Equal(t, tmpBasket.UserId, randUserId, "Basket is not correct for current user")
	assert.True(t, tmpBasket.Status, "Product is added to a basket with false status")
	assert.NotNil(t, tmpBasket.Products, "product is not added to basket")
	assert.Equal(t, tmpBasket.Products[0].Id, okProduct.Id, "Added product and expected product are not equal")
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
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up temporary database connection failed")

	service := createBasketService(conn)
	randUserId := rand.Int()

	okProduct := mockAndInsertProducts(conn, 2)
	defer destructCreatedProduct(conn, okProduct)

	_, err = service.UpdateBasketProductsAmount(randUserId, okProduct.Id, 5)
	assert.Error(t, err, "Expected error, when no active basket for user exists")
	assert.ErrorIs(t, err, NoActiveBasket, "Expected error, when no active basket for user exists")

	mockedBasket := mockAndInsertBasket(conn, 1, randUserId, true)
	defer destructBasket(conn, mockedBasket)

	productAmountInBasket := okProduct.Amount - 1
	mockedBasketProducts := mockAndInsertBasketProduct(conn, 1, mockedBasket[0].Id, okProduct.Id, productAmountInBasket)
	defer destructBasketProduct(conn, mockedBasketProducts)
	assert.Equal(t, len(mockedBasketProducts), 1, "Mocking basket products failed")

	_, err = service.UpdateBasketProductsAmount(randUserId, okProduct.Id, 5)
	assert.Error(t, err, "Adding much more than products amount should cause and throw error")
	assert.ErrorIs(t, err, InsufficientProductAmount, "Adding much more than products amount should cause and throw error")

	updatedBasketProduct, err := service.UpdateBasketProductsAmount(randUserId, okProduct.Id, 2)
	assert.NoError(t, err, "Updating basket products failed")
	assert.Equal(t, updatedBasketProduct.Products[0].Id, okProduct.Id)

	basketProduct := new(BasketProduct)
	fetchResult := conn.Where("product_id = ?", mockedBasketProducts[0].ProductId).Where("basket_id = ?", mockedBasketProducts[0].BasketId).First(basketProduct)
	assert.NoError(t, fetchResult.Error, "fetching basket product failed")
	assert.Equal(t, basketProduct.Amount, 2, "Updating basket products failed")
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

// mockAndInsertProducts into database
func mockAndInsertProducts(db *gorm.DB, amount int) *products.Product {
	randCode := rand.Int()
	randPrice := rand.Float64()

	p := &products.Product{
		Title:  "Test product",
		Code:   randCode,
		Amount: amount,
		Price:  randPrice,
	}

	db.Create(p)

	return p
}

// destructCreatedProduct for testing purpose
func destructCreatedProduct(db *gorm.DB, createdProduct *products.Product) {
	db.Unscoped().Delete(createdProduct)
}
