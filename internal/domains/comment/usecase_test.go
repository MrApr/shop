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
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setting up database connection failed")

	uC := createUseCase(db)
	ctx := context.Background()
	pId := rand.Int()

	mockedCreateRequest := mockCreateCommentRequest(pId)

	createdComment, err := uC.CreateComment(ctx, "", mockedCreateRequest)
	assert.NoError(t, err, "Expected no error comment creation")
	assert.Equal(t, 1, createdComment.UserId, "comment creation failed")
	assert.Equal(t, mockedCreateRequest.ProductId, createdComment.ProductId, "comment creation failed")
	assert.Equal(t, mockedCreateRequest.Description, createdComment.Description, "comment creation failed")
	assert.True(t, createdComment.Status, "Comment creation default status is wrong")
}

// TestCommentUseCase_DeleteComment functionality
func TestCommentUseCase_DeleteComment(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setting up database connection failed")

	uC := createUseCase(db)
	testingCount := 1
	ctx := context.Background()

	mockedCm := mockAndInsertComments(db, testingCount, 0, true)
	defer destructComments(db, mockedCm)
	assert.Equal(t, testingCount, len(mockedCm), "Mocked comments length is not equal as expected")

	err = uC.DeleteComment(ctx, mockedCm[0].Id)
	assert.NoError(t, err, "deleting comment failed")

	tmpCm := new(Comment)
	db.Where("id = ?", mockedCm[0].Id).First(tmpCm)
	assert.Zero(t, tmpCm.Id, "deleting comment failed")
}

// createUseCase and return it
func createUseCase(db *gorm.DB) CommentUseCaseInterface {
	return NewCommentUseCase(NewCommentService(NewCommentRepository(db)), func(ctx context.Context, token string) (int, error) {
		return 1, nil
	})
}

// mockCreateCommentRequest and return it
func mockCreateCommentRequest(pId int) *CreateCommentRequest {
	return &CreateCommentRequest{
		ProductId:   pId,
		Description: "Test product",
	}
}
