package payment

import "context"

// PaymentRepoContract defines set of methods that every type, which wants to play payment repository role, should obey
type PaymentRepoContract interface {
	GetPayment(id int) (*Payment, error)
	GetUserLastPayment(userId int, pendPayment bool) (*Payment, error)
	CreatePayment(payment *Payment) error
	UpdatePaymentTrace(payment *Payment, traceNum string) (*Payment, error)
	UpdatePaymentRefStatus(payment *Payment, refNum, status string) (*Payment, error)
}

// PaymentStorageServiceContract defines set of methods that every type, which wants to play payment service role, should obey
type PaymentStorageServiceContract interface {
	GetPayment(id int) (*Payment, error)
	CreatePayment(userId, BasketId, addressId, discountId, gatewayId, postTypeId int, totalPrice float64) (*Payment, error)
}

// PaymentPGWServiceContract defines set of methods that every implementation of payment gateway should obey
type PaymentPGWServiceContract interface {
	Pay(paymentId int) (*RequestPaymentResponse, error)
	Verify(paymentId int, Authority string) (*Payment, error)
}

// PaymentUseCaseContract defines set of methods that every type, which wants to play payment use case role, should obey
type PaymentUseCaseContract interface {
	CreatePayment(ctx context.Context, token string, request *CreatePaymentRequest) (*Payment, error)
	Pay(ctx context.Context, token string, paymentId int) (*RequestPaymentResponse, error)
	Verify(ctx context.Context, token string, request *PaymentVerifyRequest) (*Payment, error)
}
