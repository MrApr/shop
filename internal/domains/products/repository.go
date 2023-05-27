package products

import "gorm.io/gorm"

// CategoryRepository which implements repository interface of category
type CategoryRepository struct {
	db *gorm.DB
}

// TypeRepository implements TypeRepositoryInterface
type TypeRepository struct {
	db *gorm.DB
}

// NewCategoryRepo creates and returns a new instance of category repository
func NewCategoryRepo(db *gorm.DB) CategoriesRepositoryInterface {
	return &CategoryRepository{
		db: db,
	}
}

// GetAllCategories and return them
func (c *CategoryRepository) GetAllCategories(title *string, parentCatId, typeId, limit *int, offset int) []Category {
	var categories []Category
	db := c.db

	if title != nil {
		db = db.Where("title LIKE ?", *title)
	}

	if parentCatId != nil {
		db = db.Where("parent_cat_id = ?", *parentCatId)
	}

	if typeId != nil {
		db = db.Where("type_id = ?", *typeId)
	}

	if limit != nil {
		db = db.Limit(*limit)
	}

	db.Preload("Type").Offset(offset).Find(&categories)
	return categories
}

// TypesRepository implements type Repository interface
type TypesRepository struct {
	db *gorm.DB
}

// NewTypeRepo and return it
func NewTypeRepo(db *gorm.DB) TypeRepositoryInterface {
	return &TypeRepository{
		db: db,
	}
}

// GetAllTypes and return them based on required values
func (t *TypeRepository) GetAllTypes(name *string, limit, offset int) []Type {
	var typesSlice []Type

	db := t.db

	if name != nil {
		db = db.Where("name LIKE ?", *name)
	}

	db.Limit(limit).Offset(offset).Find(&typesSlice)

	return typesSlice
}
