package payment

// PaymentStorageService struct which implements PaymentStorageServiceContract
type PaymentStorageService struct {
	repo PaymentRepoContract
}

// NewPaymentStorageService for storage common issues
func NewPaymentStorageService(repo PaymentRepoContract) PaymentStorageServiceContract {
	return &PaymentStorageService{
		repo: repo,
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
	//TODO implement me
	panic("implement me")
}
