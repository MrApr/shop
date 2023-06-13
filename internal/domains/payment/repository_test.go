package payment

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

// TestPaymentRepository_GetPayment functionality
func TestPaymentRepository_GetPayment(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Cannot make connection to database")

	repo := createRepository(conn)
	payments := mockAndInsertData(conn, 1, false)
	defer destructPayments(conn, payments)
	assert.Equal(t, 1, len(payments), "Mocking payments failed")

	result, err := repo.GetPayment(payments[0].Id)
	assert.NoError(t, err, "Fetching single payment by id failed")

	assertTwoPayments(t, result, &payments[0])
	assert.NotEmptyf(t, result.CreatedAt, "Payment fetch failed")
	assert.NotEmptyf(t, result.UpdatedAt, "Payment fetch failed")
}

// TestPaymentRepository_GetUserLastPayment functionality
func TestPaymentRepository_GetUserLastPayment(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Cannot make connection to database")

	repo := createRepository(conn)
	payments := mockAndInsertData(conn, 3, true)
	defer destructPayments(conn, payments)
	assert.Equal(t, 3, len(payments), "Mocking payments failed")

	result, err := repo.GetUserLastPayment(payments[2].UserId, true)
	assert.NoError(t, err, "Fetching single payment by id failed")
	assert.Equal(t, result.Status, PaymentDefaultStatus, "Fetch user last created payment failed")

	result, err = repo.GetUserLastPayment(payments[2].UserId, false)
	assert.NoError(t, err, "Fetching single payment by id failed")
	assertTwoPayments(t, result, &payments[2])
}

// TestPaymentRepository_CreatePayment functionality
func TestPaymentRepository_CreatePayment(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Cannot make connection to database")

	repo := createRepository(conn)
	mockedPayment := mockPaymentInfo()

	result := repo.CreatePayment(mockedPayment)
	defer destructPayments(conn, []Payment{*mockedPayment})
	assert.NoError(t, result, "Payment creation failed")

	assert.NotZero(t, mockedPayment.Id, "Payment creation failed")
	assert.NotNil(t, mockedPayment.CreatedAt, "Payment creation failed")
	assert.NotZero(t, mockedPayment.UpdatedAt, "Payment creation failed")
}

// TestPaymentRepository_UpdatePayment functionality
func TestPaymentRepository_UpdatePaymentRef(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Cannot make connection to database")

	repo := createRepository(conn)
	payments := mockAndInsertData(conn, 1, false)
	defer destructPayments(conn, payments)
	assert.Equal(t, 1, len(payments), "Mocking payments failed")

	newRefNum := "akdkdakdaskdsa"

	newPayment, err := repo.UpdatePaymentRef(&payments[0], newRefNum)
	assert.NoError(t, err, "Payment Update ref-num operation failed")
	assert.Equal(t, newPayment.RefNum, payments[0].RefNum, "Payment Update ref-num operation failed")

	tmpPayment := new(Payment)
	conn.Where("id = ?", payments[0].Id).First(tmpPayment)

	assert.Equal(t, *tmpPayment.RefNum, newRefNum, "Payment Update ref-num operation failed")
}

// TestPaymentRepository_UpdatePayment functionality
func TestPaymentRepository_UpdatePaymentTraceStatus(t *testing.T) {
	conn, err := setupDbConnection()
	assert.NoError(t, err, "Cannot make connection to database")

	//repo := createRepository(conn)
	payments := mockAndInsertData(conn, 1, false)
	defer destructPayments(conn, payments)
	assert.Equal(t, 1, len(payments), "Mocking payments failed")

	//newPayment, err := repo.UpdatePayment(&payments[0], "akdkdakdaskdsa", "akbkasdkasdka", true)
	//assert.NoError(t, err, "Payment Update operation failed")
	//
	//assert.NotNil(t, newPayment.RefNumber, "Payment Update operation failed")
	//assert.NotNil(t, newPayment.TraceNumber, "Payment Update operation failed")
	//assert.Equal(t, *newPayment.TraceNumber, "akbkasdkasdka", "Payment Update operation failed")
	//assert.Equal(t, *newPayment.RefNumber, "akdkdakdaskdsa", "Payment Update operation failed")
	//assert.Equal(t, newPayment.Status, true, "Payment Update operation failed")
}

// mockAndInsertData and db and return them
func mockAndInsertData(db *gorm.DB, count int, withSuccessStatTest bool) []Payment {
	var createdPayments []Payment

	for i := 0; i < count; i++ {
		mockedPayment := mockPaymentInfo()
		if withSuccessStatTest && i == (count-1) {
			mockedPayment.Status = PaymentSuccessStatus
		}

		res := db.Create(mockedPayment)
		if res.Error == nil {
			createdPayments = append(createdPayments, *mockedPayment)
		}
	}

	return createdPayments
}

// mockPaymentInfo and return it for testing purpose
func mockPaymentInfo() *Payment {
	refNumber := fmt.Sprintf("%s%d", "bbbbb", rand.Int())
	tractNumber := fmt.Sprintf("%s%d", "aaaaa", rand.Int())
	return &Payment{
		UserId:     rand.Int(),
		BasketId:   rand.Int(),
		AddressId:  rand.Int(),
		DiscountId: rand.Int(),
		GatewayId:  rand.Int(),
		PostTypeId: rand.Int(),
		TotalPrice: rand.Float64(),
		RefNum:     &refNumber,
		TraceNum:   &tractNumber,
		Status:     PaymentDefaultStatus,
	}
}

// createRepository and return it
func createRepository(db *gorm.DB) PaymentRepoContract {
	return NewPaymentRepo(db)
}

// setupDbConnection and run migration
func setupDbConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(Payment{})
	return db, err
}

// assertTwoPayments together to check whether they are equal or not
func assertTwoPayments(t *testing.T, result, mockedPayment *Payment) {
	assert.Equal(t, result.UserId, mockedPayment.UserId, "Fetched and mocked payments are not equal")
	assert.Equal(t, result.BasketId, mockedPayment.BasketId, "Fetched and mocked payments are not equal")
	assert.Equal(t, result.AddressId, mockedPayment.AddressId, "Fetched and mocked payments are not equal")
	assert.Equal(t, result.DiscountId, mockedPayment.DiscountId, "Fetched and mocked payments are not equal")
	assert.Equal(t, result.GatewayId, mockedPayment.GatewayId, "Fetched and mocked payments are not equal")
	assert.Equal(t, result.PostTypeId, mockedPayment.PostTypeId, "Fetched and mocked payments are not equal")
	assert.Equal(t, result.TotalPrice, mockedPayment.TotalPrice, "Fetched and mocked payments are not equal")
	assert.Equal(t, result.Status, mockedPayment.Status, "Fetched and mocked payments are not equal")
	assert.Equal(t, result.RefNum, mockedPayment.RefNum, "Fetched and mocked payments are not equal")
	assert.Equal(t, result.TraceNum, mockedPayment.TraceNum, "Fetched and mocked payments are not equal")
}

// destructPayments and delete them (roll back func)
func destructPayments(db *gorm.DB, payments []Payment) {
	for _, payment := range payments {
		db.Unscoped().Delete(payment)
	}
}
