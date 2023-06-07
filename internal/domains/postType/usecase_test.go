package postType

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

// TestPostTypeUseCase_GetAllPostTypes functionality
func TestPostTypeUseCase_GetAllPostTypes(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up temporary database connection failed")

	uC := createPostTypeUseCase(conn)
	ctx := context.Background()

	testingObjsCount := 2
	mockedPostTypes := mockAndInsertPostTypes(conn, testingObjsCount)
	assert.Equal(t, len(mockedPostTypes), testingObjsCount, "Mocked objects are not enough as expected")

	fetchedTypes, err := uC.GetAllPostTypes(ctx)
	assert.NoError(t, err, "Fetching all post types from post type service failed. No error expected")
	assert.Equal(t, len(fetchedTypes), len(mockedPostTypes), "Mocked objects and fetched objects length is not equal")

	assertPostTypes(t, mockedPostTypes, fetchedTypes)
}

// createPostTypeUseCase and return it
func createPostTypeUseCase(db *gorm.DB) PostTypeUseCaseInterface {
	return NewUseCase(NewPostTypeService(NewRepository(db)))
}
