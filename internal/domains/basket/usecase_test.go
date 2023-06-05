package basket

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

// TestBasketUseCase_GetUserActiveBasket functionality
func TestBasketUseCase_GetUserActiveBasket(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up temporary database connection failed")

	ctx := context.Background()
	uc := createBasketUseCase(conn)
	randUserId := 1
	testingCount := 1

	mockedBasked := mockAndInsertBasket(conn, testingCount, randUserId, true)
	assert.Equal(t, len(mockedBasked), testingCount, "Created basket and required basket are not equal in Testing basket service")

	result, err := uc.GetUserActiveBasket(ctx, "")
	assert.NoError(t, err, "Fetching data from basket service failed, User should have active basket")
	assertBasketsEqual(t, mockedBasked, []Basket{*result})
}

// TestBasketUseCase_GetUserBaskets functionality
func TestBasketUseCase_GetUserBaskets(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up temporary database connection failed")

	uc := createBasketUseCase(conn)
	ctx := context.Background()
	randUserId := 1
	testingCount := 5

	mockedBasked := mockAndInsertBasket(conn, testingCount, randUserId, true)
	defer destructBasket(conn, mockedBasked)
	assert.Equal(t, len(mockedBasked), testingCount, "Created basket and required basket are not equal in Testing basket service")

	result, err := uc.GetUserBaskets(ctx, "")
	assert.NoError(t, err, "Fetching data from basket service failed, User should have baskets")
	assertBasketsEqual(t, mockedBasked, result)
}

// TestBasketUseCase_CreateBasket functionality
func TestBasketUseCase_CreateBasket(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up temporary database connection failed")

	uc := createBasketUseCase(conn)
	ctx := context.Background()
	randUserId := 1
	testingCount := 1

	mockedBasked := mockAndInsertBasket(conn, testingCount, randUserId, true)
	defer destructBasket(conn, mockedBasked)
	assert.Equal(t, len(mockedBasked), testingCount, "Created basket and required basket are not equal in Testing basket service")

	createdBasket, err := uc.CreateBasket(ctx, "")
	defer destructBasket(conn, []Basket{*createdBasket})
	assert.NoError(t, err, "User basket Creation failed in user service")
	assert.NotEqual(t, createdBasket.Id, mockedBasked[0].Id, "Previous basket is not disabled in basket service basket creation")

	disabledBasket := new(Basket)
	result := conn.Where("id = ?", mockedBasked[0].Id).First(disabledBasket)
	assert.NoError(t, result.Error, "Fetching basket in basket service failed")

	assert.False(t, disabledBasket.Status, "Previous basket is not disabled in basket service  new basket creation")
}

// TestBasketUseCase_AddProductsToBasket functionality
func TestBasketUseCase_AddProductsToBasket(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up temporary database connection failed")

	ctx := context.Background()
	useCase := createBasketUseCase(conn)
	randUserId := 1

	okProduct := mockAndInsertProducts(conn, 2)
	defer destructCreatedProduct(conn, okProduct)

	mockedAddRequest := mockAddProductToBasketRequest(okProduct.Id, 2)

	tmpBasket, err := useCase.AddProductsToBasket(ctx, "", mockedAddRequest)
	assert.NoError(t, err, "Cannot add product to basket")
	assert.NotNil(t, tmpBasket, "Expected basket, but got nil")
	assert.Equal(t, tmpBasket.UserId, randUserId, "Basket is not correct for current user")
	assert.True(t, tmpBasket.Status, "Product is added to a basket with false status")
	assert.NotNil(t, tmpBasket.Products, "product is not added to basket")
	assert.Equal(t, tmpBasket.Products[0].Id, okProduct.Id, "Added product and expected product are not equal")
}

// TestBasketUseCase_UpdateBasketProductsAmount functionality
func TestBasketUseCase_UpdateBasketProductsAmount(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up temporary database connection failed")

	useCase := createBasketUseCase(conn)
	ctx := context.Background()
	randUserId := 1

	okProduct := mockAndInsertProducts(conn, 2)
	defer destructCreatedProduct(conn, okProduct)

	mockedBasket := mockAndInsertBasket(conn, 1, randUserId, true)
	defer destructBasket(conn, mockedBasket)

	productAmountInBasket := okProduct.Amount - 1
	mockedBasketProducts := mockAndInsertBasketProduct(conn, 1, mockedBasket[0].Id, okProduct.Id, productAmountInBasket)
	defer destructBasketProduct(conn, mockedBasketProducts)
	assert.Equal(t, len(mockedBasketProducts), 1, "Mocking basket products failed")

	mockedEditAmountRequest := mockEditProductInBasketRequest(okProduct.Id, 2)
	updatedBasketProduct, err := useCase.UpdateBasketProductsAmount(ctx, "", mockedEditAmountRequest)
	assert.NoError(t, err, "Updating basket products failed")
	assert.Equal(t, updatedBasketProduct.Products[0].Id, okProduct.Id)

	basketProduct := new(BasketProduct)
	fetchResult := conn.Where("product_id = ?", mockedBasketProducts[0].ProductId).Where("basket_id = ?", mockedBasketProducts[0].BasketId).First(basketProduct)
	assert.NoError(t, fetchResult.Error, "fetching basket product failed")
	assert.Equal(t, basketProduct.Amount, 2, "Updating basket products failed")
}

// TestBasketUseCase_DisableActiveBasket functionality
func TestBasketUseCase_DisableActiveBasket(t *testing.T) {

}

// createBasketUseCase and return it
func createBasketUseCase(db *gorm.DB) BasketUseCaseInterface {
	sv := createBasketService(db)
	return NewUseCase(sv, func(ctx context.Context, token string) (int, error) {
		return 1, nil
	})
}

// mockAddProductToBasketRequest and return it
func mockAddProductToBasketRequest(productId, amount int) *AddProductsToBasketRequest {
	return &AddProductsToBasketRequest{
		ProductId: productId,
		Amount:    amount,
	}
}

// mockEditProductInBasketRequest and return it
func mockEditProductInBasketRequest(productId, amount int) *EditProductsToBasketRequest {
	return &EditProductsToBasketRequest{
		ProductId: productId,
		Amount:    amount,
	}
}
