package products

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

// TestCategoryService_GetAllCategories functionality
func TestCategoryService_GetAllCategories(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting database connection up failed")

	sv := createService(conn)

	_, err = sv.GetAllCategories(nil, nil, nil, nil, 0)
	assert.Error(t, err, "Expected categories not found error")
	assert.ErrorIs(t, err, NoCategoriesFound, "Expected categories not found error")

	createdCats := mockAndInsertCategories(conn, 2)
	defer destructCreatedType(conn, createdCats[0].TypeId)
	defer destructCreatedCategories(conn, createdCats)

	fetchedCategories, err := sv.GetAllCategories(nil, nil, nil, nil, 0)
	assert.NoError(t, err, "Fetching Categories from db failed")
	assertCategories(t, createdCats, fetchedCategories)
}

// TestTypeService_GetAllTypes functionality
func TestTypeService_GetAllTypes(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting database connection up failed")
	testingObjectsCount := 5

	sv := createTypeService(conn)

	_, err = sv.GetAllTypes(nil, testingObjectsCount, 0)
	assert.Error(t, err, "Fetching Types from type service: expected error")
	assert.ErrorIs(t, err, NoTypesFound, "Fetching Types from type service: provided wrong error type")

	mockedTypes := mockAndInsertType(conn, testingObjectsCount)
	defer destructAllTypes(conn, mockedTypes)
	assert.Equal(t, len(mockedTypes), testingObjectsCount, "Creating Mock objects failed")

	fetchedTypes, err := sv.GetAllTypes(nil, testingObjectsCount, 0)
	assert.NoError(t, err, "Fetching Types from type service:no error expected")
	assert.Equal(t, len(fetchedTypes), testingObjectsCount, "Fetching Types from repo failed")
	assertTypes(t, fetchedTypes, mockedTypes)
}

// createService and return it for testing purpose
func createService(db *gorm.DB) CategoryServiceInterface {
	return NewCategoryService(NewCategoryRepo(db))
}

// createTypeService and return it for testing purpose
func createTypeService(db *gorm.DB) TypeServiceInterface {
	return NewTypeService(NewTypeRepo(db))
}
