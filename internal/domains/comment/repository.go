package comment

import "gorm.io/gorm"

// CommentRepository implements CommentRepositoryInterface
type CommentRepository struct {
	db *gorm.DB
}

// NewCommentRepository and return it
func NewCommentRepository(db *gorm.DB) CommentRepositoryInterface {
	return &CommentRepository{
		db: db,
	}
}

// GetComment and return it
func (c *CommentRepository) GetComment(cmId int) *Comment {
	comment := new(Comment)
	c.db.Where("id = ?", cmId).First(comment)
	return comment
}

// GetAllActiveComments and return them
func (c *CommentRepository) GetAllActiveComments(productId int) []Comment {
	var comments []Comment
	c.db.Where("product_id = ?", productId).Where("status = ?", true).Find(&comments)
	return comments
}

// CreateComment and return it
func (c *CommentRepository) CreateComment(comment *Comment) error {
	return c.db.Create(comment).Error
}

// DeleteComment and return it
func (c *CommentRepository) DeleteComment(comment *Comment) error {
	return c.db.Delete(comment).Error
}
