package discount

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

// TestDiscountUseCase_GetAllDiscounts functionality
func TestDiscountUseCase_GetAllDiscounts(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "setting up database connection failed")
	testingObjCounts := 2

	uc := createDiscountUseCase(db)
	mockedGetAllRequest := mockGetAllDiscountsRequest()
	ctx := context.Background()

	mockedDiscounts := mockAndInsertDiscountCodes(db, testingObjCounts)
	defer destructDiscountCodes(db, mockedDiscounts)
	assert.Equal(t, len(mockedDiscounts), testingObjCounts, "Initializing discount codes mocked objects failed !Required amounts and created amounts are not equal")

	result, err := uc.GetAllDiscounts(ctx, mockedGetAllRequest)
	assert.Equal(t, len(result), testingObjCounts, "fetching discount codes failed. fetched items length with mocked items length are not equal")
	assert.NoError(t, err, "Fetching discounts with id from discount service failed")

	assertDiscountCodes(t, mockedDiscounts, result)
}

// TestDiscountUseCase_GetDiscountByCode functionality
func TestDiscountUseCase_GetDiscountByCode(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "setting up database connection failed")
	testingObjCounts := 1

	uC := createDiscountUseCase(db)
	ctx := context.Background()

	mockedDiscounts := mockAndInsertDiscountCodes(db, testingObjCounts)
	defer destructDiscountCodes(db, mockedDiscounts)
	assert.Equal(t, len(mockedDiscounts), testingObjCounts, "Initializing discount codes mocked objects failed !Required amounts and created amounts are not equal")

	result, err := uC.GetDiscountByCode(ctx, mockedDiscounts[0].Code)
	assert.NoError(t, err, "Fetching discounts with code from discount service failed")

	assertDiscountCodes(t, mockedDiscounts, []DiscountCode{*result})
}

// TestDiscountUseCase_GetDiscountById functionality
func TestDiscountUseCase_GetDiscountById(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "setting up database connection failed")
	testingObjCounts := 1

	uC := createDiscountUseCase(db)
	ctx := context.Background()

	mockedDiscounts := mockAndInsertDiscountCodes(db, testingObjCounts)
	defer destructDiscountCodes(db, mockedDiscounts)
	assert.Equal(t, len(mockedDiscounts), testingObjCounts, "Initializing discount codes mocked objects failed !Required amounts and created amounts are not equal")

	result, err := uC.GetDiscountById(ctx, mockedDiscounts[0].Id)
	assert.NoError(t, err, "Fetching all discounts from discount service failed")

	assertDiscountCodes(t, mockedDiscounts, []DiscountCode{*result})
}

// createDiscountUseCase and return it for testing purpose
func createDiscountUseCase(db *gorm.DB) DiscountUseCaseInterface {
	return NewUseCase(NewService(NewRepository(db)))
}

// mockGetAllDiscountsRequest and return them
func mockGetAllDiscountsRequest() *GetAllDiscountCodesRequest {
	return &GetAllDiscountCodesRequest{
		From:  0,
		Limit: 10,
	}
}
