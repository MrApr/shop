package contact_us

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

// TestContactUsRepository_CreateContactUs functionality
func TestContactUsRepository_CreateContactUs(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up database connection failed")

	repo := createRepo(conn)
	mockedData := mockContactUs()

	result := repo.CreateContactUs(mockedData)
	assert.NoError(t, result, "Contact us creation failed")

	var tmpContactUs ContactUs
	conn.First(&tmpContactUs)

	assert.NotZero(t, tmpContactUs.Id, "Contact us creation failed")
	assert.Equal(t, tmpContactUs.Title, mockedData.Title, "Contact us creation failed")
	assert.Equal(t, tmpContactUs.Description, mockedData.Description, "Contact us creation failed")
	assert.Equal(t, tmpContactUs.Email, mockedData.Email, "Contact us creation failed")
}

// createRepo and return it for testing purpose
func createRepo(db *gorm.DB) ContactUsRepositoryInterface {
	return NewContactUsRepository(db)
}

// setupDbConnection to testing in memory database
func setupDbConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(ContactUs{})
	return db, err
}

// mockContactUs and return it for test
func mockContactUs() *ContactUs {
	return &ContactUs{
		Email:       "test@testom",
		Title:       "Test contact us",
		Description: "lorem ipsum and ....lorem ipsum and ....lorem ipsum and ....lorem ipsum and ....",
	}
}
