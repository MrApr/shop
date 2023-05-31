package basket

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

// TestSqlBasket_GetBasketById functionality
func TestSqlBasket_GetBasketById(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up temporary database connection failed")

	repo := createRepo(conn)
	randUserId := rand.Int()
	testingCount := 1

	mockedBasked := mockAndInsertBasket(conn, testingCount, randUserId, true)
	assert.Equal(t, len(mockedBasked), testingCount, "Created basket and required basket are not equal in Testing basket repository")

	result, err := repo.GetBasketById(mockedBasked[0].Id)
	assert.NoError(t, err, "Fetching data from basket repository failed")
	assertBasketsEqual(t, mockedBasked, []Basket{*result})

	randWrongId := rand.Int()
	_, err = repo.GetBasketById(randWrongId)
	assert.Error(t, err, "Fetching data from basket repository failed, expected error but got none")
}

// TestSqlBasket_GetUserActiveBasket functionality
func TestSqlBasket_GetUserActiveBasket(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up temporary database connection failed")

	repo := createRepo(conn)
	randUserId := rand.Int()
	testingCount := 1

	mockedBasked := mockAndInsertBasket(conn, testingCount, randUserId, false)
	assert.Equal(t, len(mockedBasked), testingCount, "Created basket and required basket are not equal in Testing basket repository")

	_, err = repo.GetUserActiveBasket(mockedBasked[0].Id)
	assert.Error(t, err, "Fetching data from basket repository failed, no user active basket should exists")

	mockedBasked = mockAndInsertBasket(conn, testingCount, randUserId, true)
	assert.Equal(t, len(mockedBasked), testingCount, "Created basket and required basket are not equal in Testing basket repository")

	result, err := repo.GetUserActiveBasket(randUserId)
	assert.NoError(t, err, "Fetching data from basket repository failed, User should have active basket")
	assertBasketsEqual(t, mockedBasked, []Basket{*result})
}

// TestSqlBasket_GetUserBaskets functionality
func TestSqlBasket_GetUserBaskets(t *testing.T) {

}

// TestSqlBasket_BasketExists functionality
func TestSqlBasket_BasketExists(t *testing.T) {

}

// TestSqlBasket_GetBasketProduct functionality
func TestSqlBasket_GetBasketProduct(t *testing.T) {

}

// TestSqlBasket_CreateBasket functionality
func TestSqlBasket_CreateBasket(t *testing.T) {

}

// TestSqlBasket_UpdateBasketProducts functionality
func TestSqlBasket_UpdateBasketProducts(t *testing.T) {

}

// TestSqlBasket_AddProductToBasket functionality
func TestSqlBasket_AddProductToBasket(t *testing.T) {

}

// TestSqlBasket_ClearBasketProducts functionality
func TestSqlBasket_ClearBasketProducts(t *testing.T) {

}

// TestSqlBasket_DisableBasket functionality
func TestSqlBasket_DisableBasket(t *testing.T) {

}

// createRepo and return it
func createRepo(db *gorm.DB) BasketRepositoryInterface {
	return NewBasketRepository(db)
}

// setupDbConnection and run migration
func setupDbConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(Basket{}, BasketProduct{})
	return db, err
}

// mockAndInsertBasket into temporary database
func mockAndInsertBasket(db *gorm.DB, count, userId int, status bool) []Basket {
	baskets := make([]Basket, 0, count)

	for i := 0; i < count; i++ {
		mockedBasket := mockBasket(userId, status)
		result := db.Create(mockedBasket)
		if result.Error != nil {
			continue
		}
		baskets = append(baskets, *mockedBasket)
	}

	return baskets
}

// mockBasket and return it for testing purpose
func mockBasket(userId int, status bool) *Basket {
	if userId == 0 {
		userId = rand.Int()
	}

	return &Basket{
		UserId: userId,
		Status: status,
	}
}

// destructBasket which are created for test
func destructBasket(db *gorm.DB, baskets []Basket) {
	for _, basket := range baskets {
		db.Unscoped().Delete(basket)
	}
}

// assertBasketsEqual or not
func assertBasketsEqual(t *testing.T, mockedBaskets, fetchedBaskets []Basket) {
	for index := range mockedBaskets {
		assert.Equal(t, mockedBaskets[index].Id, fetchedBaskets[index].Id, "Mocked and fetched baskets are not equal")
		assert.Equal(t, mockedBaskets[index].UserId, fetchedBaskets[index].UserId, "Mocked and fetched baskets are not equal")
		assert.Equal(t, mockedBaskets[index].Status, fetchedBaskets[index].Status, "Mocked and fetched baskets are not equal")
	}
}
