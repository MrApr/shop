package payment

// CreatePaymentGateway and return it which is the factory for that
func CreatePaymentGateway(typeId int, repo PaymentRepoContract) PaymentPGWServiceContract {
	switch typeId {
	default:
		return NewZarinpal(repo, false)
	}
}
