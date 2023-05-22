package address

import "shop/pkg/authorization"

// AddressService is a type which implements all methods of address service interface and plays role as address service
type AddressService struct {
	repo AddressRepositoryInterface
}

// userAddrAuthorizerField is the field which authorization should get done and checked with that
const userAddrAuthorizerField string = "UserId"

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

// CreateAddress for user and insert it in database
func (a *AddressService) CreateAddress(userId, cityId int, address string) (*Address, error) {
	cityExists := a.repo.CityExists(cityId)
	if !cityExists {
		return nil, CityNotFound
	}

	tmpAddr := &Address{
		UserId:  userId,
		CityId:  cityId,
		Address: address,
	}

	return a.repo.CreateAddress(tmpAddr)
}

// UpdateAddress which exists in database and user is the owner of it
func (a *AddressService) UpdateAddress(requestedId, addressId, cityId int, newAddress string) (*Address, error) {
	address, err := a.repo.GetAddressById(addressId)
	if err != nil {
		return nil, AddressNotFound
	}

	if err = authorization.SimpleFieldAuthorization(*address, userAddrAuthorizerField, requestedId, YouAreNotAllowed); err != nil {
		return nil, err
	}

	cityExists := a.repo.CityExists(cityId)
	if !cityExists {
		return nil, CityNotFound
	}

	return a.repo.UpdateAddress(address, cityId, newAddress)
}

// DeleteAddress that already exists in database for user
func (a *AddressService) DeleteAddress(requestedId, addressId int) error {
	address, err := a.repo.GetAddressById(addressId)
	if err != nil {
		return AddressNotFound
	}

	if err = authorization.SimpleFieldAuthorization(*address, userAddrAuthorizerField, requestedId, YouAreNotAllowed); err != nil {
		return err
	}

	return a.repo.DeleteAddress(address)
}
