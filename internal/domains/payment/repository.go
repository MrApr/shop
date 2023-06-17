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

// GetUserPayments and return them based on user-ID
func (p *PaymentRepository) GetUserPayments(userId, from, limit int) []Payment {
	var payments []Payment
	p.db.Where("user_id = ?", userId).Offset(from).Limit(limit).Find(&payments)
	return payments
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

// UpdatePaymentTrace for fetched ref number
func (p *PaymentRepository) UpdatePaymentTrace(payment *Payment, traceNum string) (*Payment, error) {
	payment.TraceNum = &traceNum
	result := p.db.Save(payment)
	return payment, result.Error
}

// UpdatePaymentRefStatus for returned and fetched Trace Number
func (p *PaymentRepository) UpdatePaymentRefStatus(payment *Payment, refNum, status string) (*Payment, error) {
	if status != "" {
		payment.Status = status
	}

	if refNum != "" {
		payment.RefNum = &refNum
	}

	result := p.db.Save(payment)
	return payment, result.Error
}
