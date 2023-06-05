package basket

import (
	"context"
	"shop/pkg/userHandler"
)

// BasketUseCase is the struct which implements BasketUseCaseInterface
type BasketUseCase struct {
	sv        BasketServiceInterface
	decoderFn func(ctx context.Context, token string) (int, error)
}

// NewUseCase creates and returns basket use case
func NewUseCase(sv BasketServiceInterface, decoderFn func(ctx context.Context, token string) (int, error)) BasketUseCaseInterface {
	if decoderFn == nil {
		decoderFn = userHandler.ExtractUserIdFromToken
	}

	return &BasketUseCase{
		sv:        sv,
		decoderFn: decoderFn,
	}
}

// GetUserActiveBasket and return it
func (b *BasketUseCase) GetUserActiveBasket(ctx context.Context, token string) (*Basket, error) {
	userId, err := b.decoderFn(ctx, token)
	if err != nil {
		return nil, err
	}

	return b.sv.GetUserActiveBasket(userId)
}

// GetUserBaskets and return them
func (b *BasketUseCase) GetUserBaskets(ctx context.Context, token string) ([]Basket, error) {
	userId, err := b.decoderFn(ctx, token)
	if err != nil {
		return nil, err
	}

	return b.sv.GetUserBaskets(userId)
}

// CreateBasket and store it in db then return it
func (b *BasketUseCase) CreateBasket(ctx context.Context, token string) (*Basket, error) {
	userId, err := b.decoderFn(ctx, token)
	if err != nil {
		return nil, err
	}

	return b.sv.CreateBasket(userId)
}

// DisableActiveBasket and which already exists in db
func (b *BasketUseCase) DisableActiveBasket(ctx context.Context, token string) error {
	//TODO implement me
	panic("implement me")
}

// AddProductsToBasket and return general basket
func (b *BasketUseCase) AddProductsToBasket(ctx context.Context, token string, request *AddProductsToBasketRequest) (*Basket, error) {
	userId, err := b.decoderFn(ctx, token)
	if err != nil {
		return nil, err
	}

	return b.sv.AddProductsToBasket(userId, request.ProductId, request.Amount)
}

// UpdateBasketProductsAmount and return them
func (b *BasketUseCase) UpdateBasketProductsAmount(ctx context.Context, token string, request *EditProductsToBasketRequest) (*Basket, error) {
	userId, err := b.decoderFn(ctx, token)
	if err != nil {
		return nil, err
	}

	return b.sv.UpdateBasketProductsAmount(userId, request.ProductId, request.Amount)
}
