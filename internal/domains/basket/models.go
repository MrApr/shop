package basket

import (
	"shop/internal/domains/products"
	"time"
)

// Basket struct represent basket entity in system
type Basket struct {
	Id        int                `json:"id" gorm:"primaryKey"`
	UserId    int                `json:"-"`
	Status    bool               `json:"status" gorm:"index"`
	Products  []products.Product `json:"products" gorm:"many2many:basket_products;foreignKey:Id;joinForeignKey:BasketId;References:Id;joinReferences:ProductId"`
	CreatedAt *time.Time         `json:"created_at,omitempty"`
	UpdatedAt *time.Time         `json:"updated_at,omitempty"`
	DeletedAt *time.Time         `json:"deleted_at,omitempty"`
}

// BasketProduct is the struct which represents basket_products table
type BasketProduct struct {
	BasketId  int     `json:"basket_id"`
	ProductId int     `json:"product_id"`
	Amount    int     `json:"amount"`
	UnitPrice float64 `json:"unit_price"`
}

// TableName overrides default table name in gorm for them
func (*BasketProduct) TableName() string {
	return "basket_products"
}

// AddProductsToBasketRequest is the struct which represents add request
type AddProductsToBasketRequest struct {
	ProductId int `json:"product_id" validate:"required,min=1"`
	Amount    int `json:"amount" validate:"required,min=1"`
}

// EditProductsToBasketRequest is the struct which represents add request
type EditProductsToBasketRequest struct {
	ProductId int `json:"product_id" validate:"required,min=1"`
	Amount    int `json:"amount" validate:"required,min=1"`
}
