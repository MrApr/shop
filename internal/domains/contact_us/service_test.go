package contact_us

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

// TestContactUsService_CreateContactUs functionality
func TestContactUsService_CreateContactUs(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up database connection failed")

	sv := createService(conn)
	mockedData := mockContactUs()

	contactUs, err := sv.CreateContactUs(mockedData.Title, mockedData.Email, mockedData.Description)
	assert.NoError(t, err, "Contact us creation in contact us service failed")

	assert.NotZero(t, contactUs.Id, "Contact us creation in contact us service failed")
	assert.Equal(t, contactUs.Title, mockedData.Title, "Contact us creation in contact us service failed")
	assert.Equal(t, contactUs.Description, mockedData.Description, "Contact us creation in contact us service failed")
	assert.Equal(t, contactUs.Email, mockedData.Email, "Contact us creation in contact us service failed")
}

// createService and return it for testing purpose
func createService(db *gorm.DB) ContactUsServiceInterface {
	return NewContactUsService(NewContactUsRepository(db))
}
