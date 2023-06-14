package comment

import "time"

// Comment defines set of fields which represents comment entity
type Comment struct {
	Id          int        `json:"id" gorm:"primaryKey"`
	ProductId   int        `json:"product_id"`
	UserId      int        `json:"-"`
	Description string     `json:"description"`
	Status      bool       `json:"status" gorm:"index"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

// CreateCommentRequest is the request which contains necessary fields for comment creation
type CreateCommentRequest struct {
	ProductId   int    `json:"product_id" validate:"required,min=1"`
	Description string `json:"description" validate:"required,min=1,max=255"`
}
