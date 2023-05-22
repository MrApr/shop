package address

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

// TestAddressService_GetAllCities functionality
func TestAddressService_GetAllCities(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Stablishing Database connection failed")
	service := createAddressService(conn)

	_, err = service.GetAllCities()
	assert.Error(t, err, "Fetching cities from address service failed")
	assert.ErrorIs(t, err, NoCitiesFound, "Fetching cities from address service failed")

	mockedCities := mockAndInsertCity(conn, 2)
	defer destructCities(conn, mockedCities)

	cities, err := service.GetAllCities()
	assert.NoError(t, err, "Fetching cities from address service failed")
	assert.Equal(t, len(mockedCities), len(cities), "Counts of retrieved cities are not equal with created ones")

	assertCitiesEquivelance(t, mockedCities, cities)
}

// TestAddressService_GetAllUserAddresses functionality
func TestAddressService_GetAllUserAddresses(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Stablishing Database connection failed")
	service := createAddressService(conn)

	city := mockAndInsertCity(conn, 1)
	defer destructCities(conn, city)

	userId := rand.Int()
	_, err = service.GetAllUserAddresses(userId)
	assert.Error(t, err, "Fetching user addresses for user from address service failed")
	assert.ErrorIs(t, err, NoAddressesFound, "Fetching addresses for user from address service failed")

	mockedAddresses := mockAndInsertAddresses(conn, city[0].Id, userId, 2)
	defer destructAddresses(conn, mockedAddresses)

	result, err := service.GetAllUserAddresses(userId)
	assert.NoError(t, err, "Fetching all user addresses failed in user service")
	assertAddresses(t, result, mockedAddresses)
}

// TestAddressService_CreateAddress functionality
func TestAddressService_CreateAddress(t *testing.T) {

}

// TestAddressService_UpdateAddress functionality
func TestAddressService_UpdateAddress(t *testing.T) {

}

// TestAddressService_DeleteAddress functionality
func TestAddressService_DeleteAddress(t *testing.T) {

}

// createAddressService and return it
func createAddressService(conn *gorm.DB) AddressServiceInterface {
	return NewAddressService(NewAddressRepository(conn))
}
