package gateways

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

// TestGatewayTypeUseCase_GetAllTypes functionality
func TestGatewayTypeUseCase_GetAllTypes(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setting up database connection failed")

	uC := createGatewayTypeUseCase(db)
	testingObjCount := 2
	ctx := context.Background()

	mockedTypes := mockAndInsertGatewayTypes(db, testingObjCount, true)
	defer destructMockedTypes(db, mockedTypes)
	assert.Equal(t, len(mockedTypes), testingObjCount, "Mocked objects length is not as equal as required and expected")

	result, err := uC.GetAllTypes(ctx)
	assert.NoError(t, err, "Fetching all types from sv failed")
	assert.Equal(t, len(result), testingObjCount, "Fetched objects length is not as equal as expected")

	assertGatewayTypes(t, mockedTypes, result)
}

// createGatewayTypeUseCase and return it
func createGatewayTypeUseCase(db *gorm.DB) GatewayTypesUseCaseInterface {
	return NewGatewayTypeUseCase(NewGatewayTypesService(NewGatewayTypeRepo(db)))
}
