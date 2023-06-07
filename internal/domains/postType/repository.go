package postType

import (
	"errors"
	"gorm.io/gorm"
)

// PostTypeRepository plays role post type's repository and implements PostTypeRepositoryInterface
type PostTypeRepository struct {
	db *gorm.DB
}

// NewRepository and return it
func NewRepository(db *gorm.DB) PostTypeRepositoryInterface {
	return &PostTypeRepository{
		db: db,
	}
}

// GetAllPostTypes and return them
func (p *PostTypeRepository) GetAllPostTypes() []PostType {
	var postTypes []PostType
	p.db.Find(&postTypes)
	return postTypes
}

// PostTypeExists checks whether post types exist or not
func (p *PostTypeRepository) PostTypeExists(id int) bool {
	postType := new(PostType)

	result := p.db.Where("id = ?", id).First(postType)

	return result.Error == nil && !errors.Is(result.Error, gorm.ErrRecordNotFound)
}
