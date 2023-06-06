package discount

import "gorm.io/gorm"

// DiscountRepository implements DiscountRepositoryInterface
type DiscountRepository struct {
	db *gorm.DB
}

// NewRepository and return it
func NewRepository(db *gorm.DB) DiscountRepositoryInterface {
	return &DiscountRepository{
		db: db,
	}
}

// GetAllDiscounts and return them
func (d *DiscountRepository) GetAllDiscounts(from, to int) []DiscountCode {
	var discountCodes []DiscountCode
	d.db.Limit(to).Offset(from).Find(&discountCodes)
	return discountCodes
}

// GetDiscountById and return it
func (d *DiscountRepository) GetDiscountById(id int) *DiscountCode {
	discount := new(DiscountCode)
	d.db.Where("id = ?", id).First(discount)
	return discount
}

// GetDiscountByCode and return it
func (d *DiscountRepository) GetDiscountByCode(code string) *DiscountCode {
	discount := new(DiscountCode)
	d.db.Where("code = ?", code).First(discount)
	return discount
}
