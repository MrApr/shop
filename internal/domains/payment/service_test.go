package payment

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

// TestPaymentStorageService_GetPayment functionality
func TestPaymentStorageService_GetPayment(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Cannot make connection to database")

	sv := createPaymentService(conn, false)
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

// TestPaymentStorageService_GetUserPayments functionality
func TestPaymentStorageService_GetUserPayments(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Cannot make connection to database")

	sv := createPaymentService(conn, false)
	payments := mockAndInsertData(conn, 2, false)
	defer destructPayments(conn, payments)
	assert.Equal(t, 2, len(payments), "Mocking payments failed")

	result, err := sv.GetUserPayments(payments[0].UserId, 0, 10)
	assert.NoError(t, err, "Fetching user payments failed")

	assert.Equal(t, result[0].UserId, payments[0].UserId, "Fetching user payments failed !")

	randWrongId := rand.Int()
	_, err = sv.GetUserPayments(randWrongId, 0, 10)
	assert.Error(t, err, "Expected error on fetching payment with random wrong id")
	assert.ErrorIs(t, err, PaymentNotFound, "Expected error on fetching payment with random wrong id")

	_, err = sv.GetUserPayments(payments[0].UserId, 20, 10)
	assert.Error(t, err, "Expected error on fetching payment with random wrong id")
	assert.ErrorIs(t, err, PaymentNotFound, "Expected error on fetching payment with random wrong id")
}

// TestPaymentService_CreatePayment functionality
func TestPaymentService_CreatePayment(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Cannot make connection to database")

	sv := createPaymentService(conn, false)
	userId := rand.Int()
	basketId := rand.Int()
	addressId := rand.Int()
	discountId := rand.Int()
	gatewayId := rand.Int()
	postTypeId := rand.Int()
	totalPrice := rand.Float64()

	_, err = sv.CreatePayment(userId, basketId, addressId, discountId, gatewayId, postTypeId, totalPrice)
	assert.Error(t, err, "Expected error for validation function rejection")

	sv = createPaymentService(conn, true)
	result, err := sv.CreatePayment(userId, basketId, addressId, discountId, gatewayId, postTypeId, totalPrice)
	defer destructPayments(conn, []Payment{*result})
	assert.NoError(t, err, "Payment service payment creation error")

	assert.Equal(t, result.UserId, userId, "Payment service payment creation error")
	assert.Equal(t, result.BasketId, basketId, "Payment service payment creation error")
	assert.Equal(t, result.AddressId, addressId, "Payment service payment creation error")
	assert.Equal(t, result.DiscountId, discountId, "Payment service payment creation error")
	assert.Equal(t, result.GatewayId, gatewayId, "Payment service payment creation error")
	assert.Equal(t, result.PostTypeId, postTypeId, "Payment service payment creation error")
	assert.Equal(t, result.TotalPrice, totalPrice, "Payment service payment creation error")
	assert.NotNil(t, result.CreatedAt, "Payment service payment creation error")
	assert.NotNil(t, result.UpdatedAt, "Payment service payment creation error")
}

// createPaymentService and return it
func createPaymentService(db *gorm.DB, withValidFunc bool) PaymentStorageServiceContract {

	validFunc := generateValidationFunction(withValidFunc)

	return NewPaymentStorageService(NewPaymentRepo(db), validFunc, validFunc, validFunc, validFunc)
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

// generateValidationFunction for payment service when a new payment is going to get created
func generateValidationFunction(withValid bool) func(int) error {
	if withValid {
		return func(int) error {
			return nil
		}
	}

	return func(int) error {
		return errors.New("test error validation function")
	}
}
