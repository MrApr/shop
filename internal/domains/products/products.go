package products

import "context"

// CategoriesRepositoryInterface defines set of abstract methods for every type, who is going to be Category Repository
type CategoriesRepositoryInterface interface {
	GetAllCategories() []Category
}

// CategoryServiceInterface defines set of abstract methods for every type, who is going to be Category Service
type CategoryServiceInterface interface {
	GetAllCategories() ([]Category, error)
}

// CategoryUseCaseInterface defines set of abstract methods for every type, who is going to be Category Use-Case
type CategoryUseCaseInterface interface {
	GetAllCategories(ctx context.Context) ([]Category, error)
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
