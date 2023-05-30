package contact_us

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

// TestContactUsUseCase_CreateContactUs functionality
func TestContactUsUseCase_CreateContactUs(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Setting up database connection failed")

	ctx := context.Background()
	uc := createUseCase(conn)
	mockedRequest := mockCreateContactUsRequest()

	contactUs, err := uc.CreateContactUs(ctx, mockedRequest)
	assert.NoError(t, err, "Contact us creation in contact us service failed")

	assert.NotZero(t, contactUs.Id, "Contact us creation in contact us service failed")
	assert.Equal(t, contactUs.Title, mockedRequest.Title, "Contact us creation in contact us service failed")
	assert.Equal(t, contactUs.Description, mockedRequest.Description, "Contact us creation in contact us service failed")
	assert.Equal(t, contactUs.Email, mockedRequest.Email, "Contact us creation in contact us service failed")
}

// createUseCase and return it for testing purpose
func createUseCase(db *gorm.DB) ContactUsUseCaseInterface {
	return NewUseCase(NewContactUsService(NewContactUsRepository(db)))
}

// CreateContactUsRequest and return it
func mockCreateContactUsRequest() *CreateContactUsRequest {
	return &CreateContactUsRequest{
		Email:       "test@test.com",
		Title:       "test contact us title",
		Description: "Lorem ipsum ... Lorem ipsum ... Lorem ipsum ... Lorem ipsum ... Lorem ipsum ... Lorem ipsum ...",
	}
}
