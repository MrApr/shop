package products

// CategoryService is the struct which plays role as category's repository
type CategoryService struct {
	repo CategoriesRepositoryInterface
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
