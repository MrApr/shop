package postType

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

// TestPostTypeService_GetAllPostTypes functionality
func TestPostTypeService_GetAllPostTypes(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up temporary database connection failed")

	sv := createPostTypeService(conn)

	_, err = sv.GetAllPostTypes()
	assert.Error(t, err, "Expected error on empty post types received none")
	assert.ErrorIs(t, err, NotTypesFound, "Expected NotTypesFound error on empty post types received none")

	testingObjsCount := 2
	mockedPostTypes := mockAndInsertPostTypes(conn, testingObjsCount)
	assert.Equal(t, len(mockedPostTypes), testingObjsCount, "Mocked objects are not enough as expected")

	fetchedTypes, err := sv.GetAllPostTypes()
	assert.NoError(t, err, "Fetching all post types from post type service failed. No error expected")
	assert.Equal(t, len(fetchedTypes), len(mockedPostTypes), "Mocked objects and fetched objects length is not equal")

	assertPostTypes(t, mockedPostTypes, fetchedTypes)
}

// TestPostTypeService_PostTypeExists functionality
func TestPostTypeService_PostTypeExists(t *testing.T) {

}

// createPostTypeService and return it
func createPostTypeService(db *gorm.DB) PostTypeServiceInterface {
	return NewPostTypeService(NewRepository(db))
}
