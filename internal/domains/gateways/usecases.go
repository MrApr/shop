package gateways

import "context"

// GatewayTypeUseCase is the struct which implements GatewayTypesUseCaseInterface
type GatewayTypeUseCase struct {
	sv GatewayTypesServiceInterface
}

// GatewayUseCase is the struct which implements GatewayUseCaseInterface
type GatewayUseCase struct {
	sv GatewayServiceInterface
}

func NewGatewayTypeUseCase(sv GatewayTypesServiceInterface) GatewayTypesUseCaseInterface {
	return &GatewayTypeUseCase{
		sv: sv,
	}
}

// GetAllTypes and return them
func (g *GatewayTypeUseCase) GetAllTypes(ctx context.Context) ([]GatewayType, error) {
	return g.sv.GetAllTypes()
}

// NewGatewayUseCase and return it
func NewGatewayUseCase(sv GatewayServiceInterface) GatewayUseCaseInterface {
	return &GatewayUseCase{
		sv: sv,
	}
}

// GetAllGateways and return them
func (g *GatewayUseCase) GetAllGateways(ctx context.Context, request *GetAllGatewaysRequest) ([]GateWay, error) {
	onlyActives := *request.OnlyActives

	return g.sv.GetAllGateways(request.TypeId, onlyActives)
}
