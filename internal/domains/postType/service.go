package postType

// PostTypeService implements PostTypeServiceInterface
type PostTypeService struct {
	repo PostTypeRepositoryInterface
}

// NewPostTypeService and return it
func NewPostTypeService(repo PostTypeRepositoryInterface) PostTypeServiceInterface {
	return &PostTypeService{
		repo: repo,
	}
}

// GetAllPostTypes and return them
func (p *PostTypeService) GetAllPostTypes() ([]PostType, error) {
	postTypes := p.repo.GetAllPostTypes()
	if len(postTypes) == 0 {
		return nil, NotTypesFound
	}

	return postTypes, nil
}

// PostTypeExists or not based on given it
func (p *PostTypeService) PostTypeExists(id int) error {
	//TODO implement me
	panic("implement me")
}
