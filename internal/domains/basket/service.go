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

// AddProductsToBasket for user which has already an active basket
func (b *BasketService) AddProductsToBasket(userId, productId, amount int) (*Basket, error) {
	err := b.checkAndUpdateProductStack(productId, amount)
	if err != nil {
		return nil, err
	}

	activeBasket, err := b.getOrCreateUserActiveBasket(userId)
	if err != nil {
		return nil, advancedError.New(err, "Cannot create or get user active basket")
	}

	price := b.getProductPrice(productId)
	basketProduct := b.createBasketProduct(productId, amount, price)

	result := b.basketRepo.AddProductToBasket(activeBasket, basketProduct)
	if result != nil {
		return nil, err
	}

	basket, err := b.basketRepo.GetBasketById(activeBasket.Id)
	if err != nil {
		return nil, err
	}

	return basket, nil
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

// getProductPrice and return it
func (b *BasketService) getProductPrice(productId int) float64 {
	product, _ := b.productService.GetProduct(productId)
	return product.Price
}

// checkAndUpdateProductStack
func (b *BasketService) checkAndUpdateProductStack(productId, amount int) error {
	product, err := b.productService.GetProduct(productId)
	if err != nil {
		return err
	}

	if err = b.checkProductStack(product.Amount, amount); err != nil {
		return err
	}

	product, err = b.productService.UpdateProductInventory(productId, amount)
	if err != nil {
		return err
	}

	return nil
}

// checkProductStack and see whether it's enough or not
func (b *BasketService) checkProductStack(amount, requestedAmount int) error {
	switch {
	case amount == 0:
		return ProductIsFinished
	case amount < requestedAmount:
		return InsufficientProductAmount
	default:
		return nil
	}
}

// getOrCreateUserActiveBasket is a wrapper for simplifying code
func (b *BasketService) getOrCreateUserActiveBasket(userId int) (*Basket, error) {
	basket, err := b.GetUserActiveBasket(userId)
	if err == nil && basket.Id != 0 {
		return basket, nil
	}

	return b.CreateBasket(userId)
}

// createBasketProduct and return it after initializing it
func (b *BasketService) createBasketProduct(productId, amount int, price float64) *BasketProduct {
	return &BasketProduct{
		ProductId: productId,
		Amount:    amount,
		UnitPrice: price,
	}
}
