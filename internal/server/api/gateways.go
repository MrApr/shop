package api

import (
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"shop/internal/domains/gateways"
)

// gatewayEchoHandler is the type which attaches rest api end points to domain functions
type gatewayEchoHandler struct {
	uC gateways.GatewayUseCaseInterface
}

// AttachGatewayHandlerToGatewayDomain for working with rest Apis
func AttachGatewayHandlerToGatewayDomain(engine *echo.Echo, db *gorm.DB) {
	uC := gateways.NewGatewayUseCase(gateways.NewGatewayService(gateways.NewGatewayRepository(db)))

	setupGatewayHandler(engine, &gatewayEchoHandler{
		uC: uC,
	})
}

// setupGatewayHandler which are accessible through http URI
func setupGatewayHandler(engine *echo.Echo, handler *gatewayEchoHandler) {
	gatewayRouter := engine.Group("gateways")
	gatewayRouter.GET("", handler.GetAllGateways)
}

// GetAllGateways and return it
func (gH *gatewayEchoHandler) GetAllGateways(e echo.Context) error {
	getAllGateWayRequests := gH.getDefaultRequest()
	ctx := context.Background()

	fetchedGateways, err := gH.uC.GetAllGateways(ctx, getAllGateWayRequests)
	if err != nil {
		return e.JSON(http.StatusNotFound, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(fetchedGateways, nil))
}

// getDefaultRequest because gateway types are hided from user currently
func (gH *gatewayEchoHandler) getDefaultRequest() *gateways.GetAllGatewaysRequest {
	trueVal := true
	return &gateways.GetAllGatewaysRequest{
		TypeId:      1,
		OnlyActives: &trueVal,
	}
}
