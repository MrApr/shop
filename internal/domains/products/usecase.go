package products

import "context"

// CategoriesUseCase is categories use case struct
type CategoriesUseCase struct {
	sv CategoryServiceInterface
}

// NewCategoryUseCase and return it
func NewCategoryUseCase(sv CategoryServiceInterface) CategoryUseCaseInterface {
	return &CategoriesUseCase{
		sv: sv,
	}
}

// GetAllCategories and return them
func (c *CategoriesUseCase) GetAllCategories(ctx context.Context) ([]Category, error) {
	return c.sv.GetAllCategories()
}
