package address

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

// TestAddressRepository_GetAllCities functionality
func TestAddressRepository_GetAllCities(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Stablishing Database connection failed")
	repo := createRepository(conn)

	mockedCities := mockAndInsertCity(conn, 2)
	defer destructCities(conn, mockedCities)

	cities, err := repo.GetAllCities()
	assert.NoError(t, err, "Fetching cities from repository failed")
	assert.Equal(t, len(mockedCities), len(cities), "Counts of retrieved cities are not equal with created ones")

	assertCitiesEquivelance(t, mockedCities, cities)
}

// TestAddressRepository_CityExists functionality
func TestAddressRepository_CityExists(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Stablishing Database connection failed")
	repo := createRepository(conn)

	mockedCities := mockAndInsertCity(conn, 1)
	defer destructCities(conn, mockedCities)

	exists := repo.CityExists(mockedCities[0].Id)
	assert.True(t, exists, "checking city existence from repository failed")

	exists = repo.CityExists(rand.Int())
	assert.False(t, exists, "checking city existence from repository failed")
}

// TestAddressRepository_GetAllUserAddresses functionality
func TestAddressRepository_GetAllUserAddresses(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Stablishing Database connection failed")
	repo := createRepository(conn)

	city := mockAndInsertCity(conn, 1)
	defer destructCities(conn, city)

	userId := rand.Int()
	mockedAddresses := mockAndInsertAddresses(conn, city[0].Id, userId, 2)
	defer destructAddresses(conn, mockedAddresses)

	result, err := repo.GetAllUserAddresses(userId)
	assert.NoError(t, err, "Fetching all user addresses failed")
	assertAddresses(t, result, mockedAddresses)
}

// TestAddressRepository_GetAddressById functionality
func TestAddressRepository_GetAddressById(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Stablishing Database connection failed")
	repo := createRepository(conn)

	city := mockAndInsertCity(conn, 1)
	defer destructCities(conn, city)

	mockedAddresses := mockAndInsertAddresses(conn, city[0].Id, 0, 1)
	defer destructAddresses(conn, mockedAddresses)

	result, err := repo.GetAddressById(mockedAddresses[0].Id)
	assert.NoError(t, err, "Fetching Address by id failed")
	assertAddresses(t, []Address{*result}, mockedAddresses)
}

// TestAddressRepository_CreateAddress functionality
func TestAddressRepository_CreateAddress(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Stablishing Database connection failed")
	repo := createRepository(conn)

	city := mockAndInsertCity(conn, 1)
	defer destructCities(conn, city)

	userId := rand.Int()
	mockedAddress := mockAddress(city[0].Id, userId)

	result, err := repo.CreateAddress(mockedAddress)
	defer destructAddresses(conn, []Address{*result})
	assert.NoError(t, err, "User address creation failed")
	assertAddresses(t, []Address{*mockedAddress}, []Address{*result})
}

// TestAddressRepository_UpdateAddress functionality
func TestAddressRepository_UpdateAddress(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Stablishing Database connection failed")
	repo := createRepository(conn)

	city := mockAndInsertCity(conn, 2)
	defer destructCities(conn, city)

	mockedAddresses := mockAndInsertAddresses(conn, city[0].Id, 0, 1)
	oldAddress := mockedAddresses[0]
	defer destructAddresses(conn, mockedAddresses)

	newAddress := "Edited address"

	result, err := repo.UpdateAddress(&mockedAddresses[0], city[1].Id, newAddress)
	assert.NoError(t, err, "user address edit failed")

	assert.NotEqual(t, result.Address, oldAddress.Address, "user address edit failed")
	assert.NotEqual(t, result.CityId, oldAddress.CityId, "user address edit failed")
}

// TestAddressRepository_DeleteAddress functionality
func TestAddressRepository_DeleteAddress(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Stablishing Database connection failed")
	repo := createRepository(conn)

	city := mockAndInsertCity(conn, 2)
	defer destructCities(conn, city)

	mockedAddresses := mockAndInsertAddresses(conn, city[0].Id, 0, 1)
	defer destructAddresses(conn, mockedAddresses)

	err = repo.DeleteAddress(&mockedAddresses[0])
	assert.NoError(t, err, "Deleting address failed")

	assertIsAddressDeleted(t, conn, mockedAddresses[0].Id)
}

// mockAndInsertCity in database
func mockAndInsertCity(conn *gorm.DB, count int) []City {
	var cities []City
	for i := 0; i < count; i++ {
		mockedCity := mockCity()

		result := conn.Create(mockedCity)
		if result.Error != nil {
			fmt.Println(result.Error)
			continue
		}
		cities = append(cities, *mockedCity)
	}

	return cities
}

// createRepository and return it
func createRepository(conn *gorm.DB) AddressRepositoryInterface {
	return NewAddressRepository(conn)
}

// mockAndInsertAddresses into database and return them
func mockAndInsertAddresses(db *gorm.DB, cityId, userId, count int) []Address {
	var addresses []Address

	for i := 0; i < count; i++ {
		mockedAddress := mockAddress(cityId, userId)
		result := db.Create(mockedAddress)
		if result.Error != nil {
			fmt.Println(result.Error)
			continue
		}

		addresses = append(addresses, *mockedAddress)
	}

	return addresses
}

// mockAddress and return it
func mockAddress(cityId, userId int) *Address {
	if userId == 0 {
		userId = rand.Int()
	}

	return &Address{
		UserId:  userId,
		CityId:  cityId,
		Address: "Test Address for user",
	}
}

// mockCity struct and return it
func mockCity() *City {
	return &City{
		Title: "Test City",
	}
}

// setupDbConnection and run migration
func setupDbConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(City{}, Address{})
	return db, err
}

// assertCitiesEquivelance to check whether cities are equal or not
func assertCitiesEquivelance(t *testing.T, mockedCities, fetchedCities []City) {
	for index := range mockedCities {
		assert.Equal(t, mockedCities[index].Id, fetchedCities[index].Id)
		assert.Equal(t, mockedCities[index].Title, fetchedCities[index].Title)
	}
}

// assertAddresses to check whether they are equal or not
func assertAddresses(t *testing.T, fetchedAddresses, mockedAddresses []Address) {
	for index := range mockedAddresses {
		assert.Equal(t, mockedAddresses[index].Id, fetchedAddresses[index].Id, "User addresses are not equal")
		//assert.Equal(t, mockedAddresses[index].CityId, fetchedAddresses[index].CityId, "User addresses are not equal")
		assert.Equal(t, mockedAddresses[index].UserId, fetchedAddresses[index].UserId, "User addresses are not equal")
		assert.Equal(t, mockedAddresses[index].Address, fetchedAddresses[index].Address, "User addresses are not equal")
	}
}

// destructCities which are created for testing purpose
func destructCities(conn *gorm.DB, cities []City) {
	for _, city := range cities {
		conn.Unscoped().Delete(city)
	}
}

// destructAddresses which are created for testing
func destructAddresses(conn *gorm.DB, addresses []Address) {
	for _, address := range addresses {
		conn.Unscoped().Delete(address)
	}
}

// assertIsAddressDeleted for checking delete functionalities
func assertIsAddressDeleted(t *testing.T, conn *gorm.DB, addressId int) {
	var addr Address
	fetchResult := conn.Where("id = ?", addressId).First(&addr)
	assert.Error(t, fetchResult.Error, "Deleting address failed")
	assert.ErrorIs(t, fetchResult.Error, gorm.ErrRecordNotFound, "Deleting address failed")
}
