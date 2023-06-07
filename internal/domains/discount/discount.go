package discount

import "context"

// DiscountRepositoryInterface defines set of abstract methods for discount repo
type DiscountRepositoryInterface interface {
	GetAllDiscounts(from, to int) []DiscountCode
	GetDiscountById(id int) *DiscountCode
	GetDiscountByCode(code string) *DiscountCode
}

// DiscountServiceInterface defines set of abstract methods that each type who wants to be known and treated as discount service should obey
type DiscountServiceInterface interface {
	GetAllDiscounts(from, to int) ([]DiscountCode, error)
	GetDiscountById(id int) (*DiscountCode, error)
	GetDiscountByCode(code string) (*DiscountCode, error)
}

// DiscountUseCaseInterface defines set of abstract methods that each type who wants to be known and treated as discount use-case should obey
type DiscountUseCaseInterface interface {
	GetAllDiscounts(ctx context.Context, request *GetAllDiscountCodesRequest) ([]DiscountCode, error)
	GetDiscountById(ctx context.Context, id int) (*DiscountCode, error)
	GetDiscountByCode(ctx context.Context, code string) (*DiscountCode, error)
}

//Todo Should get fetched only by authorized users ( Middleware limitation )
