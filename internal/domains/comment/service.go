package comment

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
		return nil, NoProductsFound
	}

	return comments, nil
}

// CreateComment and return it
func (c *CommentService) CreateComment(userId, productId int, description string) (*Comment, error) {
	//TODO implement me
	panic("implement me")
}

// DeleteComment and which exists in db
func (c *CommentService) DeleteComment(cmId int) error {
	//TODO implement me
	panic("implement me")
}
