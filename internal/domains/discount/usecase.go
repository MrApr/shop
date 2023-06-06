package discount

import "context"

// DiscountUseCase is the type which implements DiscountUseCaseInterface
type DiscountUseCase struct {
	sv DiscountServiceInterface
}

// NewUseCase and return it
func NewUseCase(sv DiscountServiceInterface) DiscountUseCaseInterface {
	return &DiscountUseCase{
		sv: sv,
	}
}

// GetAllDiscounts and return the
func (d *DiscountUseCase) GetAllDiscounts(ctx context.Context, request *GetAllDiscountCodesRequest) ([]DiscountCode, error) {
	return d.sv.GetAllDiscounts(request.From, request.Limit)
}

// GetDiscountById and return it
func (d *DiscountUseCase) GetDiscountById(ctx context.Context, id int) (*DiscountCode, error) {
	return d.sv.GetDiscountById(id)
}

// GetDiscountByCode and return it
func (d *DiscountUseCase) GetDiscountByCode(ctx context.Context, code string) (*DiscountCode, error) {
	return d.sv.GetDiscountByCode(code)
}
