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

func (a *AddressUseCase) CreateAddress(ctx context.Context, token string, request *CreateAddressRequest) (*Address, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AddressUseCase) UpdateAddress(ctx context.Context, token string, request *UpdateAddressRequest) (*Address, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AddressUseCase) DeleteAddress(ctx context.Context, token string, request *DeleteAddressRequest) error {
	//TODO implement me
	panic("implement me")
}
