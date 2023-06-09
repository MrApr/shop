package gateways

// GatewayTypeService implements GatewayTypesServiceInterface
type GatewayTypeService struct {
	repo GatewayTypesRepositoryInterface
}

// NewGatewayTypesService and return it
func NewGatewayTypesService(repo GatewayTypesRepositoryInterface) GatewayTypesServiceInterface {
	return &GatewayTypeService{
		repo: repo,
	}
}

// GetAllTypes and return them
func (g *GatewayTypeService) GetAllTypes() ([]GatewayType, error) {
	types := g.repo.GetAllTypes()
	if len(types) == 0 {
		return nil, NoTypesFound
	}
	return types, nil
}
