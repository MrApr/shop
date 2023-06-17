package products

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"testing"
)

// TestCategoryRepository_GetAllCategories functionality
func TestCategoryRepository_GetAllCategories(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Connecting to database failed")

	repo := createCategoryRepository(db)
	createdCats := mockAndInsertCategories(db, 5)
	defer destructCreatedType(db, createdCats[0].TypeId)
	defer destructCreatedCategories(db, createdCats)

	fetchedCats := repo.GetAllCategories(nil, nil, nil, nil, 10)
	assert.Equal(t, len(fetchedCats), 0, "Zero categories must be fetched")

	falseParentCatId := rand.Int()
	fetchedCats = repo.GetAllCategories(nil, &falseParentCatId, nil, nil, 0)
	assert.Equal(t, len(fetchedCats), 0, "Zero categories must be fetched")

	falseTypeId := rand.Int()
	fetchedCats = repo.GetAllCategories(nil, nil, &falseTypeId, nil, 0)
	assert.Equal(t, len(fetchedCats), 0, "Zero categories must be fetched")

	limit := 1
	fetchedCats = repo.GetAllCategories(nil, nil, nil, &limit, 0)
	assert.Equal(t, len(fetchedCats), limit, "one Category must be fetched")

	falseTitle := "Test irrelevant category title which not exists"
	fetchedCats = repo.GetAllCategories(&falseTitle, nil, nil, nil, 0)
	assert.Equal(t, len(fetchedCats), 0, "zero Category must be fetched")

	fetchedCats = repo.GetAllCategories(nil, nil, nil, nil, 0)
	assert.NotZero(t, len(fetchedCats), "Zero categories fetched")
	assert.Equal(t, len(fetchedCats), 5, "Fetched categories are not equal")
	assertCategories(t, createdCats, fetchedCats)

	fetchedCats = repo.GetAllCategories(nil, nil, &createdCats[0].TypeId, nil, 0)
	assert.NotZero(t, len(fetchedCats), "Zero categories fetched")
	assertCategories(t, createdCats, fetchedCats)

	fetchedCats = repo.GetAllCategories(&createdCats[0].Title, nil, &createdCats[0].TypeId, nil, 0)
	assert.NotZero(t, len(fetchedCats), "Zero categories fetched")
	assertCategories(t, createdCats, fetchedCats)
}

// TestTypeRepository_GetAllTypes functionality
func TestTypeRepository_GetAllTypes(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up database connection failed")

	testingObjectCount := 5

	repo := createTypeRepo(conn)

	mockedTypes := mockAndInsertType(conn, testingObjectCount)
	defer destructAllTypes(conn, mockedTypes)
	assert.Equal(t, len(mockedTypes), testingObjectCount, "Creating Mock objects failed")

	fetchedTypes := repo.GetAllTypes(nil, testingObjectCount, 0)
	assert.Equal(t, len(fetchedTypes), testingObjectCount, "Fetching Types from repo failed")
	assertTypes(t, fetchedTypes, mockedTypes)

	fetchedTypes = repo.GetAllTypes(nil, testingObjectCount, 10)
	assert.Equal(t, len(fetchedTypes), 0, "Fetching Types from repo failed, on offset 10")

	var wrongName string = "TestIrreleventname"
	fetchedTypes = repo.GetAllTypes(&wrongName, testingObjectCount, 0)
	assert.Equal(t, len(fetchedTypes), 0, "Fetching Types from repo failed, on offset 10")
}

// TestProductRepository_GetAllProducts functionality
func TestProductRepository_GetAllProducts(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up database connection failed")

	testingObjectCount := 5

	repo := createProductRepository(conn)

	mockedProducts := mockAndInsertProducts(conn, testingObjectCount)
	assert.Equal(t, len(mockedProducts), testingObjectCount, "Mocking products failed")
	defer destructCreatedType(conn, mockedProducts[0].Categories[0].TypeId)
	defer destructCreatedCategories(conn, mockedProducts[0].Categories)
	defer destructCreatedProducts(conn, mockedProducts)

	fetchedProducts := repo.GetAllProducts([]int{rand.Int()}, nil, nil, nil, nil, nil, nil, nil, 0)
	assert.Equal(t, len(fetchedProducts), 0, "Fetching products with wrong category id failed")

	wrongText := "wrong given title"
	fetchedProducts = repo.GetAllProducts(nil, &wrongText, nil, nil, nil, nil, nil, nil, 0)
	assert.Equal(t, len(fetchedProducts), 0, "Fetching products with wrong title failed")

	fetchedProducts = repo.GetAllProducts(nil, nil, &wrongText, nil, nil, nil, nil, nil, 0)
	assert.Equal(t, len(fetchedProducts), 0, "Fetching products with wrong description failed")

	wrongInt := 50
	fetchedProducts = repo.GetAllProducts(nil, nil, nil, &wrongInt, nil, nil, nil, nil, 0)
	assert.Equal(t, len(fetchedProducts), 0, "Fetching products with wrong weight failed")

	wrongInt = 10
	fetchedProducts = repo.GetAllProducts(nil, nil, nil, nil, &wrongInt, nil, nil, nil, 0)
	assert.Equal(t, 0, len(fetchedProducts), "Fetching products with wrong weight failed")

	wrongPrice := 55.00
	fetchedProducts = repo.GetAllProducts(nil, nil, nil, nil, nil, &wrongPrice, nil, nil, 0)
	assert.Equal(t, len(fetchedProducts), 0, "Fetching products with wrong price failed")

	wrongPrice = 10.00
	fetchedProducts = repo.GetAllProducts(nil, nil, nil, nil, nil, nil, &wrongPrice, nil, 0)
	assert.Equal(t, len(fetchedProducts), 0, "Fetching products with wrong price failed")

	oneLimit := 1
	fetchedProducts = repo.GetAllProducts(nil, nil, nil, nil, nil, nil, nil, &oneLimit, 0)
	assert.Equal(t, len(fetchedProducts), oneLimit, "Fetching products with one limit failed")

	fetchedProducts = repo.GetAllProducts(nil, nil, nil, nil, nil, nil, nil, nil, 10)
	assert.Equal(t, len(fetchedProducts), oneLimit, "Fetching products with wrong offset failed")

	correctMinWeight := 15
	correctMaxWeight := 20
	correctMinPrice := 12.5
	correctMaxPrice := 20.00
	fetchedProducts = repo.GetAllProducts([]int{mockedProducts[0].Categories[0].Id}, &mockedProducts[0].Title, nil, &correctMinWeight, &correctMaxWeight, &correctMinPrice, &correctMaxPrice, nil, 0)
	assert.Equal(t, len(fetchedProducts), testingObjectCount, "Fetching products with correct data failed")
	assertProducts(t, mockedProducts, fetchedProducts)
}

// TestProductRepository_GetProduct functionality
func TestProductRepository_GetProduct(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up database connection failed")

	repo := createProductRepository(conn)

	mockedProducts := mockAndInsertProducts(conn, 1)
	assert.Equal(t, len(mockedProducts), 1, "Mocking products failed")
	defer destructCreatedType(conn, mockedProducts[0].Categories[0].TypeId)
	defer destructCreatedCategories(conn, mockedProducts[0].Categories)
	defer destructCreatedProducts(conn, mockedProducts)

	product := repo.GetProduct(mockedProducts[0].Id)
	assert.NotNil(t, product.Categories)
	assertProducts(t, mockedProducts, []Product{*product})
}

// TestProductRepository_UpdateProduct functionality
func TestProductRepository_UpdateProduct(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up database connection failed")

	repo := createProductRepository(conn)

	mockedProducts := mockAndInsertProducts(conn, 1)
	assert.Equal(t, len(mockedProducts), 1, "Mocking products failed")
	defer destructCreatedType(conn, mockedProducts[0].Categories[0].TypeId)
	defer destructCreatedCategories(conn, mockedProducts[0].Categories)
	defer destructCreatedProducts(conn, mockedProducts)

	newAmount := rand.Int()
	newTitle := "new test title for update"
	newCode := rand.Int()
	newWeight := rand.Int()

	mockedProducts[0].Amount = newAmount
	mockedProducts[0].Weight = &newWeight
	mockedProducts[0].Code = newCode
	mockedProducts[0].Title = newTitle

	err = repo.UpdateProduct(&mockedProducts[0])
	assert.NoError(t, err, "product update failed")

	tmpProduct := new(Product)
	conn.Where("id = ?", mockedProducts[0].Id).First(tmpProduct)
	assert.Equal(t, tmpProduct.Amount, newAmount, "product update failed")
	assert.Equal(t, *tmpProduct.Weight, newWeight, "product update failed")
	assert.Equal(t, tmpProduct.Title, newTitle, "product update failed")
	assert.Equal(t, tmpProduct.Code, newCode, "product update failed")
}

// TestLikeDislikeRepository_LikeProduct functionality
func TestLikeDislikeRepository_LikeProduct(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up database connection failed")

	repo := createLikeDislikeRepo(conn)

	productId := 1
	userId := 5

	result := repo.LikeProduct(productId, userId)
	assert.Equal(t, result.ProductId, productId, "Liking product failed")
	assert.Equal(t, result.UserId, userId, "Liking product failed")

	var tmpLikeProduct Likes
	err = conn.Where("user_id = ?", userId).Where("product_id = ?", productId).First(&tmpLikeProduct).Error
	assert.NoError(t, err, "Liking product failed")
}

// createCategoryRepository for testing purpose
func createCategoryRepository(db *gorm.DB) CategoriesRepositoryInterface {
	return NewCategoryRepo(db)
}

// setupDbConnection to testing in memory database
func setupDbConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(Type{}, Category{}, Product{}, Likes{}, DisLikes{})
	return db, err
}

// mockAndInsertCategories in database
func mockAndInsertCategories(db *gorm.DB, count int) []Category {
	customType := createTypeInDb(db)

	categories := make([]Category, 0, count)

	for i := 0; i < count; i++ {
		mockedCat := mockCategory(customType.Id)
		result := db.Create(mockedCat)
		if result.Error != nil {
			continue
		}
		categories = append(categories, *mockedCat)
	}

	return categories
}

// createTypeInDb mocks and insert a type in database
func createTypeInDb(db *gorm.DB) *Type {
	tmpType := mockType()
	db.Create(tmpType)
	return tmpType
}

// mockAndInsertType in temporary database
func mockAndInsertType(conn *gorm.DB, count int) []Type {
	types := make([]Type, 0, count)

	for i := 0; i < count; i++ {
		mockedType := mockType()
		result := conn.Create(mockedType)
		if result.Error != nil {
			log.Println(result.Error)
			continue
		}
		types = append(types, *mockedType)
	}

	return types
}

// mockType struct for test
func mockType() *Type {
	return &Type{
		Title: "Test default type",
	}
}

// mockCategory for test
func mockCategory(typeId int) *Category {
	return &Category{
		TypeId:         typeId,
		ParentCatId:    nil,
		ParentCategory: nil,
		Title:          "Test category",
		Indent:         1,
		Order:          1,
	}
}

// assertCategories check whether they are equal or not
func assertCategories(t *testing.T, createdCats, fetchedCats []Category) {
	for index := range createdCats {
		assert.Equal(t, createdCats[index].Id, fetchedCats[index].Id, "Categories Repository test: Ids are not equal")
		assert.Equal(t, createdCats[index].Title, fetchedCats[index].Title, "Categories Repository test: titles are not equal")
		assert.Equal(t, createdCats[index].TypeId, fetchedCats[index].TypeId, "Categories Repository test: type ids are not equal")
		assert.Equal(t, createdCats[index].Indent, fetchedCats[index].Indent, "Categories Repository test: indents are not equal")
		assert.Equal(t, createdCats[index].Order, fetchedCats[index].Order, "Categories Repository test: orders are not equal")
		assert.NotNil(t, fetchedCats[index].Type, "Categories Repository test: Type is not eager loaded properly")
	}
}

// destructCreatedCategories and delete them from db
func destructCreatedCategories(db *gorm.DB, cats []Category) {
	for _, cat := range cats {
		db.Unscoped().Delete(cat)
	}
}

// destructCreatedType and deleted it from DB
func destructCreatedType(db *gorm.DB, typeId int) {
	db.Unscoped().Delete(Type{}, typeId)
}

// createTypeRepo and return it for testing purpose
func createTypeRepo(conn *gorm.DB) TypeRepositoryInterface {
	return NewTypeRepo(conn)
}

// assertTypes to check whether they are equal or not !
func assertTypes(t *testing.T, fetchedTypes, mockedTypes []Type) {
	for index := range mockedTypes {
		assert.Equal(t, fetchedTypes[index].Title, mockedTypes[index].Title, "Types are not equal")
		assert.Equal(t, fetchedTypes[index].Id, mockedTypes[index].Id, "Types are not equal")
	}
}

// destructTypes which has been created in db
func destructAllTypes(conn *gorm.DB, types []Type) {
	for _, tmpType := range types {
		conn.Unscoped().Delete(tmpType)
	}
}

// createProductRepository and return it
func createProductRepository(db *gorm.DB) ProductsRepositoryInterface {
	return NewProductRepository(db)
}

// mockAndInsertProducts in database
func mockAndInsertProducts(db *gorm.DB, count int) []Product {
	cat := mockAndInsertCategories(db, 1)

	products := make([]Product, 0, count)

	for i := 0; i < count; i++ {
		mockedProduct := mockProduct(cat[0])
		result := db.Create(mockedProduct)
		if result.Error != nil {
			log.Println(result.Error)
			continue
		}

		products = append(products, *mockedProduct)
	}

	return products
}

// mockProduct and return it for testing purpose
func mockProduct(category Category) *Product {
	randCode := rand.Int()
	randAmount := rand.Int()
	weight := 15
	description := "Test description for you, Which I prepared"
	return &Product{
		Categories:  []Category{category},
		Title:       "Test title",
		Code:        randCode,
		Amount:      randAmount,
		Price:       12.5,
		Weight:      &weight,
		Description: &description,
	}
}

// assertProducts in test operation
func assertProducts(t *testing.T, mockedProducts, fetchedProducts []Product) {
	for index := range mockedProducts {
		assert.NotNil(t, fetchedProducts[index].Categories, "Fetched products doesnt have categories preloaded")
		assert.Equal(t, mockedProducts[index].Id, fetchedProducts[index].Id, "Fetched products are not equal")
		assert.Equal(t, mockedProducts[index].Title, fetchedProducts[index].Title, "Fetched products are not equal")
		assert.Equal(t, mockedProducts[index].Code, fetchedProducts[index].Code, "Fetched products are not equal")
		assert.Equal(t, mockedProducts[index].Amount, fetchedProducts[index].Amount, "Fetched products are not equal")
		assert.Equal(t, mockedProducts[index].Price, fetchedProducts[index].Price, "Fetched products are not equal")
		assert.Equal(t, mockedProducts[index].Weight, fetchedProducts[index].Weight, "Fetched products are not equal")
		assert.Equal(t, mockedProducts[index].Description, fetchedProducts[index].Description, "Fetched products are not equal")
	}
}

// destructCreatedProducts that inserted during test
func destructCreatedProducts(db *gorm.DB, products []Product) {
	for _, product := range products {
		db.Unscoped().Delete(product)
	}
}

// createLikeDislikeRepo and return it for testing purpose
func createLikeDislikeRepo(db *gorm.DB) LikeDislikeRepositoryInterface {
	return NewLikeDislikeRepository(db)
}
