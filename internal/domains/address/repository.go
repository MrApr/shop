package address

import "gorm.io/gorm"

// AddressRepository is the type which has the responsibility for address data access layer
type AddressRepository struct {
	db *gorm.DB
}

// NewAddressRepository and return its abstraction
func NewAddressRepository(db *gorm.DB) AddressRepositoryInterface {
	return &AddressRepository{
		db: db,
	}
}

// GetAllCities and return them as slice of cities
func (a *AddressRepository) GetAllCities() ([]City, error) {
	var cities []City
	result := a.db.Find(&cities)
	return cities, result.Error
}

func (a *AddressRepository) GetAddressById(id int) (*Address, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AddressRepository) GetAllUserAddresses(userId int) ([]Address, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AddressRepository) CreateAddress(address *Address) (*Address, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AddressRepository) UpdateAddress(address *Address, cityId int, newAddress string) (*Address, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AddressRepository) DeleteAddress(address *Address) (bool, error) {
	//TODO implement me
	panic("implement me")
}
