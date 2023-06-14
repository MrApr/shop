package comment

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

// TestCommentUseCase_GetAllActiveComments functionality
func TestCommentUseCase_GetAllActiveComments(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setting up database connection failed")

	uC := createUseCase(db)
	testingCount := 2
	pId := rand.Int()
	ctx := context.Background()

	mockedCm := mockAndInsertComments(db, testingCount, pId, true)
	defer destructComments(db, mockedCm)
	assert.Equal(t, testingCount, len(mockedCm), "Mocked comments length is not equal as expected")

	fetchedCms, err := uC.GetAllActiveComments(ctx, pId)
	assert.NoError(t, err, "Fetching all active comments failed with error")
	assert.Equal(t, len(fetchedCms), testingCount, "Expected no active comments for product but received some")

	assertComments(t, fetchedCms, mockedCm)
}

// TestCommentUseCase_CreateComment functionality
func TestCommentUseCase_CreateComment(t *testing.T) {

}

// TestCommentUseCase_DeleteComment functionality
func TestCommentUseCase_DeleteComment(t *testing.T) {

}

// createUseCase and return it
func createUseCase(db *gorm.DB) CommentUseCaseInterface {
	return NewCommentUseCase(NewCommentService(NewCommentRepository(db)))
}
