package postType

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"math/rand"
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
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up temporary database connection failed")

	sv := createPostTypeService(conn)
	testingObjsCount := 1

	mockedPostTypes := mockAndInsertPostTypes(conn, testingObjsCount)
	assert.Equal(t, len(mockedPostTypes), testingObjsCount, "Mocked objects are not enough as expected")

	err = sv.PostTypeExists(mockedPostTypes[0].Id)
	assert.NoError(t, err, "Post type exists, expected no error but received one")

	randWrongId := rand.Int()
	err = sv.PostTypeExists(randWrongId)
	assert.Error(t, err, "Post type with given id doesnt exist, expected error but received none")
	assert.ErrorIs(t, err, PostTypeDoesntExists, "Post type with given id doesnt exist, expected PostTypeDoesntExists error but received none")
}

// createPostTypeService and return it
func createPostTypeService(db *gorm.DB) PostTypeServiceInterface {
	return NewPostTypeService(NewRepository(db))
}
