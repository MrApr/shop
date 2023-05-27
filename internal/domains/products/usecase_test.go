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

	_, err = uC.GetAllCategories(ctx)
	assert.Error(t, err, "Expected categories not found error")
	assert.ErrorIs(t, err, NoCategoriesFound, "Expected categories not found error")

	createdCats := mockAndInsertCategories(conn, 2)

	fetchedCategories, err := uC.GetAllCategories(ctx)
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
	assert.Equal(t, len(mockedTypes), testingObjectsCount, "Creating Mock objects failed")
	mockedRequest := mockGetAllTypesRequest()

	fetchedTypes, err := uc.GetAllTypes(ctx, mockedRequest)
	assert.NoError(t, err, "Fetching Types from type service:no error expected")
	assert.Equal(t, len(fetchedTypes), testingObjectsCount, "Fetching Types from repo failed")
	assertTypes(t, fetchedTypes, mockedTypes)
}

// createUseCase and return it for testing
func createUseCase(db *gorm.DB) CategoryUseCaseInterface {
	return NewCategoryUseCase(NewCategoryService(NewCategoryRepo(db)))
}

// createTypeUseCase and return it for testing purpose
func createTypeUseCase(db *gorm.DB) TypeUseCaseInterface {
	return NewTypeUseCase(NewTypeService(NewTypeRepo(db)))
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
