package payment

import (
	"shop/pkg/advancedError"
	"shop/pkg/paymentHandler"
)

// PaymentStorageService struct which implements PaymentStorageServiceContract
type PaymentStorageService struct {
	repo              PaymentRepoContract
	basketValidator   func(int) error
	addressValidator  func(int) error
	discountValidator func(int) error
	postTypeValidator func(int) error
}

// NewPaymentStorageService for storage common issues
func NewPaymentStorageService(repo PaymentRepoContract, basketValidator, addressValidator, discountValidator, postTypeValidator func(id int) error) PaymentStorageServiceContract {
	if basketValidator == nil {
		basketValidator = paymentHandler.BasketIsValid
	}

	if addressValidator == nil {
		addressValidator = paymentHandler.AddressIsValid
	}

	if discountValidator == nil {
		discountValidator = paymentHandler.IsDiscountValid
	}

	if postTypeValidator == nil {
		postTypeValidator = paymentHandler.IsPostTypeValid
	}

	return &PaymentStorageService{
		repo:              repo,
		basketValidator:   basketValidator,
		addressValidator:  addressValidator,
		discountValidator: discountValidator,
		postTypeValidator: postTypeValidator,
	}
}

// GetPayment and return it based on given Id
func (p *PaymentStorageService) GetPayment(id int) (*Payment, error) {
	payment, err := p.repo.GetPayment(id)
	if err != nil {
		return nil, PaymentNotFound
	}
	return payment, nil
}

// CreatePayment and return it based on given Data
func (p *PaymentStorageService) CreatePayment(userId, BasketId, addressId, discountId, gatewayId, postTypeId int, totalPrice float64) (*Payment, error) {
	payment := Payment{
		UserId:     userId,
		BasketId:   BasketId,
		AddressId:  addressId,
		DiscountId: discountId,
		GatewayId:  gatewayId,
		PostTypeId: postTypeId,
		TotalPrice: totalPrice,
		Status:     PaymentDefaultStatus,
	}

	err := p.paymentCanCreated(&payment)
	if err != nil {
		return nil, err
	}

	err = p.repo.CreatePayment(&payment)
	if err != nil {
		return nil, advancedError.New(err, "cannot create a new payment")
	}
	return &payment, nil
}

// paymentCanCreated checks whether payment is allowed to insert in db or not!
func (p *PaymentStorageService) paymentCanCreated(payment *Payment) error {
	if err := p.basketValidator(payment.BasketId); err != nil {
		return err
	}

	if err := p.addressValidator(payment.AddressId); err != nil {
		return err
	}

	if err := p.discountValidator(payment.DiscountId); err != nil {
		return err
	}

	if err := p.postTypeValidator(payment.PostTypeId); err != nil {
		return err
	}

	return nil
}
