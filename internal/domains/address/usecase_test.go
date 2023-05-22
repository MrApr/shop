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
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Stablishing Database connection failed")
	ctx := context.Background()
	usecase := createUseCase(conn)

	city := mockAndInsertCity(conn, 1)
	defer destructCities(conn, city)

	userId := 1
	mockedRequest := mockCreateRequest(city[0].Id)

	result, err := usecase.CreateAddress(ctx, "", mockedRequest)
	assert.NoError(t, err, "User address creation failed in address use-case")

	assert.Equal(t, result.Address, mockedRequest.Address, "User address creation failed in address use-case")
	assert.Equal(t, result.CityId, mockedRequest.CityId, "User address creation failed in address use-case")
	assert.Equal(t, result.UserId, userId, "User address creation failed in address use-case")
}

// TestAddressUseCase_UpdateAddress functionality
func TestAddressUseCase_UpdateAddress(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Stablishing Database connection failed")
	ctx := context.Background()
	usecase := createUseCase(conn)

	cities := mockAndInsertCity(conn, 2)
	defer destructCities(conn, cities)

	userId := 1
	mockedAddresses := mockAndInsertAddresses(conn, cities[0].Id, userId, 1)
	defer destructAddresses(conn, mockedAddresses)

	mockedEditRequest := mockEditRequest(cities[1].Id, mockedAddresses[0].Id)

	editedAddress, err := usecase.UpdateAddress(ctx, "", mockedEditRequest)
	assert.NoError(t, err, "UserAddress use-case update functionality failed")
	assert.Equal(t, editedAddress.Id, mockedEditRequest.AddressId, "UserAddress use-case update functionality failed")
	assert.Equal(t, editedAddress.Address, mockedEditRequest.Address, "UserAddress use-case update functionality failed")
	assert.Equal(t, editedAddress.CityId, mockedEditRequest.CityId, "UserAddress use-case update functionality failed")
}

// TestAddressUseCase_DeleteAddress functionality
func TestAddressUseCase_DeleteAddress(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Stablishing Database connection failed")
	ctx := context.Background()
	usecase := createUseCase(conn)

	cities := mockAndInsertCity(conn, 1)
	defer destructCities(conn, cities)

	userId := 1
	mockedAddresses := mockAndInsertAddresses(conn, cities[0].Id, userId, 1)
	defer destructAddresses(conn, mockedAddresses)

	mockedDeleteRequest := mockDeleteRequest(mockedAddresses[0].Id)

	err = usecase.DeleteAddress(ctx, "", mockedDeleteRequest)
	assert.NoError(t, err, "UserAddress use-case delete functionality failed")
}

// createUseCase and return it
func createUseCase(conn *gorm.DB) AddressUseCaseInterface {
	return NewUseCase(NewAddressService(NewAddressRepository(conn)), func(ctx context.Context, token string) (int, error) {
		return 1, nil
	})
}

// mockCreateRequest for address creation
func mockCreateRequest(cityId int) *CreateAddressRequest {
	return &CreateAddressRequest{
		CityId:  cityId,
		Address: "Test address",
	}
}

// mockEditRequest for update operation for addresses
func mockEditRequest(cityId, addressId int) *UpdateAddressRequest {
	return &UpdateAddressRequest{
		AddressId: addressId,
		CityId:    cityId,
		Address:   "new edited address",
	}
}

// mockDeleteRequest for deleting adresses operation
func mockDeleteRequest(addrId int) *DeleteAddressRequest {
	return &DeleteAddressRequest{
		AddressId: addrId,
	}
}
