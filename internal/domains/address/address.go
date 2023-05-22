package address

import "context"

// AddressRepositoryInterface defines set of methods that each type who wants to play role as address repository should obey
type AddressRepositoryInterface interface {
	GetAllCities() ([]City, error)
	GetAddressById(id int) (*Address, error)
	GetAllUserAddresses(userId int) ([]Address, error)
	CreateAddress(address *Address) (*Address, error)
	UpdateAddress(address *Address, cityId int, newAddress string) (*Address, error)
	DeleteAddress(address *Address) error
}

// AddressServiceInterface defines set of methods that each type who wants to play role as address service should obey
type AddressServiceInterface interface {
	GetAllCities() ([]City, error)
	GetAllUserAddresses(userId int) ([]Address, error)
	CreateAddress(userId, cityId int, address string) (*Address, error)
	UpdateAddress(requestedId, addressId, cityId int, newAddress string) (*Address, error) //Todo add Authorization on address edit and delete operation check user is the owner or not
	DeleteAddress(requestedId, addressId int) error
}

// AddressUseCaseInterface defines set of methods that each type who wants to play role as address use-case should obey
type AddressUseCaseInterface interface {
	GetAllCities(ctx context.Context) ([]City, error)
	GetAllUserAddresses(ctx context.Context, userId int) ([]Address, error)
	CreateAddress(ctx context.Context, token string, request *CreateAddressRequest) (*Address, error)
	UpdateAddress(ctx context.Context, token string, request *UpdateAddressRequest) (*Address, error)
	DeleteAddress(ctx context.Context, token string, request *DeleteAddressRequest) error
}

//Todo add global package for handing Authorization in all domains
