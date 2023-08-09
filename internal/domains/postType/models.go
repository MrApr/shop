package postType

import (
	"gorm.io/gorm"
	"time"
)

// PostType struct represents post type entity in system
type PostType struct {
	Id              int            `json:"id" gorm:"primaryKey"`
	Title           string         `json:"title"`
	Price           float64        `json:"price"`
	DeliverableTime uint64         `json:"deliverable_time"` //In minutes
	CreatedAt       *time.Time     `json:"created_at,omitempty"`
	UpdatedAt       *time.Time     `json:"updated_at,omitempty"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at,omitempty"`
}
