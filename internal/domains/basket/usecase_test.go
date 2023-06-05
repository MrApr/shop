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

}

// TestBasketUseCase_AddProductsToBasket functionality
func TestBasketUseCase_AddProductsToBasket(t *testing.T) {

}

// TestBasketUseCase_UpdateBasketProductsAmount functionality
func TestBasketUseCase_UpdateBasketProductsAmount(t *testing.T) {

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
