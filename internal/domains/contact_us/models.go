package contact_us

import "time"

// ContactUs struct represents contact us entity in system
type ContactUs struct {
	Id          int        `json:"id" gorm:"primaryKey"`
	Email       string     `json:"email"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

// CreateContactUsRequest is the struct which describes and represents, new contact us creation fields
type CreateContactUsRequest struct {
	Email       string `json:"email" validate:"required,email,max=255"`
	Title       string `json:"title" validate:"required,min=2,max=255"`
	Description string `json:"description" validate:"required,min=2,max=500"`
}
