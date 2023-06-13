package internal

import (
	"shop/internal/domains/payment"
	"shop/internal/domains/payment/internal/zarinpal"
)

// CreatePaymentGateway and return it which is the factory for that
func CreatePaymentGateway(typeId int, repo payment.PaymentRepoContract) payment.PaymentPGWServiceContract {
	switch typeId {
	default:
		return zarinpal.NewZarinpal(repo, false)
	}
}
