package address

// AddressService is a type which implements all methods of address service interface and plays role as address service
type AddressService struct {
	repo AddressRepositoryInterface
}

// NewAddressService and return it
func NewAddressService(repo AddressRepositoryInterface) AddressServiceInterface {
	return &AddressService{
		repo: repo,
	}
}

// GetAllCities and return them
func (a *AddressService) GetAllCities() ([]City, error) {
	cities, err := a.repo.GetAllCities()
	if len(cities) == 0 {
		return nil, NoCitiesFound
	}
	return cities, err
}

// GetAllUserAddresses bases on user id and return them
func (a *AddressService) GetAllUserAddresses(userId int) ([]Address, error) {
	addresses, err := a.repo.GetAllUserAddresses(userId)

	if len(addresses) == 0 {
		return nil, NoAddressesFound
	}

	return addresses, err
}

func (a *AddressService) CreateAddress(userId, cityId int, address string) (*Address, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AddressService) UpdateAddress(requestedId, addressId, cityId int, newAddress string) (*Address, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AddressService) DeleteAddress(requestedId, addressId int) error {
	//TODO implement me
	panic("implement me")
}
