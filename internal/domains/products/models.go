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
