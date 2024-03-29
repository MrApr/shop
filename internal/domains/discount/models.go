package discount

import (
	"gorm.io/gorm"
	"time"
)

// DiscountCode is the struct which represents discount entity in system
type DiscountCode struct {
	Id              int            `json:"id" gorm:"primaryKey"`
	Title           string         `json:"title"`
	Code            string         `json:"code" gorm:"uniqueIndex"`
	DiscountPercent int            `json:"discount_percent"`
	Status          bool           `json:"status"`
	CreatedAt       *time.Time     `json:"created_at,omitempty"`
	UpdatedAt       *time.Time     `json:"updated_at,omitempty"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at,omitempty"`
}

// GetAllDiscountCodesRequest for incoming http requests that want to fetch all discount codes
type GetAllDiscountCodesRequest struct {
	From  int `json:"from" validate:"required,min=0"`
	Limit int `json:"Limit" validate:"required,min=1"`
}
