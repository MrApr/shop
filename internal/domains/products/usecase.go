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

// LikeDislikeUseCase is the type which implements LikeDislikeUseCaseInterface
type LikeDislikeUseCase struct {
	sv LikeDislikeServiceInterface
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

// NewLikeDislikeUseCase and return it
func NewLikeDislikeUseCase(sv LikeDislikeServiceInterface) LikeDislikeUseCaseInterface {
	return &LikeDislikeUseCase{
		sv: sv,
	}
}

// LikeProduct and store it in db
func (l *LikeDislikeUseCase) LikeProduct(ctx context.Context, token string, request *LikeDislikeRequest) error {
	return l.sv.LikeProduct(request.ProductId, request.UserId)
}

// DislikeProduct and store it in db
func (l *LikeDislikeUseCase) DislikeProduct(ctx context.Context, token string, request *LikeDislikeRequest) error {
	//TODO implement me
	panic("implement me")
}
