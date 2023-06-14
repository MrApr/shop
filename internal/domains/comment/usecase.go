package comment

import "context"

// CommentUseCase struct which implements CommentUseCaseInterface
type CommentUseCase struct {
	sv CommentServiceInterface
}

// NewCommentUseCase and return it
func NewCommentUseCase(sv CommentServiceInterface) CommentUseCaseInterface {
	return &CommentUseCase{
		sv: sv,
	}
}

// GetAllActiveComments and return it
func (c *CommentUseCase) GetAllActiveComments(ctx context.Context, productId int) ([]Comment, error) {
	return c.sv.GetAllActiveComments(productId)
}

// CreateComment and return it
func (c *CommentUseCase) CreateComment(ctx context.Context, request *CreateCommentRequest) (*Comment, error) {
	//TODO implement me
	panic("implement me")
}

// DeleteComment which already exists in database
func (c *CommentUseCase) DeleteComment(ctx context.Context, cmId int) error {
	//TODO implement me
	panic("implement me")
}
