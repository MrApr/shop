package discount

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

// TestDiscountRepository_GetAllDiscounts functionality
func TestDiscountRepository_GetAllDiscounts(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "setting up database connection failed")

	testingObjCounts := 2

	repo := createRepository(db)
	mockedDiscounts := mockAndInsertDiscountCodes(db, testingObjCounts)
	defer destructDiscountCodes(db, mockedDiscounts)
	assert.Equal(t, len(mockedDiscounts), testingObjCounts, "Initializing discount codes mocked objects failed !Required amounts and created amounts are not equal")

	result := repo.GetAllDiscounts(0, testingObjCounts)
	assert.Equal(t, len(result), testingObjCounts, "fetching discount codes failed. fetched items length with mocked items length are not equal")

	assertDiscountCodes(t, mockedDiscounts, result)
}

// TestDiscountRepository_GetDiscountByCode functionality
func TestDiscountRepository_GetDiscountByCode(t *testing.T) {

}

// TestDiscountRepository_GetDiscountById functionality
func TestDiscountRepository_GetDiscountById(t *testing.T) {

}

// setupDbConnection and run migration
func setupDbConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(DiscountCode{})
	return db, err
}

// createRepository and return it
func createRepository(db *gorm.DB) DiscountRepositoryInterface {
	return NewRepository(db)
}

// mockAndInsertDiscountCodes and return them
func mockAndInsertDiscountCodes(db *gorm.DB, count int) []DiscountCode {
	discounts := make([]DiscountCode, 0, count)

	for i := 0; i < count; i++ {
		mockedDiscount := mockDiscountCode()
		res := db.Create(mockedDiscount)
		if res.Error != nil {
			continue
		}

		discounts = append(discounts, *mockedDiscount)
	}

	return discounts
}

// mockDiscountCode and return it
func mockDiscountCode() *DiscountCode {
	randInt := rand.Int()
	randCode := fmt.Sprintf("%s%d", "discountCode", randInt)

	return &DiscountCode{
		Title:           "Test discount code",
		Code:            randCode,
		DiscountPercent: rand.Int(),
		Status:          true,
	}
}

// destructDiscountCodes which created during test
func destructDiscountCodes(db *gorm.DB, discounts []DiscountCode) {
	for _, discount := range discounts {
		db.Unscoped().Delete(discount)
	}
}

// assertDiscountCodes and check whether they are equal or not
func assertDiscountCodes(t *testing.T, mockedDiscounts, fetchedDiscounts []DiscountCode) {
	for index := range mockedDiscounts {
		assert.Equal(t, mockedDiscounts[index].Title, fetchedDiscounts[index].Title, "Discounts are not equal")
		assert.Equal(t, mockedDiscounts[index].Code, fetchedDiscounts[index].Code, "Discounts are not equal")
		assert.Equal(t, mockedDiscounts[index].DiscountPercent, fetchedDiscounts[index].DiscountPercent, "Discounts are not equal")
		assert.Equal(t, mockedDiscounts[index].Status, fetchedDiscounts[index].Status, "Discounts are not equal")
	}
}
