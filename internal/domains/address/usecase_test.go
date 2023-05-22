package address

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

// TestAddressUseCase_GetAllCities functionality
func TestAddressUseCase_GetAllCities(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Stablishing Database connection failed")
	ctx := context.Background()
	usecase := createUseCase(conn)

	mockedCities := mockAndInsertCity(conn, 2)
	defer destructCities(conn, mockedCities)

	cities, err := usecase.GetAllCities(ctx)
	assert.NoError(t, err, "Fetching cities from address usecae failed")
	assert.Equal(t, len(mockedCities), len(cities), "Counts of retrieved cities are not equal with created ones")

	assertCitiesEquivelance(t, mockedCities, cities)

}

// TestAddressUseCase_GetAllUserAddresses functionality
func TestAddressUseCase_GetAllUserAddresses(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Stablishing Database connection failed")
	ctx := context.Background()
	usecase := createUseCase(conn)

	city := mockAndInsertCity(conn, 1)
	defer destructCities(conn, city)

	userId := 1
	mockedAddresses := mockAndInsertAddresses(conn, city[0].Id, userId, 2)
	defer destructAddresses(conn, mockedAddresses)

	result, err := usecase.GetAllUserAddresses(ctx, "")
	assert.NoError(t, err, "Fetching all user addresses failed in user usecase")
	assertAddresses(t, result, mockedAddresses)
}

// TestAddressUseCase_CreateAddress functionality
func TestAddressUseCase_CreateAddress(t *testing.T) {

}

// TestAddressUseCase_UpdateAddress functionality
func TestAddressUseCase_UpdateAddress(t *testing.T) {

}

// TestAddressUseCase_DeleteAddress functionality
func TestAddressUseCase_DeleteAddress(t *testing.T) {

}

// createUseCase and return it
func createUseCase(conn *gorm.DB) AddressUseCaseInterface {
	return NewUseCase(NewAddressService(NewAddressRepository(conn)), func(ctx context.Context, token string) (int, error) {
		return 1, nil
	})
}
