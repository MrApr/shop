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

// TestGatewayRepository_GetAllGateways functionality
func TestGatewayRepository_GetAllGateways(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setting up database connection failed")

	repo := createGatewayRepo(db)
	testingObjCount := 2

	mockedTypes := mockAndInsertGatewayTypes(db, 1, false)
	defer destructMockedTypes(db, mockedTypes)

	mockedGateways := mockAndInsertGateways(db, testingObjCount, mockedTypes[0].Id, false)
	defer destructMockedGateways(db, mockedGateways)
	assert.Equal(t, len(mockedGateways), testingObjCount, "Mocked objects length is not as equal as required and expected")

	result := repo.GetAllGateways(mockedTypes[0].Id, true)
	assert.Equal(t, len(result), 0, "expected no result because requested for only active gateways")

	result = repo.GetAllGateways(mockedTypes[0].Id, false)
	assertGateways(t, mockedGateways, result)
}

// createGatewayTypeRepo and return it
func createGatewayTypeRepo(db *gorm.DB) GatewayTypesRepositoryInterface {
	return NewGatewayTypeRepo(db)
}

// createGatewayRepo and return it
func createGatewayRepo(db *gorm.DB) GatewayRepositoryInterface {
	return NewGatewayRepository(db)
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

// mockAndInsertGateways into database for testing
func mockAndInsertGateways(db *gorm.DB, count, typeId int, status bool) []GateWay {
	gateways := make([]GateWay, 0, count)

	for i := 0; i < count; i++ {
		mockedGateway := mockGateway(typeId, status)
		result := db.Create(mockedGateway)
		if result.Error != nil {
			continue
		}

		gateways = append(gateways, *mockedGateway)
	}

	return gateways
}

// mockGateway and return it
func mockGateway(typeId int, status bool) *GateWay {
	return &GateWay{
		Name:          "Test gateway",
		GatewayTypeId: typeId,
		Token:         "Test token for gateways",
		Status:        status,
	}
}

// assertGatewayTypes and check whether they are equal or not
func assertGatewayTypes(t *testing.T, mockedTypes, fetchedTypes []GatewayType) {
	for index := range mockedTypes {
		assert.Equal(t, mockedTypes[index].Id, fetchedTypes[index].Id, "mocked types and fetched types are not equal")
		assert.Equal(t, mockedTypes[index].Title, fetchedTypes[index].Title, "mocked types and fetched types are not equal")
		assert.Equal(t, mockedTypes[index].Status, fetchedTypes[index].Status, "mocked types and fetched types are not equal")
		assert.True(t, fetchedTypes[index].Status, "Only active types should got received")
	}
}

// assertGateways and check whether they are equal or not
func assertGateways(t *testing.T, mockedGateways, fetchedGateways []GateWay) {
	for index := range mockedGateways {
		assert.Equal(t, mockedGateways[index].Id, fetchedGateways[index].Id, "mocked and fetched gateways are not equal")
		assert.Equal(t, mockedGateways[index].Name, fetchedGateways[index].Name, "mocked and fetched gateways are not equal")
		assert.Equal(t, mockedGateways[index].Token, fetchedGateways[index].Token, "mocked and fetched gateways are not equal")
		assert.Equal(t, mockedGateways[index].GatewayTypeId, fetchedGateways[index].GatewayTypeId, "mocked and fetched gateways are not equal")
		assert.Equal(t, mockedGateways[index].Status, fetchedGateways[index].Status, "mocked and fetched gateways are not equal")
	}
}

// destructMockedTypes which are created during testing
func destructMockedTypes(db *gorm.DB, gatewayTypes []GatewayType) {
	for _, gatewayType := range gatewayTypes {
		db.Unscoped().Delete(gatewayType)
	}
}

// destructMockedGateways which are created during testing
func destructMockedGateways(db *gorm.DB, gateways []GateWay) {
	for _, gateway := range gateways {
		db.Unscoped().Delete(gateway)
	}
}
