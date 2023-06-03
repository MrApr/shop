package basket

import (
	"errors"
	"gorm.io/gorm"
	"shop/internal/domains/products"
	"shop/pkg/advancedError"
)

// BasketService is the type which implements BasketServiceInterface
type BasketService struct {
	basketRepo     BasketRepositoryInterface
	productService products.ProductServiceInterface
}

// NewBasketService and return it
func NewBasketService(basketRepo BasketRepositoryInterface, productSv products.ProductServiceInterface) BasketServiceInterface {
	return &BasketService{
		basketRepo:     basketRepo,
		productService: productSv,
	}
}

// GetUserActiveBasket and return it based on its userId
func (b *BasketService) GetUserActiveBasket(userId int) (*Basket, error) {
	return b.basketRepo.GetUserActiveBasket(userId)
}

// GetUserBaskets and return them based on their Id
func (b *BasketService) GetUserBaskets(userId int) ([]Basket, error) {
	return b.basketRepo.GetUserBaskets(userId)
}

// CreateBasket for user and return it
func (b *BasketService) CreateBasket(userId int) (*Basket, error) {
	err := b.checkAndDisableActiveBasket(userId)
	if err != nil {
		return nil, advancedError.New(err, "Disabling active basket for user, in new basket creation failed")
	}
	newBasket := &Basket{
		UserId: userId,
		Status: true,
	}

	err = b.basketRepo.CreateBasket(newBasket)
	if err != nil {
		return nil, advancedError.New(err, "Creating new basket for user failed")
	}

	return newBasket, err
}

// DisableUserActiveBasket which is already inserted and present in database
func (b *BasketService) DisableUserActiveBasket(userId int) error {
	return b.checkAndDisableActiveBasket(userId)
}

func (b *BasketService) AddProductsToBasket(productId, amount int) (*Basket, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BasketService) UpdateBasketProductsAmount(productId, amount int) (*Basket, error) {
	//TODO implement me
	panic("implement me")
}

// checkAndDisableActiveBasket for user in order to create a new one
func (b *BasketService) checkAndDisableActiveBasket(userId int) error {
	activeBasket, err := b.basketRepo.GetUserActiveBasket(userId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	err = b.basketRepo.DisableBasket(activeBasket)
	return err
}
