package basket

import "context"

// BasketRepositoryInterface defines set of methods that represent basket's repository functionalities
type BasketRepositoryInterface interface {
	GetUserActiveBasket(userId int) (*Basket, error)
	GetBasketById(id int) (*Basket, error)
	GetUserBaskets(userId int) ([]Basket, error)
	GetBasketProduct(basketId int, productId int) (*BasketProduct, error)
	BasketExists(basketId int) bool
	CreateBasket(userBasket *Basket) error
	DisableBasket(userBasket *Basket) error
	AddProductToBasket(userBasket *Basket, basketProduct *BasketProduct) error
	UpdateBasketProducts(basketProduct *BasketProduct) error
	ClearBasketProducts(userBasket *Basket) error
}

// BasketServiceInterface defines set of methods that represent basket's service functionalities
type BasketServiceInterface interface {
	GetUserActiveBasket(userId int) (*Basket, error)
	GetUserBaskets(userId int) ([]Basket, error)
	CreateBasket(userId int) (*Basket, error)
	DisableUserActiveBasket(userId int) error
	AddProductsToBasket(userId, productId, amount int) (*Basket, error) //Todo should fetch product price from Product domain
	UpdateBasketProductsAmount(productId, amount int) (*Basket, error)  //Todo should fetch product price from Product domain
}

// BasketUseCaseInterface defines set of methods that represent basket's use case functionalities
type BasketUseCaseInterface interface {
	GetUserActiveBasket(ctx context.Context, token string) (*Basket, error)
	GetUserBaskets(ctx context.Context, token string) ([]Basket, error)
	CreateBasket(ctx context.Context, token string) (*Basket, error)
	DisableActiveBasket(ctx context.Context, token string) error
	AddProductsToBasket(ctx context.Context, token string, request *AddProductsToBasketRequest) (*Basket, error)         //Todo should fetch product price from Product domain
	UpdateBasketProductsAmount(ctx context.Context, token string, request *EditProductsToBasketRequest) (*Basket, error) //Todo should fetch product price from Product domain
}
