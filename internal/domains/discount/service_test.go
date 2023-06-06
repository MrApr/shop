package discount

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
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
	assert.NoError(t, err, "Fetching all discounts from discount service failed")

	assertDiscountCodes(t, mockedDiscounts, result)
}

// TestDiscountService_GetDiscountByCode functionality
func TestDiscountService_GetDiscountByCode(t *testing.T) {

}

// TestDiscountService_GetDiscountByCode functionality
func TestDiscountService_GetDiscountById(t *testing.T) {

}

// createDiscountService and return it
func createDiscountService(db *gorm.DB) DiscountServiceInterface {
	return NewService(createRepository(db))
}
