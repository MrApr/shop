package postType

import "context"

// PostTypeUseCase implements PostTypeUseCaseInterface
type PostTypeUseCase struct {
	sv PostTypeServiceInterface
}

// NewUseCase and return it
func NewUseCase(sv PostTypeServiceInterface) PostTypeUseCaseInterface {
	return &PostTypeUseCase{
		sv: sv,
	}
}

// GetAllPostTypes and return them
func (p *PostTypeUseCase) GetAllPostTypes(ctx context.Context) ([]PostType, error) {
	return p.sv.GetAllPostTypes()
}
