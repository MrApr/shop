package payment

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

// TestPaymentStorageService_GetPayment functionality
func TestPaymentStorageService_GetPayment(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Cannot make connection to database")

	sv := createPaymentService(conn)
	payments := mockAndInsertData(conn, 1, false)
	defer destructPayments(conn, payments)
	assert.Equal(t, 1, len(payments), "Mocking payments failed")

	result, err := sv.GetPayment(payments[0].Id)
	assert.NoError(t, err, "Fetching single payment by id failed")

	assertTwoPayments(t, result, &payments[0])
	assert.NotEmptyf(t, result.CreatedAt, "Payment fetch failed")
	assert.NotEmptyf(t, result.UpdatedAt, "Payment fetch failed")

	randWrongId := rand.Int()
	_, err = sv.GetPayment(randWrongId)
	assert.Error(t, err, "Expected error on fetching payment with random wrong id")
	assert.ErrorIs(t, err, PaymentNotFound, "Expected error on fetching payment with random wrong id")
}

// TestPaymentService_CreatePayment functionality
func TestPaymentService_CreatePayment(t *testing.T) {
	//conn, err := setupDbConnection()
	//assert.NoError(t, err, "Cannot make connection to database")
	//
	//sv := createPaymentService(conn)
	//mockedPaymentable := mockPaymentable()
	//userId := rand.Int()

	//result, err := sv.CreatePayment(&mockedPaymentable, userId)
	//defer destructPayments(conn, []Payment{*result})
	//assert.NoError(t, err, "Payment service payment creation error")
	//
	//assert.Equal(t, result.UserId, userId, "Payment service payment creation error")
	//assert.Equal(t, result.Amount, mockedPaymentable.Amount(), "Payment service payment creation error")
	//assert.Equal(t, result.PaymentableID, mockedPaymentable.PaymentableId(), "Payment service payment creation error")
	//assert.Equal(t, result.PaymentableType, mockedPaymentable.PaymentableType(), "Payment service payment creation error")
	//assert.NotNil(t, result.CreatedAt, "Payment service payment creation error")
	//assert.NotNil(t, result.UpdatedAt, "Payment service payment creation error")
}

// createPaymentService and return it
func createPaymentService(db *gorm.DB) PaymentStorageServiceContract {
	return NewPaymentStorageService(NewPaymentRepo(db))
}

// customPaymentableType defines a struct which implements paymentable interface
type customPaymentableType struct {
	paymentableType string
	paymentableId   int
	amount          float64
}

// PaymentableType returns it type
func (c *customPaymentableType) PaymentableType() string {
	return c.paymentableType
}

// PaymentableId return it's id
func (c *customPaymentableType) PaymentableId() int {
	return c.paymentableId
}

// Amount returns it's amount
func (c *customPaymentableType) Amount() float64 {
	return c.amount
}

// mockPaymentable and return it
func mockPaymentable() customPaymentableType {
	return customPaymentableType{
		paymentableType: "TestType",
		paymentableId:   rand.Int(),
		amount:          rand.Float64(),
	}
}
