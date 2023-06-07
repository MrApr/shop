package postType

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"testing"
	"time"
)

// TestPostTypeRepository_GetAllPostTypes functionality
func TestPostTypeRepository_GetAllPostTypes(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up temporary database connection failed")

	repo := createPostTypeRepository(conn)
	testingObjsCount := 2

	mockedPostTypes := mockAndInsertPostTypes(conn, testingObjsCount)
	assert.Equal(t, len(mockedPostTypes), testingObjsCount, "Mocked objects are not enough as expected")

	fetchedTypes := repo.GetAllPostTypes()
	assert.Equal(t, len(fetchedTypes), len(mockedPostTypes), "Mocked objects and fetched objects length is not equal")

	assertPostTypes(t, mockedPostTypes, fetchedTypes)
}

// TestPostTypeRepository_PostTypeExists functionality
func TestPostTypeRepository_PostTypeExists(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up temporary database connection failed")

	repo := createPostTypeRepository(conn)
	testingObjsCount := 1

	mockedPostTypes := mockAndInsertPostTypes(conn, testingObjsCount)
	assert.Equal(t, len(mockedPostTypes), testingObjsCount, "Mocked objects are not enough as expected")

	exists := repo.PostTypeExists(mockedPostTypes[0].Id)
	assert.True(t, exists, "Existence method of repository is not correct, expected true")

	randWrongId := rand.Int()
	exists = repo.PostTypeExists(randWrongId)
	assert.False(t, exists, "Existence method of repository is not correct, expected False")
}

// setupDbConnection and run migration
func setupDbConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(PostType{})
	return db, err
}

// createPostTypeRepository and return it for testing purpose
func createPostTypeRepository(db *gorm.DB) PostTypeRepositoryInterface {
	return NewRepository(db)
}

// mockAndInsertPostTypes in temporary db
func mockAndInsertPostTypes(db *gorm.DB, count int) []PostType {
	postTypes := make([]PostType, 0, count)

	for i := 0; i < count; i++ {
		mockedPostType := mockPostType()

		result := db.Create(mockedPostType)
		if result.Error != nil {
			continue
		}

		postTypes = append(postTypes, *mockedPostType)
	}

	return postTypes
}

// mockPostType and return it
func mockPostType() *PostType {
	randPrice := rand.Float64()
	return &PostType{
		Title:           "Test post type",
		Price:           randPrice,
		DeliverableTime: time.Now(),
	}
}

// destructPostTypes which are already created in db
func destructPostTypes(db *gorm.DB, postTypes []PostType) {
	for _, pType := range postTypes {
		db.Unscoped().Delete(pType)
	}
}

// assertPostTypes and check they are equal or not
func assertPostTypes(t *testing.T, mockedPostType, fetchedPostType []PostType) {
	for index := range mockedPostType {
		assert.Equal(t, mockedPostType[index].Id, fetchedPostType[index].Id, "Mocked and fetched post types are not equal in id field")
		assert.Equal(t, mockedPostType[index].Title, fetchedPostType[index].Title, "Mocked and fetched post types are not equal in title field")
		assert.Equal(t, mockedPostType[index].Price, fetchedPostType[index].Price, "Mocked and fetched post types are not equal in price field")
	}
}
