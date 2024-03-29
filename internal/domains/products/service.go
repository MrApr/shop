package products

// CategoryService is the struct which plays role as category's Service
type CategoryService struct {
	repo CategoriesRepositoryInterface
}

// TypeService is the struct which plays a role as Type service
type TypeService struct {
	repo TypeRepositoryInterface
}

// ProductService is a struct which implements ProductServiceInterface
type ProductService struct {
	repo ProductsRepositoryInterface
}

// LikeDislikeService is a struct which implements LikeDislikeServiceInterface
type LikeDislikeService struct {
	repo LikeDislikeRepositoryInterface
}

// NewCategoryService and return it
func NewCategoryService(repo CategoriesRepositoryInterface) CategoryServiceInterface {
	return &CategoryService{
		repo: repo,
	}
}

// GetAllCategories and return them
func (c *CategoryService) GetAllCategories(title *string, parentCatId, typeId, limit *int, offset int) ([]Category, error) {
	cats := c.repo.GetAllCategories(title, parentCatId, typeId, limit, offset)
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

// NewProductsService and return it
func NewProductsService(repo ProductsRepositoryInterface) ProductServiceInterface {
	return &ProductService{
		repo: repo,
	}
}

// GetAllProducts and return them
func (p *ProductService) GetAllProducts(categories []int, title, description *string, minWeight, maxWeight *int, minPrice, maxPrice *float64, limit *int, offset int) ([]Product, error) {
	products := p.repo.GetAllProducts(categories, title, description, minWeight, maxWeight, minPrice, maxPrice, limit, offset)
	if len(products) == 0 {
		return nil, NoProductsFound
	}
	return products, nil
}

// GetProduct and return it based on given ID
func (p *ProductService) GetProduct(id int) (*Product, error) {
	product := p.repo.GetProduct(id)
	if product.Id == 0 {
		return nil, ProductNotFound
	}

	return product, nil
}

// UpdateProductInventory with new amount of that
func (p *ProductService) UpdateProductInventory(productId, newInventory int) (*Product, error) {
	product := p.repo.GetProduct(productId)
	if product.Id == 0 {
		return nil, ProductNotFound
	}
	product.Amount = newInventory

	updateResult := p.repo.UpdateProduct(product)
	if updateResult != nil {
		return nil, updateResult
	}

	return product, nil
}

// NewLikeDislikeService and return it
func NewLikeDislikeService(repo LikeDislikeRepositoryInterface) LikeDislikeServiceInterface {
	return &LikeDislikeService{
		repo: repo,
	}
}

// LikeProduct and store it if none exists. If exists remove it
func (l *LikeDislikeService) LikeProduct(productId, userId int) error {
	if l.repo.LikeExists(productId, userId) {
		return l.repo.RemoveLike(productId, userId)
	}

	if err := l.removeDislike(productId, userId); err != nil {
		return err
	}

	liked := l.repo.LikeProduct(productId, userId)
	if liked.ProductId == productId && liked.UserId == userId {
		return nil
	}

	return InternalServerError
}

// DislikeProduct and store it if none exists. If exists remove it
func (l *LikeDislikeService) DislikeProduct(productId, userId int) error {
	if l.repo.DisLikeExists(productId, userId) {
		return l.repo.RemoveDislike(productId, userId)
	}

	if err := l.removeLike(productId, userId); err != nil {
		return err
	}

	disliked := l.repo.DislikeProduct(productId, userId)
	if disliked.ProductId == productId && disliked.UserId == userId {
		return nil
	}

	return InternalServerError
}

// removeDislike when it exists
func (l *LikeDislikeService) removeDislike(productId, userId int) error {
	if l.repo.DisLikeExists(productId, userId) {
		return l.repo.RemoveDislike(productId, userId)
	}
	return nil
}

// removeLike when it exists
func (l *LikeDislikeService) removeLike(productId, userId int) error {
	if l.repo.LikeExists(productId, userId) {
		return l.repo.RemoveLike(productId, userId)
	}
	return nil
}
