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

// createUseCase and return it for testing
func createUseCase(db *gorm.DB) CategoryUseCaseInterface {
	return NewCategoryUseCase(NewCategoryService(NewCategoryRepo(db)))
}
