package image

import "time"

// Image struct Defines image model for image domain
type Image struct {
	Id            int       `json:"id" gorm:"primaryKey"`
	ImageableID   int       `json:"imageable_id" gorm:"column:imageable_id"`
	ImageableType string    `json:"imageable_type" gorm:"column:imageable_type"`
	Path          string    `json:"path"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// HasImageable interface defines set of methods, that every model who has images should satisfy these.
type HasImageable interface {
	GetImageableId() int
	GetImageableType() string
}
