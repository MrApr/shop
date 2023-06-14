package comment

import "context"

// CommentRepositoryInterface defines set of methods which every type who wants to play role as comment repo should obey
type CommentRepositoryInterface interface {
	GetComment(cmId int) *Comment
	GetAllActiveComments(productId int) []Comment
	CreateComment(comment *Comment) error
	DeleteComment(comment *Comment) error
}

// CommentServiceInterface defines set of methods which every type who wants to play role as comment service should obey
type CommentServiceInterface interface {
	GetAllActiveComments(productId int) ([]Comment, error)
	CreateComment(userId, productId int, description string) (*Comment, error)
	DeleteComment(cmId int) error
}

// CommentUseCaseInterface defines set of methods which every type who wants to play role as comment use-case should obey
type CommentUseCaseInterface interface {
	GetAllActiveComments(ctx context.Context, productId int) ([]Comment, error)
	CreateComment(ctx context.Context, token string, request *CreateCommentRequest) (*Comment, error)
	DeleteComment(ctx context.Context, cmId int) error
}
