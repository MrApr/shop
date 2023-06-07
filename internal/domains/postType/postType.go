package postType

import "context"

// PostTypeRepositoryInterface defines set of methods for every type who is going to be and used as Repository for post type
type PostTypeRepositoryInterface interface {
	GetAllPostTypes() []PostType
	PostTypeExists(id int) bool
}

// PostTypeServiceInterface defines set of methods for every type who is going to be and used as Service for post type
type PostTypeServiceInterface interface {
	GetAllPostTypes() ([]PostType, error)
	PostTypeExists(id int) bool
}

// PostTypeUseCaseInterface defines set of methods for every type who is going to be and used as Use Case for post type
type PostTypeUseCaseInterface interface {
	GetAllPostTypes(ctx context.Context) ([]PostType, error)
}
