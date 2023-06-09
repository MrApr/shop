package gateways

import "context"

// GatewayTypesRepositoryInterface defines set of methods for every type who wants to play role as gateway type repository
type GatewayTypesRepositoryInterface interface {
	GetAllTypes() []GatewayType
}

// GatewayTypesServiceInterface defines set of methods for every type who wants to play role as gateway type service
type GatewayTypesServiceInterface interface {
	GetAllTypes() ([]GatewayType, error)
}

// GatewayTypesUseCaseInterface defines set of methods for every type who wants to play role as gateway type use case
type GatewayTypesUseCaseInterface interface {
	GetAllTypes(ctx context.Context) ([]GatewayType, error)
}

// GatewayRepositoryInterface defines set of abstract methods for every type who wants to play role as gateway repository
type GatewayRepositoryInterface interface {
	GetAllGateways(onlyActives bool) []GateWay
}

// GatewayServiceInterface defines set of abstract methods for every type who wants to play role as gateway service
type GatewayServiceInterface interface {
	GetAllGateways() ([]GateWay, error)
}

// GatewayUseCaseInterface defines set of abstract methods for every type who wants to play role as gateway service
type GatewayUseCaseInterface interface {
	GetAllGateways(ctx context.Context) ([]GateWay, error)
}
