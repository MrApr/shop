package products

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

// TestCategoriesUseCase_GetAllCategories functionality
func TestCategoriesUseCase_GetAllCategories(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting database connection up failed")

	ctx := context.Background()
	uC := createUseCase(conn)
	mockGetAllRequest := mockEmptyAllCategoriesRequest()

	_, err = uC.GetAllCategories(ctx, mockGetAllRequest)
	assert.Error(t, err, "Expected categories not found error")
	assert.ErrorIs(t, err, NoCategoriesFound, "Expected categories not found error")

	createdCats := mockAndInsertCategories(conn, 2)
	defer destructCreatedType(conn, createdCats[0].TypeId)
	defer destructCreatedCategories(conn, createdCats)

	fetchedCategories, err := uC.GetAllCategories(ctx, mockGetAllRequest)
	assert.NoError(t, err, "Fetching Categories from db failed")
	assertCategories(t, createdCats, fetchedCategories)
}

// TestTypesUseCase_GetAllTypes functionality
func TestTypesUseCase_GetAllTypes(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting database connection up failed")
	testingObjectsCount := 5
	ctx := context.Background()

	uc := createTypeUseCase(conn)

	mockedTypes := mockAndInsertType(conn, testingObjectsCount)
	defer destructAllTypes(conn, mockedTypes)
	assert.Equal(t, len(mockedTypes), testingObjectsCount, "Creating Mock objects failed")
	mockedRequest := mockGetAllTypesRequest()

	fetchedTypes, err := uc.GetAllTypes(ctx, mockedRequest)
	assert.NoError(t, err, "Fetching Types from type service:no error expected")
	assert.Equal(t, len(fetchedTypes), testingObjectsCount, "Fetching Types from repo failed")
	assertTypes(t, fetchedTypes, mockedTypes)
}

// TestProductUseCase_GetAllProducts functionality
func TestProductUseCase_GetAllProducts(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up database connection failed")
	ctx := context.Background()

	testingObjectCount := 5

	uc := createProductUseCase(conn)
	mockedGetAllRequest := mockEmptyGetAllProductsRequest()

	mockedProducts := mockAndInsertProducts(conn, testingObjectCount)
	assert.Equal(t, len(mockedProducts), testingObjectCount, "Mocking products failed")
	defer destructCreatedType(conn, mockedProducts[0].Categories[0].TypeId)
	defer destructCreatedCategories(conn, mockedProducts[0].Categories)
	defer destructCreatedProducts(conn, mockedProducts)

	products, err := uc.GetAllProducts(ctx, mockedGetAllRequest)
	assert.NoError(t, err, "Fetching products from service failed")
	assertProducts(t, mockedProducts, products)
}

// TestProductUseCase_GetProduct functionality
func TestProductUseCase_GetProduct(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up database connection failed")
	ctx := context.Background()

	uc := createProductUseCase(conn)

	mockedProducts := mockAndInsertProducts(conn, 1)
	assert.Equal(t, len(mockedProducts), 1, "Mocking products failed")
	defer destructCreatedType(conn, mockedProducts[0].Categories[0].TypeId)
	defer destructCreatedCategories(conn, mockedProducts[0].Categories)
	defer destructCreatedProducts(conn, mockedProducts)

	product, err := uc.GetProduct(ctx, mockedProducts[0].Id)
	assert.NotNil(t, product.Categories)
	assertProducts(t, mockedProducts, []Product{*product})
}

// createUseCase and return it for testing
func createUseCase(db *gorm.DB) CategoryUseCaseInterface {
	return NewCategoryUseCase(NewCategoryService(NewCategoryRepo(db)))
}

// createTypeUseCase and return it for testing purpose
func createTypeUseCase(db *gorm.DB) TypeUseCaseInterface {
	return NewTypeUseCase(NewTypeService(NewTypeRepo(db)))
}

// createProductUseCase and return it
func createProductUseCase(db *gorm.DB) ProductUseCaseInterface {
	return NewProductUseCase(NewProductsService(NewProductRepository(db)))
}

// GetAllTypesRequest and return it
func mockGetAllTypesRequest() *GetAllTypesRequest {
	offset := 0
	return &GetAllTypesRequest{
		Name:   nil,
		Limit:  5,
		Offset: &offset,
	}
}

// mockEmptyAllCategoriesRequest and return it for testing purpose
func mockEmptyAllCategoriesRequest() *GetAllCategoriesRequest {
	return &GetAllCategoriesRequest{
		Title:       nil,
		TypeId:      nil,
		ParentCatId: nil,
		Limit:       nil,
		Offset:      0,
	}
}

// mockEmptyGetAllProductsRequest for testing purpose
func mockEmptyGetAllProductsRequest() *GetAllProductsRequest {
	return &GetAllProductsRequest{
		CategoryIds: nil,
		Title:       nil,
		Description: nil,
		MinWeight:   nil,
		MaxWeight:   nil,
		MinPrice:    nil,
		MaxPrice:    nil,
		Limit:       nil,
		Offset:      0,
	}
}
