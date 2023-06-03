package basket

import (
	"shop/internal/domains/products"
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

func (b *BasketService) GetUserBaskets(userId int) ([]Basket, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BasketService) CreateBasket(userId int) (*Basket, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BasketService) DisableActiveBasket() error {
	//TODO implement me
	panic("implement me")
}

func (b *BasketService) AddProductsToBasket(productId, amount int) (*Basket, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BasketService) UpdateBasketProductsAmount(productId, amount int) (*Basket, error) {
	//TODO implement me
	panic("implement me")
}
