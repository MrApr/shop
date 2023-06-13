package payment

import "gorm.io/gorm"

// Define payment status types
const (
	PaymentDefaultStatus string = "pending"
	PaymentSuccessStatus string = "success"
	PaymentFailureStatus string = "failed"
)

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
	var payment Payment
	db := p.db.Where("user_id = ?", userId).Order("created_at DESC")
	if pendPayment {
		db = p.db.Where("status = ?", PaymentDefaultStatus)
	}
	result := db.First(&payment)
	return &payment, result.Error
}

// CreatePayment in order to store it in db
func (p *PaymentRepository) CreatePayment(payment *Payment) error {
	result := p.db.Create(payment)
	return result.Error
}

// UpdatePaymentRef for fetched ref number
func (p *PaymentRepository) UpdatePaymentRef(payment *Payment, refNum string) (*Payment, error) {
	payment.RefNum = &refNum
	result := p.db.Save(payment)
	return payment, result.Error
}

// UpdatePaymentTraceStatus for returned and fetched Trace Number
func (p *PaymentRepository) UpdatePaymentTraceStatus(payment *Payment, traceNum, status string) (*Payment, error) {
	//TODO implement me
	panic("implement me")
}
