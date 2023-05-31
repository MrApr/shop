package basket

import "gorm.io/gorm"

// SqlBasket implements BasketRepositoryInterface
type SqlBasket struct {
	db *gorm.DB
}

// NewBasketRepository and return it
func NewBasketRepository(db *gorm.DB) BasketRepositoryInterface {
	return &SqlBasket{
		db: db,
	}
}

// GetUserActiveBasket which has status true
func (b *SqlBasket) GetUserActiveBasket(userId int) (*Basket, error) {
	activeBasket := new(Basket)
	result := b.db.Where("user_id = ?", userId).Where("status = ?", true).First(activeBasket)
	return activeBasket, result.Error
}

// GetBasketById and return it
func (b *SqlBasket) GetBasketById(id int) (*Basket, error) {
	basket := new(Basket)
	result := b.db.Where("id = ?", id).First(basket)
	return basket, result.Error
}

// GetUserBaskets that exists in system
func (b *SqlBasket) GetUserBaskets(userId int) ([]Basket, error) {
	var userBaskets []Basket
	result := b.db.Where("user_id = ?", userId).Find(&userBaskets)
	return userBaskets, result.Error
}

// GetBasketProduct and return it for
func (b *SqlBasket) GetBasketProduct(basketId int, productId int) (*BasketProduct, error) {
	//TODO implement me
	panic("implement me")
}

func (b *SqlBasket) BasketExists(basketId int) bool {
	//TODO implement me
	panic("implement me")
}

func (b *SqlBasket) CreateBasket(userBasket *Basket) error {
	//TODO implement me
	panic("implement me")
}

func (b *SqlBasket) DisableBasket(userBasket *Basket) error {
	//TODO implement me
	panic("implement me")
}

func (b *SqlBasket) AddProductToBasket(userBasket *Basket, basketProduct *BasketProduct) error {
	//TODO implement me
	panic("implement me")
}

func (b *SqlBasket) UpdateBasketProducts(userBasket *Basket, basketProduct *BasketProduct) error {
	//TODO implement me
	panic("implement me")
}

func (b *SqlBasket) ClearBasketProducts(userBasket *Basket) error {
	//TODO implement me
	panic("implement me")
}
