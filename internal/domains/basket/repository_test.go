package basket

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"shop/internal/domains/products"
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
	defer destructBasket(conn, mockedBasked)
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
	defer destructBasket(conn, mockedBasked)
	assert.Equal(t, len(mockedBasked), testingCount, "Created basket and required basket are not equal in Testing basket repository")

	_, err = repo.GetUserActiveBasket(randUserId)
	assert.Error(t, err, "Fetching data from basket repository failed, no user active basket should exists")

	mockedBasked = mockAndInsertBasket(conn, testingCount, randUserId, true)
	assert.Equal(t, len(mockedBasked), testingCount, "Created basket and required basket are not equal in Testing basket repository")

	result, err := repo.GetUserActiveBasket(randUserId)
	assert.NoError(t, err, "Fetching data from basket repository failed, User should have active basket")
	assertBasketsEqual(t, mockedBasked, []Basket{*result})
}

// TestSqlBasket_GetUserBaskets functionality
func TestSqlBasket_GetUserBaskets(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up temporary database connection failed")

	repo := createRepo(conn)
	randUserId := rand.Int()
	testingCount := 5

	baskets, err := repo.GetUserBaskets(randUserId)
	assert.Equal(t, len(baskets), 0, "Fetching data from basket repository failed, User doesnt have any basket")

	mockedBasked := mockAndInsertBasket(conn, testingCount, randUserId, true)
	defer destructBasket(conn, mockedBasked)
	assert.Equal(t, len(mockedBasked), testingCount, "Created basket and required basket are not equal in Testing basket repository")

	result, err := repo.GetUserBaskets(randUserId)
	assert.NoError(t, err, "Fetching data from basket repository failed, User should have baskets")
	assertBasketsEqual(t, mockedBasked, result)
}

// TestSqlBasket_BasketExists functionality
func TestSqlBasket_BasketExists(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up temporary database connection failed")

	repo := createRepo(conn)
	randUserId := rand.Int()
	testingCount := 1

	mockedBasked := mockAndInsertBasket(conn, testingCount, randUserId, true)
	defer destructBasket(conn, mockedBasked)
	assert.Equal(t, len(mockedBasked), testingCount, "Created basket and required basket are not equal in Testing basket repository")

	exists := repo.BasketExists(mockedBasked[0].Id)
	assert.True(t, exists, "Fetching data from basket repository failed, basket exists")

	randWrongId := rand.Int()
	exists = repo.BasketExists(randWrongId)
	assert.False(t, exists, "Fetching data from basket repository failed, Basket does not exists")
}

// TestSqlBasket_GetBasketProduct functionality
func TestSqlBasket_GetBasketProduct(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up temporary database connection failed")

	repo := createRepo(conn)
	randUserId := rand.Int()
	testingCount := 1

	mockedBasket := mockAndInsertBasket(conn, 1, randUserId, true)
	defer destructBasket(conn, mockedBasket)

	mockedBasketProducts := mockAndInsertBasketProduct(conn, testingCount, mockedBasket[0].Id, 0, 0)
	defer destructBasketProduct(conn, mockedBasketProducts)
	assert.Equal(t, len(mockedBasketProducts), testingCount, "Mocking basket products failed")

	result, err := repo.GetBasketProduct(mockedBasket[0].Id, mockedBasketProducts[0].ProductId)
	assert.NoError(t, err, "fetching basket product failed")
	assertBasketProductsEqual(t, mockedBasketProducts, []BasketProduct{*result})

	randWrongBasketId := rand.Int()
	result, err = repo.GetBasketProduct(randWrongBasketId, mockedBasketProducts[0].ProductId)
	assert.Error(t, err, "Expected error on wrong basket Id")

	randWrongProductId := rand.Int()
	result, err = repo.GetBasketProduct(mockedBasket[0].Id, randWrongProductId)
	assert.Error(t, err, "Expected error on wrong basket Id")
}

// TestSqlBasket_CreateBasket functionality
func TestSqlBasket_CreateBasket(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up temporary database connection failed")

	repo := createRepo(conn)
	randUserId := rand.Int()

	mockedBasket := mockBasket(randUserId, true)

	err = repo.CreateBasket(mockedBasket)
	defer destructBasket(conn, []Basket{*mockedBasket})
	assert.NoError(t, err, "User basket creation failed")

	fetchedBasket := new(Basket)
	result := conn.Where("id = ?", mockedBasket.Id).First(fetchedBasket)
	assert.NoError(t, result.Error, "Basket did not created properly")
}

// TestSqlBasket_UpdateBasketProducts functionality
func TestSqlBasket_UpdateBasketProducts(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up temporary database connection failed")

	repo := createRepo(conn)
	randUserId := rand.Int()

	mockedBasket := mockAndInsertBasket(conn, 1, randUserId, true)
	defer destructBasket(conn, mockedBasket)

	mockedBasketProducts := mockAndInsertBasketProduct(conn, 1, mockedBasket[0].Id, 0, 0)
	defer destructBasketProduct(conn, mockedBasketProducts)
	assert.Equal(t, len(mockedBasketProducts), 1, "Mocking basket products failed")

	newProductId := rand.Int()
	newAmount := rand.Int()
	newPrice := rand.Float64()

	mockedBasketProducts[0].ProductId = newProductId
	mockedBasketProducts[0].Amount = newAmount
	mockedBasketProducts[0].UnitPrice = newPrice

	err = repo.UpdateBasketProducts(&mockedBasketProducts[0])
	assert.NoError(t, err, "Updating basket product failed, expected no error")

	tmpBasketProduct := new(BasketProduct)
	result := conn.Where("basket_id = ?", mockedBasketProducts[0].BasketId).Where("product_id = ?", mockedBasketProducts[0].ProductId).First(tmpBasketProduct)
	assert.NoError(t, result.Error, "Adding product to basket failed")
	assertBasketProductsEqual(t, []BasketProduct{mockedBasketProducts[0]}, []BasketProduct{*tmpBasketProduct})
}

// TestSqlBasket_AddProductToBasket functionality
func TestSqlBasket_AddProductToBasket(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up temporary database connection failed")

	repo := createRepo(conn)
	randUserId := rand.Int()

	mockedBasket := mockAndInsertBasket(conn, 1, randUserId, true)
	defer destructBasket(conn, mockedBasket)

	mockedBasketProduct := mockBasketProduct(mockedBasket[0].Id, 0, 0)

	err = repo.AddProductToBasket(&mockedBasket[0], mockedBasketProduct)
	defer destructBasketProduct(conn, []BasketProduct{*mockedBasketProduct})
	assert.NoError(t, err, "Adding product to basket failed")

	tmpBasketProduct := new(BasketProduct)
	result := conn.Where("basket_id = ?", mockedBasketProduct.BasketId).Where("product_id = ?", mockedBasketProduct.ProductId).First(tmpBasketProduct)
	assert.NoError(t, result.Error, "Adding product to basket failed")
	assertBasketProductsEqual(t, []BasketProduct{*mockedBasketProduct}, []BasketProduct{*tmpBasketProduct})
}

// TestSqlBasket_ClearBasketProducts functionality
func TestSqlBasket_ClearBasketProducts(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up temporary database connection failed")

	repo := createRepo(conn)
	randUserId := rand.Int()

	mockedBasket := mockAndInsertBasket(conn, 1, randUserId, true)
	defer destructBasket(conn, mockedBasket)

	mockedBasketProducts := mockAndInsertBasketProduct(conn, 1, mockedBasket[0].Id, 0, 0)
	defer destructBasketProduct(conn, mockedBasketProducts)
	assert.Equal(t, len(mockedBasketProducts), 1, "Mocking basket products failed")

	err = repo.ClearBasketProducts(&mockedBasket[0])
	assert.NoError(t, err, "Clearing basket products failed")

	var tmpBasketProducts []BasketProduct
	result := conn.Where("basket_id = ?", mockedBasket[0].Id).Find(&tmpBasketProducts)
	assert.NoError(t, result.Error, "Clearing products from basket failed")
	assert.Equal(t, len(tmpBasketProducts), 0, "Clearing products from basket failed")
}

// TestSqlBasket_DisableBasket functionality
func TestSqlBasket_DisableBasket(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up temporary database connection failed")

	repo := createRepo(conn)
	randUserId := rand.Int()

	mockedBasket := mockAndInsertBasket(conn, 1, randUserId, true)
	defer destructBasket(conn, mockedBasket)

	err = repo.DisableBasket(&mockedBasket[0])
	assert.NoError(t, err, "Disabling basket failed")
	assert.False(t, mockedBasket[0].Status, "Disabling basket failed")
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
	err = db.AutoMigrate(Basket{}, BasketProduct{}, products.Product{})
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

// mockAndInsertBasketProduct and return them
func mockAndInsertBasketProduct(db *gorm.DB, count, basketId, productId, amount int) []BasketProduct {
	basketProducts := make([]BasketProduct, 0, count)

	for i := 0; i < count; i++ {
		mockedBasketProduct := mockBasketProduct(basketId, productId, amount)
		result := db.Create(mockedBasketProduct)
		if result.Error != nil {
			continue
		}
		basketProducts = append(basketProducts, *mockedBasketProduct)
	}

	return basketProducts
}

// mockBasketProduct and return it
func mockBasketProduct(basketId, productId, amount int) *BasketProduct {
	if productId == 0 {
		productId = rand.Int()
	}

	if amount == 0 {
		amount = rand.Int()
	}

	price := rand.Float64()
	return &BasketProduct{
		BasketId:  basketId,
		ProductId: productId,
		Amount:    amount,
		UnitPrice: price,
	}
}

// destructBasket which are created for test
func destructBasket(db *gorm.DB, baskets []Basket) {
	for _, basket := range baskets {
		db.Unscoped().Delete(basket)
	}
}

// destructBasketProduct and delete them
func destructBasketProduct(db *gorm.DB, basketProducts []BasketProduct) {
	for _, basketProduct := range basketProducts {
		db.Where("basket_id = ?", basketProduct.BasketId).Where("product_id = ?", basketProduct.ProductId).Delete(basketProduct)
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

// assertBasketProductsEqual or not
func assertBasketProductsEqual(t *testing.T, mockedBaskedProducts, fetchedBasketProducts []BasketProduct) {
	for index := range mockedBaskedProducts {
		assert.Equal(t, mockedBaskedProducts[index].BasketId, fetchedBasketProducts[index].BasketId)
		assert.Equal(t, mockedBaskedProducts[index].ProductId, fetchedBasketProducts[index].ProductId)
		assert.Equal(t, mockedBaskedProducts[index].Amount, fetchedBasketProducts[index].Amount)
		assert.Equal(t, mockedBaskedProducts[index].UnitPrice, fetchedBasketProducts[index].UnitPrice)
	}
}
