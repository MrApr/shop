package address

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

// TestAddressRepository_GetAllCities functionality
func TestAddressRepository_GetAllCities(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Stablishing Database connection failed")
	repo := createRepository(conn)

	mockedCities := mockAndInsertCity(conn, 2)

	cities, err := repo.GetAllCities()
	assert.NoError(t, err, "Fetching cities from repository failed")
	assert.Equal(t, len(mockedCities), len(cities), "Counts of retrieved cities are not equal with created ones")

	assertCitiesEquivelance(t, mockedCities, cities)
}

// TestAddressRepository_GetAllUserAddresses functionality
func TestAddressRepository_GetAllUserAddresses(t *testing.T) {

}

// TestAddressRepository_GetAddressById functionality
func TestAddressRepository_GetAddressById(t *testing.T) {

}

// TestAddressRepository_UpdateAddress functionality
func TestAddressRepository_UpdateAddress(t *testing.T) {

}

// TestAddressRepository_DeleteAddress functionality
func TestAddressRepository_DeleteAddress(t *testing.T) {

}

// TestAddressRepository_CreateAddress functionality
func TestAddressRepository_CreateAddress(t *testing.T) {

}

// mockAndInsertCity in database
func mockAndInsertCity(conn *gorm.DB, count int) []City {
	var cities []City
	for i := 0; i < count; i++ {
		mockedCity := mockCity()

		result := conn.Create(mockedCity)
		if result.Error != nil {
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
	err = db.AutoMigrate(City{})
	return db, err
}

// assertCitiesEquivelance to check whether cities are equal or not
func assertCitiesEquivelance(t *testing.T, mockedCities, fetchedCities []City) {
	for index := range mockedCities {
		assert.Equal(t, mockedCities[index].Id, fetchedCities[index].Id)
		assert.Equal(t, mockedCities[index].Title, fetchedCities[index].Title)
	}
}
