package gateways

import "gorm.io/gorm"

// GatewayTypesRepository is a struct which implements GatewayTypesRepositoryInterface
type GatewayTypesRepository struct {
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
