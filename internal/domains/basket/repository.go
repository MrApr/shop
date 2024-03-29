package basket

import (
	"errors"
	"gorm.io/gorm"
)

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
	result := b.db.Preload("Products").Where("user_id = ?", userId).Where("status = ?", true).First(activeBasket)
	return activeBasket, result.Error
}

// GetBasketById and return it
func (b *SqlBasket) GetBasketById(id int) (*Basket, error) {
	basket := new(Basket)
	result := b.db.Preload("Products").Where("id = ?", id).First(basket)
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
	basketProduct := new(BasketProduct)
	result := b.db.Where("basket_id = ?", basketId).Where("product_id = ?", productId).First(basketProduct)
	return basketProduct, result.Error
}

// BasketExists or not with given Id
func (b *SqlBasket) BasketExists(basketId int) bool {
	basket := new(Basket)

	result := b.db.Where("id = ?", basketId).First(basket)

	return result.Error == nil && !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

// CreateBasket and insert it in database
func (b *SqlBasket) CreateBasket(userBasket *Basket) error {
	result := b.db.Create(userBasket)
	return result.Error
}

// DisableBasket by making it's Status false
func (b *SqlBasket) DisableBasket(userBasket *Basket) error {
	return b.db.Model(Basket{}).Where("id = ?", userBasket.Id).Update("status", false).Error
}

// AddProductToBasket which is already doesn't exist
func (b *SqlBasket) AddProductToBasket(userBasket *Basket, basketProduct *BasketProduct) error {
	basketProduct.BasketId = userBasket.Id
	return b.db.Create(basketProduct).Error
}

// UpdateBasketProducts which is already exists in database
func (b *SqlBasket) UpdateBasketProducts(basketProduct *BasketProduct) error {
	return b.db.Model(basketProduct).Where("product_id = ?", basketProduct.ProductId).Where("basket_id = ?", basketProduct.BasketId).Updates(BasketProduct{
		Amount:    basketProduct.Amount,
		UnitPrice: basketProduct.UnitPrice,
	}).Error
}

// ClearBasketProducts which already exists in db for a specific basket
func (b *SqlBasket) ClearBasketProducts(userBasket *Basket) error {
	return b.db.Where("basket_id = ?", userBasket.Id).Delete(&BasketProduct{}).Error
}
