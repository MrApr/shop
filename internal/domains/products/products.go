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
