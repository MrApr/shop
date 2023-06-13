package postType

import "time"

// PostType struct represents post type entity in system
type PostType struct {
	Id              int        `json:"id" gorm:"primaryKey"`
	Title           string     `json:"title"`
	Price           float64    `json:"price"`
	DeliverableTime time.Time  `json:"deliverable_time"`
	CreatedAt       *time.Time `json:"created_at,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
}
