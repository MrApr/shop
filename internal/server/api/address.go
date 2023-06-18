package api

import (
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"shop/internal/domains/address"
	"shop/pkg/reqTokenHandler"
	"shop/pkg/validation"
)

// addressDeletationMessage in order to sent to client
const addressDeletationMessage string = "Address deleted successfully"

// addressEchoHandler is the type which attaches rest api end points to domain functions
type addressEchoHandler struct {
	uC address.AddressUseCaseInterface
}

// AttachAddressHandlerWithAddressDomain for working with rest Apis
func AttachAddressHandlerWithAddressDomain(engine *echo.Echo, db *gorm.DB) {
	uC := address.NewUseCase(address.NewAddressService(address.NewAddressRepository(db)), nil)
	handler := &addressEchoHandler{
		uC: uC,
	}
	setupAddressRoutes(engine, handler)
}

// setupAddressRoutes which connect with address handler methods
func setupAddressRoutes(engine *echo.Echo, addrHandler *addressEchoHandler) {
	cities := engine.Group("/cities")
	cities.GET("", addrHandler.GetAllCities)

	addressRoutes := engine.Group("/addresses")
	addressRoutes.GET("", addrHandler.GetAllUserAddresses)
}

// GetAllCities and return them
func (a *addressEchoHandler) GetAllCities(e echo.Context) error {
	ctx := context.Background()

	cities, err := a.uC.GetAllCities(ctx)
	if err != nil {
		return e.JSON(http.StatusNotFound, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(cities, nil))
}

// GetAllUserAddresses by calling addresses use-case
func (a *addressEchoHandler) GetAllUserAddresses(e echo.Context) error {
	bearerToken, err := reqTokenHandler.ExtractBearerToken(e.Request())
	if err != nil {
		return e.JSON(http.StatusForbidden, generateResponse(nil, err))
	}

	ctx := context.Background()

	addresses, err := a.uC.GetAllUserAddresses(ctx, bearerToken)
	if err != nil {
		return e.JSON(http.StatusNotFound, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(addresses, nil))
}

// CreateAddress for user in system
func (a *addressEchoHandler) CreateAddress(e echo.Context) error {
	createAddressRequest := new(address.CreateAddressRequest)

	if err := e.Bind(createAddressRequest); err != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, err))
	}

	if errs := validation.Validate(createAddressRequest); errs != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, errs))
	}

	bearerToken, err := reqTokenHandler.ExtractBearerToken(e.Request())
	if err != nil {
		return e.JSON(http.StatusForbidden, generateResponse(nil, err))
	}
	ctx := context.Background()

	newAddress, err := a.uC.CreateAddress(ctx, bearerToken, createAddressRequest)
	if err != nil {
		return e.JSON(http.StatusNotFound, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(newAddress, nil))
}

// UpdateAddress of user which exists in system
func (a *addressEchoHandler) UpdateAddress(e echo.Context) error {
	updateAddrRequest := new(address.UpdateAddressRequest)

	if err := e.Bind(updateAddrRequest); err != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, err))
	}

	if errs := validation.Validate(updateAddrRequest); errs != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, errs))
	}

	bearerToken, err := reqTokenHandler.ExtractBearerToken(e.Request())
	if err != nil {
		return e.JSON(http.StatusForbidden, generateResponse(nil, err))
	}
	ctx := context.Background()

	updatedAddress, err := a.uC.UpdateAddress(ctx, bearerToken, updateAddrRequest)
	if err != nil {
		return e.JSON(http.StatusNotFound, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(updatedAddress, nil))
}

// DeleteAddress which exists in system
func (a *addressEchoHandler) DeleteAddress(e echo.Context) error {
	deleteAddressRequest := new(address.DeleteAddressRequest)

	if err := e.Bind(deleteAddressRequest); err != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, err))
	}

	if errs := validation.Validate(deleteAddressRequest); errs != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, errs))
	}

	bearerToken, err := reqTokenHandler.ExtractBearerToken(e.Request())
	if err != nil {
		return e.JSON(http.StatusForbidden, generateResponse(nil, err))
	}
	ctx := context.Background()

	err = a.uC.DeleteAddress(ctx, bearerToken, deleteAddressRequest)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(nil, addressDeletationMessage))
}
