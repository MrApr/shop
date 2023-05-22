package address

import (
	"context"
	"shop/pkg/advancedError"
	"shop/pkg/userHandler"
)

// AddressUseCase is the type which implements user address use-case interface methods
type AddressUseCase struct {
	sv        AddressServiceInterface
	decoderFn func(ctx context.Context, token string) (int, error)
}

// NewUseCase and return it
func NewUseCase(sv AddressServiceInterface, decoderFn func(ctx context.Context, token string) (int, error)) AddressUseCaseInterface {
	if decoderFn == nil {
		decoderFn = userHandler.ExtractUserIdFromToken
	}

	return &AddressUseCase{
		sv:        sv,
		decoderFn: decoderFn,
	}
}

// GetAllCities and return them
func (a *AddressUseCase) GetAllCities(ctx context.Context) ([]City, error) {
	return a.sv.GetAllCities()
}

// GetAllUserAddresses by their user id and return it
func (a *AddressUseCase) GetAllUserAddresses(ctx context.Context, token string) ([]Address, error) {
	userId, err := a.decoderFn(ctx, token)
	if err != nil {
		return nil, advancedError.New(err, "Decoding token failed")
	}

	return a.sv.GetAllUserAddresses(userId)
}

// CreateAddress for user and store it in db
func (a *AddressUseCase) CreateAddress(ctx context.Context, token string, request *CreateAddressRequest) (*Address, error) {
	userId, err := a.decoderFn(ctx, token)
	if err != nil {
		return nil, advancedError.New(err, "Decoding token failed")
	}

	return a.sv.CreateAddress(userId, request.CityId, request.Address)
}

// UpdateAddress which exists already in database for user
func (a *AddressUseCase) UpdateAddress(ctx context.Context, token string, request *UpdateAddressRequest) (*Address, error) {
	userId, err := a.decoderFn(ctx, token)
	if err != nil {
		return nil, advancedError.New(err, "Decoding token failed")
	}

	return a.sv.UpdateAddress(userId, request.AddressId, request.CityId, request.Address)
}

// DeleteAddress which already exists for user in database
func (a *AddressUseCase) DeleteAddress(ctx context.Context, token string, request *DeleteAddressRequest) error {
	userId, err := a.decoderFn(ctx, token)
	if err != nil {
		return advancedError.New(err, "Decoding token failed")
	}

	return a.sv.DeleteAddress(userId, request.AddressId)
}
