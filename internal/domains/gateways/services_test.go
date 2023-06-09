package gateways

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

// TestGatewayTypeService_GetAllTypes functionality
func TestGatewayTypeService_GetAllTypes(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setting up database connection failed")

	sv := createGatewayTypeService(db)

	_, err = sv.GetAllTypes()
	assert.Error(t, err, "Expected error on empty types in db")
	assert.ErrorIs(t, err, NoTypesFound, "Expected error on empty types in db")

	testingObjCount := 2

	mockedTypes := mockAndInsertGatewayTypes(db, testingObjCount, true)
	defer destructMockedTypes(db, mockedTypes)
	assert.Equal(t, len(mockedTypes), testingObjCount, "Mocked objects length is not as equal as required and expected")

	result, err := sv.GetAllTypes()
	assert.NoError(t, err, "Fetching all types from sv failed")
	assert.Equal(t, len(result), testingObjCount, "Fetched objects length is not as equal as expected")

	assertGatewayTypes(t, mockedTypes, result)
}

// createGatewayTypeService and return it for testing
func createGatewayTypeService(db *gorm.DB) GatewayTypesServiceInterface {
	return NewGatewayTypesService(NewGatewayTypeRepo(db))
}
