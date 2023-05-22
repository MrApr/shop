package address

import "time"

// Address is a struct which defines address entity fields and properties in system
type Address struct {
	Id        int        `json:"id" gorm:"primaryKey"`
	UserId    int        `json:"-"`
	CityId    int        `json:"city_id"`
	City      *City      `json:"city,omitempty" gorm:"foreignKey:CityId;references:Id"`
	Address   string     `json:"address"`
	CreatedAt *time.Time `json:"created_at"`
	UpdateAt  *time.Time `json:"update_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// TableName overrides table name for gorm model
func (Address) TableName() string {
	return "user_addresses"
}

// City is a struct which contains city's entity fields
type City struct {
	Id    int    `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
}

// CreateAddressRequest defines set of parameters for user address creation
type CreateAddressRequest struct {
	CityId  int    `json:"city_id" validate:"required,numeric,min=1"`
	Address string `json:"address" validate:"required,min=5,max=500"`
}

// UpdateAddressRequest defines set of parameters for user address modification
type UpdateAddressRequest struct {
	AddressId int    `json:"address_id" validate:"required,numeric,min=1"`
	CityId    int    `json:"city_id" validate:"required,numeric,min=1"`
	Address   string `json:"address" validate:"required,min=5,max=500"`
}

// DeleteAddressRequest defines set of methods for user address deletation
type DeleteAddressRequest struct {
	AddressId int `json:"address_id" validate:"required,numeric,min=1"`
}
