package products

import "context"

// CategoriesRepositoryInterface defines set of abstract methods for every type, who is going to be Category Repository
type CategoriesRepositoryInterface interface {
	GetAllCategories(title *string, parentCatId, typeId, limit *int, offset int) []Category
}

// CategoryServiceInterface defines set of abstract methods for every type, who is going to be Category Service
type CategoryServiceInterface interface {
	GetAllCategories(title *string, parentCatId, typeId, limit *int, offset int) ([]Category, error)
}

// CategoryUseCaseInterface defines set of abstract methods for every type, who is going to be Category Use-Case
type CategoryUseCaseInterface interface {
	GetAllCategories(ctx context.Context, request *GetAllCategoriesRequest) ([]Category, error)
}

// TypeRepositoryInterface defines set of abstract methods for every type, who is going to be Type repository
type TypeRepositoryInterface interface {
	GetAllTypes(name *string, limit, offset int) []Type
}

// TypeServiceInterface defines set of abstract methods for every type who is going to be Type Service
type TypeServiceInterface interface {
	GetAllTypes(name *string, limit, offset int) ([]Type, error)
}

// TypeUseCaseInterface defines set of abstract methods for type use case, which every Type use case should implement it
type TypeUseCaseInterface interface {
	GetAllTypes(ctx context.Context, request *GetAllTypesRequest) ([]Type, error)
}

// ProductsRepositoryInterface defines set of abstract methods for every type who wants to play role as Product repository
type ProductsRepositoryInterface interface {
	GetAllProducts(categories []int, title, description *string, minWeight, maxWeight *int, minPrice, maxPrice *float64) []Product
	GetProduct(id int) *Product
}

// ProductServiceInterface defines set of abstract methods for every type who wants to play role as Product service
type ProductServiceInterface interface {
	GetAllProducts(categories []int, title, description *string, minWeight, maxWeight *int, minPrice, maxPrice *float64) ([]Product, error)
	GetProduct(id int) (*Product, error)
}

// ProductUseCaseInterface defines set of abstract methods for every type who wants to play role as Product service
type ProductUseCaseInterface interface {
	GetAllProducts(ctx context.Context, request *GetAllProductsRequest) ([]Product, error)
	GetProduct(id int) (*Product, error)
}
