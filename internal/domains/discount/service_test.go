package discount

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

// TestDiscountService_GetAllDiscounts functionality
func TestDiscountService_GetAllDiscounts(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "setting up database connection failed")
	testingObjCounts := 2

	sv := createDiscountService(db)

	_, err = sv.GetAllDiscounts(0, testingObjCounts)
	assert.Error(t, err, "Fetching empty discounts failed, expected error")
	assert.ErrorIs(t, err, NoDiscountsExists, "Fetching empty discounts failed, expected error")

	mockedDiscounts := mockAndInsertDiscountCodes(db, testingObjCounts)
	defer destructDiscountCodes(db, mockedDiscounts)
	assert.Equal(t, len(mockedDiscounts), testingObjCounts, "Initializing discount codes mocked objects failed !Required amounts and created amounts are not equal")

	result, err := sv.GetAllDiscounts(0, testingObjCounts)
	assert.Equal(t, len(result), testingObjCounts, "fetching discount codes failed. fetched items length with mocked items length are not equal")
	assert.NoError(t, err, "Fetching discounts with id from discount service failed")

	assertDiscountCodes(t, mockedDiscounts, result)
}

// TestDiscountService_GetDiscountByCode functionality
func TestDiscountService_GetDiscountByCode(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "setting up database connection failed")
	testingObjCounts := 1

	sv := createDiscountService(db)

	mockedDiscounts := mockAndInsertDiscountCodes(db, testingObjCounts)
	defer destructDiscountCodes(db, mockedDiscounts)
	assert.Equal(t, len(mockedDiscounts), testingObjCounts, "Initializing discount codes mocked objects failed !Required amounts and created amounts are not equal")

	wrongCode := "wrongDiscountCode"
	_, err = sv.GetDiscountByCode(wrongCode)
	assert.Error(t, err, "Expected error on fetching discount with wrong code")
	assert.ErrorIs(t, err, DiscountNotFound, "Expected error on fetching discount with wrong code")

	result, err := sv.GetDiscountByCode(mockedDiscounts[0].Code)
	assert.NoError(t, err, "Fetching discounts with code from discount service failed")

	assertDiscountCodes(t, mockedDiscounts, []DiscountCode{*result})
}

// TestDiscountService_GetDiscountByCode functionality
func TestDiscountService_GetDiscountById(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "setting up database connection failed")
	testingObjCounts := 1

	sv := createDiscountService(db)

	mockedDiscounts := mockAndInsertDiscountCodes(db, testingObjCounts)
	defer destructDiscountCodes(db, mockedDiscounts)
	assert.Equal(t, len(mockedDiscounts), testingObjCounts, "Initializing discount codes mocked objects failed !Required amounts and created amounts are not equal")

	randWrongId := rand.Int()
	_, err = sv.GetDiscountById(randWrongId)
	assert.Error(t, err, "Expected error on fetching discount with wrong id")
	assert.ErrorIs(t, err, DiscountNotFound, "Expected error on fetching discount with wrong id")

	result, err := sv.GetDiscountById(mockedDiscounts[0].Id)
	assert.NoError(t, err, "Fetching all discounts from discount service failed")

	assertDiscountCodes(t, mockedDiscounts, []DiscountCode{*result})
}

// createDiscountService and return it
func createDiscountService(db *gorm.DB) DiscountServiceInterface {
	return NewService(createRepository(db))
}
