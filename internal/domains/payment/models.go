package payment

import (
	"gorm.io/gorm"
	"time"
)

// Payment defines set of fields for payment entity in system
type Payment struct {
	Id         int            `json:"id" gorm:"primaryKey"`
	UserId     int            `json:"-"`
	BasketId   int            `json:"basket_id"`
	AddressId  int            `json:"address_id"`
	DiscountId int            `json:"discount_id"`
	GatewayId  int            `json:"gateway_id"`
	PostTypeId int            `json:"post_type_id"`
	TotalPrice float64        `json:"total_price"`
	RefNum     *string        `json:"ref_num,omitempty" gorm:"uniqueIndex"`
	TraceNum   *string        `json:"trace_num,omitempty" gorm:"uniqueIndex"`
	Status     string         `json:"status"`
	CreatedAt  *time.Time     `json:"created_at,omitempty"`
	UpdatedAt  *time.Time     `json:"updated_at,omitempty"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at,omitempty"`
}

// CreatePaymentRequest defines set of methods for payment creation
type CreatePaymentRequest struct {
	BasketId   int     `json:"basket_id" validate:"required,min=1"`
	AddressId  int     `json:"address_id" validate:"required,min=1"`
	DiscountId int     `json:"discount_id" validate:"required,min=1"`
	GatewayId  int     `json:"gateway_id" validate:"required,min=1"`
	PostTypeId int     `json:"post_type_id" validate:"required,min=1"`
	TotalPrice float64 `json:"total_price" validate:"required,min=1"`
}

// PaymentVerifyRequest defines set of fields that are required for payment verification
type PaymentVerifyRequest struct {
	PaymentId int    `query:"payment_id" validate:"required,min=1"`
	Authority string `query:"authority" validate:"required,min=36,max=36"`
}

// GetUserPaymentsRequest is the struct which defines necessary fields for desired operation
type GetUserPaymentsRequest struct {
	From *int `json:"from,omitempty" validate:"required,min=1"`
	To   *int `json:"to,omitempty" validate:"required,min=1,gtefield=From"`
}

// RequestPaymentResponse defines set of fields in which is returned by payment gateway before starting payment operation and redirecting to bank
type RequestPaymentResponse struct {
	Url             string
	Key             string
	OperationStatus int
}
