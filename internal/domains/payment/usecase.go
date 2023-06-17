package payment

import (
	"context"
	"gorm.io/gorm"
	"shop/pkg/advancedError"
)

// PaymentUseCase is the usecase which implements payment use case contract and satisfies all of it's methods
type PaymentUseCase struct {
	paymentSvStorage PaymentStorageServiceContract
	tokenDecoder     func(ctx context.Context, token string) (int, error)
	db               *gorm.DB
}

// NewPaymentUseCase creates and returns payment usecase
func NewPaymentUseCase(svStorage PaymentStorageServiceContract, db *gorm.DB, decoder func(ctx context.Context, token string) (int, error)) PaymentUseCaseContract {
	return &PaymentUseCase{
		paymentSvStorage: svStorage,
		db:               db,
		tokenDecoder:     decoder,
	}
}

// GetUserPayments and return them
func (p *PaymentUseCase) GetUserPayments(ctx context.Context, token string, request *GetUserPaymentsRequest) ([]Payment, error) {
	panic("implement me")
}

// CreatePayment and return it
func (p *PaymentUseCase) CreatePayment(ctx context.Context, token string, request *CreatePaymentRequest) (*Payment, error) {
	userId, err := p.tokenDecoder(ctx, token)
	if err != nil {
		return nil, err
	}

	return p.paymentSvStorage.CreatePayment(userId, request.BasketId, request.AddressId, request.DiscountId, request.GatewayId, request.PostTypeId, request.TotalPrice)
}

// Pay operates payment
func (p *PaymentUseCase) Pay(ctx context.Context, token string, paymentId int) (*RequestPaymentResponse, error) {
	doesntHavePermission := p.userCanPay(ctx, paymentId, token)
	if doesntHavePermission != nil {
		return nil, doesntHavePermission
	}

	paymentSvPGW, err := p.getPaymentSvPGWForUser(paymentId)
	if err != nil {
		return nil, err
	}
	return paymentSvPGW.Pay(paymentId)
}

// Verify done payment by user
func (p *PaymentUseCase) Verify(ctx context.Context, token string, request *PaymentVerifyRequest) (*Payment, error) {
	doesntHavePermission := p.userCanPay(ctx, request.PaymentId, token)
	if doesntHavePermission != nil {
		return nil, doesntHavePermission
	}

	tx := p.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	paymentSvPGW, err := p.getPaymentSvPGWForUser(request.PaymentId)
	if err != nil {
		return nil, err
	}

	result, err := paymentSvPGW.Verify(request.PaymentId, request.Authority)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, advancedError.New(err, SomethingWentWrong.Error())
	}

	return result, nil
}

// userCanPay check whether user can do required operation or not
func (p *PaymentUseCase) userCanPay(ctx context.Context, paymentId int, token string) error {
	userId, err := p.tokenDecoder(ctx, token)
	if err != nil {
		return err
	}

	payment, err := p.paymentSvStorage.GetPayment(paymentId)
	if err != nil {
		return err
	}

	if payment.UserId != userId {
		return Unauthorized
	}

	return nil
}

// getPaymentSvPGWForUser for and return it to do payment operation
func (p *PaymentUseCase) getPaymentSvPGWForUser(paymentId int) (PaymentPGWServiceContract, error) {
	payment, err := p.paymentSvStorage.GetPayment(paymentId)
	if err != nil {
		return nil, err
	}
	return CreatePaymentGateway(payment.GatewayId, NewPaymentRepo(p.db)), nil
}
