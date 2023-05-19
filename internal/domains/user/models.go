package user

import "time"

// User struct defines user entity which is used in database
type User struct {
	Id          int        `json:"-" gorm:"primaryKey"`
	UUID        string     `json:"uuid" gorm:"uniqueIndex"`
	PhoneNumber string     `json:"phone_number" gorm:"uniqueIndex"`
	Name        *string    `json:"name,omitempty"`
	Password    string     `json:"password"`
	LastLoginAt *time.Time `json:"last_login_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// UserRegisterRequest defines a struct for user creation operation
type UserRegisterRequest struct {
	PhoneNumber     string  `json:"phone_number" validate:"required,len=11"`
	Password        string  `json:"password" validate:"required,min:8,max:255"`
	PasswordConfirm string  `json:"password_confirm" validate:"required,min=8,max=255,eqfield=Password"`
	Name            *string `json:"name,omitempty" validate:"omitempty,min=3,max=255"`
}

// UserLoginRequest defines a fields for login operation and appropriate validation rules for that too
type UserLoginRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required,len=11"`
	Password    string `json:"password" validate:"required,min:8,max:255"`
}

// UpdateUserRequest defines a struct which contains necessary fields and data for updating user operation
type UpdateUserRequest struct {
	Name            *string `json:"name,omitempty" validate:"omitempty,min:3,max=255"`
	Password        *string `json:"password,omitempty" validate:"omitempty,min:8,max:255"`
	PasswordConfirm *string `json:"password_confirm,omitempty" validate:"omitempty,min=8,max=255,eqfield=Password"`
}

// AuthResponse defines struct for authentication operation that contains necessary fields for auth operation
type AuthResponse struct {
	Tokens map[string]string `json:"authorization"`
	User   *User             `json:"user"`
}
