package products

// CategoryService is the struct which plays role as category's Service
type CategoryService struct {
	repo CategoriesRepositoryInterface
}

// TypeService is the struct which plays a role as Type service
type TypeService struct {
	repo TypeRepositoryInterface
}

// NewCategoryService and return it
func NewCategoryService(repo CategoriesRepositoryInterface) CategoryServiceInterface {
	return &CategoryService{
		repo: repo,
	}
}

// GetAllCategories and return them
func (c *CategoryService) GetAllCategories() ([]Category, error) {
	cats := c.repo.GetAllCategories()
	if len(cats) == 0 {
		return nil, NoCategoriesFound
	}

	return cats, nil
}

// NewTypeService and return it
func NewTypeService(repo TypeRepositoryInterface) TypeServiceInterface {
	return &TypeService{
		repo: repo,
	}
}

// GetAllTypes and return them
func (t *TypeService) GetAllTypes(name *string, limit, offset int) ([]Type, error) {
	results := t.repo.GetAllTypes(name, limit, offset)
	if len(results) == 0 {
		return nil, NoTypesFound
	}
	return results, nil
}
