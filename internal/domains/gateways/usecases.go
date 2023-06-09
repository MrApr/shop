package gateways

import "context"

// GatewayTypeUseCase is the struct which implements GatewayTypesUseCaseInterface
type GatewayTypeUseCase struct {
	sv GatewayTypesServiceInterface
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
