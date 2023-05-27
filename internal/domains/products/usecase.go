package products

import "context"

// CategoriesUseCase is categories use case struct
type CategoriesUseCase struct {
	sv CategoryServiceInterface
}

// TypesUseCase is Type UseCase handler
type TypesUseCase struct {
	sv TypeServiceInterface
}

// ProductUseCase is the type which implements Products use case interface
type ProductUseCase struct {
	sv ProductServiceInterface
}

// defaultOffset defines default starting point for requests
const defaultOffset int = 0

// NewCategoryUseCase and return it
func NewCategoryUseCase(sv CategoryServiceInterface) CategoryUseCaseInterface {
	return &CategoriesUseCase{
		sv: sv,
	}
}

// GetAllCategories and return them
func (c *CategoriesUseCase) GetAllCategories(ctx context.Context, request *GetAllCategoriesRequest) ([]Category, error) {
	return c.sv.GetAllCategories(request.Title, request.ParentCatId, request.TypeId, request.Limit, request.Offset)
}

// NewTypeUseCase and return it
func NewTypeUseCase(sv TypeServiceInterface) TypeUseCaseInterface {
	return &TypesUseCase{
		sv: sv,
	}
}

// GetAllTypes and return them
func (t *TypesUseCase) GetAllTypes(ctx context.Context, request *GetAllTypesRequest) ([]Type, error) {
	var offset = defaultOffset
	if request.Offset != nil {
		offset = *request.Offset
	}

	return t.sv.GetAllTypes(request.Name, request.Limit, offset)
}

// NewProductUseCase and return it
func NewProductUseCase(sv ProductServiceInterface) ProductUseCaseInterface {
	return &ProductUseCase{
		sv: sv,
	}
}

// GetAllProducts and return them
func (p *ProductUseCase) GetAllProducts(ctx context.Context, request *GetAllProductsRequest) ([]Product, error) {
	return p.sv.GetAllProducts(request.CategoryIds, request.Title, request.Description, request.MinWeight, request.MaxWeight, request.MinPrice, request.MaxPrice, request.Limit, request.Offset)
}

// GetProduct and return it based on given id
func (p *ProductUseCase) GetProduct(ctx context.Context, id int) (*Product, error) {
	return p.sv.GetProduct(id)
}
