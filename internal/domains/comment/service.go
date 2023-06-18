package comment

// CommentDefaultStatus defines default status for comment
const CommentDefaultStatus bool = true

// CommentService implements CommentServiceInterface
type CommentService struct {
	repo CommentRepositoryInterface
}

// NewCommentService and return it
func NewCommentService(repo CommentRepositoryInterface) CommentServiceInterface {
	return &CommentService{
		repo: repo,
	}
}

// GetAllActiveComments and return them
func (c *CommentService) GetAllActiveComments(productId int) ([]Comment, error) {
	comments := c.repo.GetAllActiveComments(productId)

	if len(comments) == 0 {
		return nil, NoCommentsFound
	}

	return comments, nil
}

// CreateComment and return it
func (c *CommentService) CreateComment(userId, productId int, description string) (*Comment, error) {
	cm := &Comment{
		ProductId:   productId,
		UserId:      userId,
		Description: description,
		Status:      CommentDefaultStatus,
	}

	err := c.repo.CreateComment(cm)
	return cm, err
}

// DeleteComment and which exists in db
func (c *CommentService) DeleteComment(cmId, userId int) error {
	comment := c.repo.GetComment(cmId)
	if comment.Id == 0 {
		return CommentNotFound
	}

	if comment.UserId != userId {
		return OperationNotAllowed
	}

	return c.repo.DeleteComment(comment)
}
