package payment

import "gorm.io/gorm"

// PaymentRepository is the type which implements PaymentRepoContract
type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepo(db *gorm.DB) PaymentRepoContract {
	return &PaymentRepository{
		db: db,
	}
}

// GetPayment and return it
func (p *PaymentRepository) GetPayment(id int) (*Payment, error) {
	payment := new(Payment)
	result := p.db.Where("id = ?", id).First(payment)
	return payment, result.Error
}

// GetUserLastPayment and return it
func (p *PaymentRepository) GetUserLastPayment(userId int, pendPayment bool) (*Payment, error) {
	//TODO implement me
	panic("implement me")
}

// CreatePayment in order to store it in db
func (p *PaymentRepository) CreatePayment(payment *Payment) error {
	//TODO implement me
	panic("implement me")
}

// UpdatePaymentRef for fetched ref number
func (p *PaymentRepository) UpdatePaymentRef(payment *Payment, refNum string) (*Payment, error) {
	//TODO implement me
	panic("implement me")
}

// UpdatePaymentTraceStatus for returned and fetched Trace Number
func (p *PaymentRepository) UpdatePaymentTraceStatus(payment *Payment, traceNum, status string) (*Payment, error) {
	//TODO implement me
	panic("implement me")
}
