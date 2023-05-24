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
