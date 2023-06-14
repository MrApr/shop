package comment

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

// TestCommentService_GetAllActiveComments functionality
func TestCommentService_GetAllActiveComments(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setting up database connection failed")

	sv := createService(db)
	testingCount := 2
	pId := rand.Int()

	_, err = sv.GetAllActiveComments(pId)
	assert.Error(t, err, "Expected error on null comments")
	assert.ErrorIs(t, err, NoProductsFound, "Expected error on null comments")

	mockedCm := mockAndInsertComments(db, testingCount, pId, true)
	defer destructComments(db, mockedCm)
	assert.Equal(t, testingCount, len(mockedCm), "Mocked comments length is not equal as expected")

	fetchedCms, err := sv.GetAllActiveComments(pId)
	assert.NoError(t, err, "Fetching all active comments failed with error")
	assert.Equal(t, len(fetchedCms), testingCount, "Expected no active comments for product but received some")

	assertComments(t, fetchedCms, mockedCm)
}

// TestCommentService_CreateComment functionality
func TestCommentService_CreateComment(t *testing.T) {

}

// TestCommentService_DeleteComment functionality
func TestCommentService_DeleteComment(t *testing.T) {

}

// createService and return it for testing purpose
func createService(db *gorm.DB) CommentServiceInterface {
	return NewCommentService(NewCommentRepository(db))
}
