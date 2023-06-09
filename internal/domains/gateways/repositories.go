package gateways

import (
	"gorm.io/gorm"
)

// GatewayTypesRepository is a struct which implements GatewayTypesRepositoryInterface
type GatewayTypesRepository struct {
	db *gorm.DB
}

// GatewayRepository is a struct which implements GatewayRepositoryInterface
type GatewayRepository struct {
	db *gorm.DB
}

// NewGatewayTypeRepo and return it
func NewGatewayTypeRepo(db *gorm.DB) GatewayTypesRepositoryInterface {
	return &GatewayTypesRepository{
		db: db,
	}
}

// GetAllTypes and return them
func (g *GatewayTypesRepository) GetAllTypes() []GatewayType {
	var gatewayTypes []GatewayType
	g.db.Where("status = ?", true).Find(&gatewayTypes)
	return gatewayTypes
}

// NewGatewayRepository and return it
func NewGatewayRepository(db *gorm.DB) GatewayRepositoryInterface {
	return &GatewayRepository{
		db: db,
	}
}

// GetAllGateways and return it bases on requested type and status
func (g *GatewayRepository) GetAllGateways(typeId int, onlyActives bool) []GateWay {
	var gateways []GateWay

	db := g.db.Where("gateway_type_id = ?", typeId)
	if onlyActives {
		db = db.Where("status = ?", true)
	}
	db.Find(&gateways)

	return gateways
}
