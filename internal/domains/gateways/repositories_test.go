package gateways

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

// TestGatewayTypesRepository_GetAllTypes functionality
func TestGatewayTypesRepository_GetAllTypes(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setting up database connection failed")

	repo := createGatewayTypeRepo(db)
	testingObjCount := 2
	mockedTypes := mockAndInsertGatewayTypes(db, testingObjCount, false)
	defer destructMockedTypes(db, mockedTypes)
	assert.Equal(t, len(mockedTypes), testingObjCount, "Mocked objects length is not as equal as required and expected")

	result := repo.GetAllTypes()
	assert.Equal(t, len(result), 0, "repository should not return disabled mocked objects")

	mockedTypes = mockAndInsertGatewayTypes(db, testingObjCount, true)
	defer destructMockedTypes(db, mockedTypes)
	assert.Equal(t, len(mockedTypes), testingObjCount, "Mocked objects length is not as equal as required and expected")

	result = repo.GetAllTypes()
	assert.Equal(t, len(result), testingObjCount, "Fetched objects length is not as equal as expected")

	assertGatewayTypes(t, mockedTypes, result)
}

// createGatewayTypeRepo and return it
func createGatewayTypeRepo(db *gorm.DB) GatewayTypesRepositoryInterface {
	return NewGatewayTypeRepo(db)
}

// setupDbConnection and run migration
func setupDbConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(GatewayType{}, GateWay{})
	return db, err
}

// mockAndInsertGatewayTypes into temporary database
func mockAndInsertGatewayTypes(db *gorm.DB, count int, status bool) []GatewayType {
	gatewayTypes := make([]GatewayType, 0, count)

	for i := 0; i < count; i++ {
		mockedType := mockGatewayType(status)
		result := db.Create(mockedType)
		if result.Error != nil {
			continue
		}

		gatewayTypes = append(gatewayTypes, *mockedType)
	}

	return gatewayTypes
}

// mockGatewayType and return it
func mockGatewayType(status bool) *GatewayType {
	return &GatewayType{
		Title:  "test type title",
		Status: status,
	}
}

// assertGatewayTypes and check whether they are equal or not
func assertGatewayTypes(t *testing.T, mockedTypes, fetchedTypes []GatewayType) {
	for index := range mockedTypes {
		assert.Equal(t, mockedTypes[index].Id, fetchedTypes[index].Id, "mocked types and fetched types are not equal")
		assert.Equal(t, mockedTypes[index].Title, fetchedTypes[index].Title, "mocked types and fetched types are not equal")
		assert.Equal(t, mockedTypes[index].Status, fetchedTypes[index].Status, "mocked types and fetched types are not equal")
		assert.True(t, fetchedTypes[index].Status, "Only active types should get received")
	}
}

// destructMockedTypes which are created during testing
func destructMockedTypes(db *gorm.DB, gatewayTypes []GatewayType) {
	for _, gatewayType := range gatewayTypes {
		db.Unscoped().Delete(gatewayType)
	}
}
