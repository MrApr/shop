package gateways

// GatewayTypeService implements GatewayTypesServiceInterface
type GatewayTypeService struct {
	repo GatewayTypesRepositoryInterface
}

// GatewayService implements GatewayServiceInterface
type GatewayService struct {
	repo GatewayRepositoryInterface
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

// NewGatewayService and return it
func NewGatewayService(repo GatewayRepositoryInterface) GatewayServiceInterface {
	return &GatewayService{
		repo: repo,
	}
}

// GetAllGateways and return them based on requested params
func (g *GatewayService) GetAllGateways(typeId int, onlyActives bool) ([]GateWay, error) {
	gateways := g.repo.GetAllGateways(typeId, onlyActives)
	if len(gateways) == 0 {
		return nil, NoGatewaysFound
	}
	return gateways, nil
}
