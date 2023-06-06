package discount

// DiscountService is the type which Implements DiscountServiceInterface
type DiscountService struct {
	repo DiscountRepositoryInterface
}

// NewService and return it
func NewService(repo DiscountRepositoryInterface) DiscountServiceInterface {
	return &DiscountService{
		repo: repo,
	}
}

// GetAllDiscounts and return them if any exists
func (d *DiscountService) GetAllDiscounts(from, to int) ([]DiscountCode, error) {
	discounts := d.repo.GetAllDiscounts(from, to)
	if len(discounts) == 0 {
		return nil, NoDiscountsExists
	}

	return discounts, nil
}

// GetDiscountById and return it if any exists
func (d *DiscountService) GetDiscountById(id int) (*DiscountCode, error) {
	//TODO implement me
	panic("implement me")
}

// GetDiscountByCode and return it if any exists
func (d *DiscountService) GetDiscountByCode(code string) (*DiscountCode, error) {
	//TODO implement me
	panic("implement me")
}
