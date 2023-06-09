package gateways

import "time"

// GatewayType represents gateways entity type
type GatewayType struct {
	Id        int        `json:"id" gorm:"primaryKey"`
	Title     string     `json:"title"`
	Status    bool       `json:"status" gorm:"index"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// GateWay represents Gateway entity which handles payments in system
type GateWay struct {
	Id            int          `json:"id" gorm:"primaryKey"`
	Name          string       `json:"name"`
	GatewayTypeId int          `json:"gateway_type_id"`
	Type          *GatewayType `json:"GatewayType" gorm:"foreignKey:GatewayTypeId;references:Id"`
	Token         string       `json:"token" gorm:"index"`
	Status        bool         `json:"status"  gorm:"index"`
	CreatedAt     *time.Time   `json:"created_at"`
	UpdateAt      *time.Time   `json:"update_at"`
	DeletedAt     *time.Time   `json:"deleted_at"`
}

// TableName overrides table name in gorm db
func (GateWay) TableName() string {
	return "gateways"
}

// GetAllGatewaysRequest represents request structure for fetching all gateways from system
type GetAllGatewaysRequest struct {
	TypeId      int   `json:"type_id" validate:"required,min=1"`
	OnlyActives *bool `json:"only_actives,omitempty" validate:"omitempty,boolean"`
}
