package products

import "time"

// Type struct defines type entity
type Type struct {
	Id        int        `json:"id" gorm:"primaryKey"`
	Title     string     `json:"title"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// GetAllTypesRequest is the struct which represents get all types request
type GetAllTypesRequest struct {
	Name   *string `json:"name,omitempty" validate:"omitempty,min=2,max=255"`
	Limit  int     `json:"limit" validate:"required,gte=1"`
	Offset *int    `json:"offset,omitempty" validate:"omitempty,min=0"`
}

// Category defines category entity in system
type Category struct {
	Id             int        `json:"id" gorm:"primaryKey"`
	TypeId         int        `json:"type_id"`
	Type           Type       `json:"type" gorm:"foreignKey:TypeId;references:Id"`
	ParentCatId    *int       `json:"parent_cat_id,omitempty"`
	ParentCategory *Category  `json:"parent_category,omitempty" gorm:"foreignKey:ParentCatId;reference:Id"`
	Title          string     `json:"title"`
	Indent         int        `json:"indent"`
	Order          int        `json:"order"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
}

// GetAllCategoriesRequest is the struct which represents get all categories request
type GetAllCategoriesRequest struct {
	Title       *string `json:"title,omitempty" validate:"omitempty,min=3,max=255"`
	TypeId      *int    `json:"typeId,omitempty" validate:"omitempty,min=1"`
	ParentCatId *int    `json:"parent_cat_id,omitempty" validate:"omitempty,min=1"`
	Limit       *int    `json:"limit" validate:"required,gte=1"`
	Offset      int     `json:"offset" validate:"omitempty,min=0"`
}

// Product represent product entity in system
type Product struct {
	Id          int        `json:"id" gorm:"primaryKey"`
	Categories  []Category `json:"categories" gorm:"many2many:product_categories;foreignKey:Id;joinForeignKey:ProductId;References:Id;joinReferences:CategoryId"`
	Likes       []Likes    `json:"likes,omitempty" gorm:"many2many:likes;foreignKey:Id;joinForeignKey:ProductId;References:Id;joinReferences:UserId"`
	Dislikes    []DisLikes `json:"dislikes,omitempty" gorm:"many2many:dislikes;foreignKey:Id;joinForeignKey:ProductId;References:Id;joinReferences:UserId"`
	Title       string     `json:"title" gorm:"index"`
	Code        int        `json:"code" gorm:"uniqueIndex"`
	Amount      int        `json:"amount"`
	Price       float64    `json:"price" gorm:"index"`
	Weight      *int       `json:"weight"`
	Description *string    `json:"description"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

// GetAllProductsRequest and return them
type GetAllProductsRequest struct {
	CategoryIds []int    `json:"category_ids,omitempty" validate:"omitempty"`
	Title       *string  `json:"title,omitempty" validate:"omitempty;min=3;max=255"`
	Description *string  `json:"description,omitempty"  validate:"omitempty;min=10;max=500"`
	MinWeight   *int     `json:"min_weight,omitempty" validate:"omitempty:min=1"`
	MaxWeight   *int     `json:"high_weigh,omitempty" validate:"omitempty;min=1;gtefield:MinWeight"`
	MinPrice    *float64 `json:"min_price,omitempty" validate:"omitempty:min=1"`
	MaxPrice    *float64 `json:"max_price,omitempty" validate:"omitempty:min=1;gtefield:MinPrice"`
	Limit       *int     `json:"limit,omitempty" validate:"omitempty:min=1"`
	Offset      int      `json:"offset" validate:"omitempty:min=1"`
}

// Likes determines Likes entity in system
type Likes struct {
	ProductId int `json:"product_id" gorm:"primaryKey"`
	UserId    int `json:"user_id" gorm:"primaryKey"`
}

// DisLikes determines DisLikes entity in system
type DisLikes struct {
	ProductId int `json:"product_id" gorm:"primaryKey"`
	UserId    int `json:"user_id" gorm:"primaryKey"`
}
