package comment

import (
	"context"
	"shop/pkg/userHandler"
)

// CommentUseCase struct which implements CommentUseCaseInterface
type CommentUseCase struct {
	sv        CommentServiceInterface
	decoderFn func(ctx context.Context, token string) (int, error)
}

// NewCommentUseCase and return it
func NewCommentUseCase(sv CommentServiceInterface, decoderFn func(ctx context.Context, token string) (int, error)) CommentUseCaseInterface {
	if decoderFn == nil {
		decoderFn = userHandler.ExtractUserIdFromToken
	}

	return &CommentUseCase{
		sv:        sv,
		decoderFn: decoderFn,
	}
}

// GetAllActiveComments and return it
func (c *CommentUseCase) GetAllActiveComments(ctx context.Context, productId int) ([]Comment, error) {
	return c.sv.GetAllActiveComments(productId)
}

// CreateComment and return it
func (c *CommentUseCase) CreateComment(ctx context.Context, token string, request *CreateCommentRequest) (*Comment, error) {
	userId, err := c.decoderFn(ctx, token)
	if err != nil {
		return nil, err
	}

	return c.sv.CreateComment(userId, request.ProductId, request.Description)
}

// DeleteComment which already exists in database
func (c *CommentUseCase) DeleteComment(ctx context.Context, cmId int) error {
	return c.sv.DeleteComment(cmId)
}
