package comment

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

// TestCommentRepository_GetComment functionality
func TestCommentRepository_GetComment(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setting up database connection failed")

	repo := createRepository(db)

	mockedCm := mockAndInsertComments(db, 1, true)
	assert.Equal(t, 1, len(mockedCm), "Mocked comments length is not equal as expected")

	fetchedCm := repo.GetComment(mockedCm[0].Id)
	assertComments(t, []Comment{*fetchedCm}, mockedCm)
}

// TestCommentRepository_GetAllActiveComments functionality
func TestCommentRepository_GetAllActiveComments(t *testing.T) {

}

// TestCommentRepository_CreateComment functionality
func TestCommentRepository_CreateComment(t *testing.T) {

}

// TestCommentRepository_DeleteComment functionality
func TestCommentRepository_DeleteComment(t *testing.T) {

}

// createRepository and return it for testing purpose
func createRepository(db *gorm.DB) CommentRepositoryInterface {
	return NewCommentRepository(db)
}

// setupDbConnection and run migration
func setupDbConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(Comment{})
	return db, err
}

// mockAndInsertComments into database
func mockAndInsertComments(db *gorm.DB, count int, status bool) []Comment {
	comments := make([]Comment, 0, count)
	for i := 0; i < count; i++ {
		mockedCm := mockComment(status)
		result := db.Create(mockedCm)
		if result.Error != nil {
			continue
		}

		comments = append(comments, *mockedCm)
	}
	return comments
}

// mockComment and return it for testing purpose
func mockComment(status bool) *Comment {
	return &Comment{
		Id:          0,
		ProductId:   rand.Int(),
		UserId:      rand.Int(),
		Description: "Test description for it",
		Status:      status,
	}
}

// destructComments which are created in testing procedure
func destructComments(db *gorm.DB, comments []Comment) {
	for _, cm := range comments {
		db.Unscoped().Delete(cm)
	}
}

// assertComments and checks whether they are equal or not
func assertComments(t *testing.T, fetchedCms, mockedCms []Comment) {
	for index := range mockedCms {
		assert.Equal(t, mockedCms[index].Id, fetchedCms[index].Id, "mocked and fetched comments are not equal")
		assert.Equal(t, mockedCms[index].ProductId, fetchedCms[index].ProductId, "mocked and fetched comments are not equal")
		assert.Equal(t, mockedCms[index].UserId, fetchedCms[index].UserId, "mocked and fetched comments are not equal")
		assert.Equal(t, mockedCms[index].Description, fetchedCms[index].Description, "mocked and fetched comments are not equal")
		assert.Equal(t, mockedCms[index].Status, fetchedCms[index].Status, "mocked and fetched comments are not equal")
	}
}
