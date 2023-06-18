package api

import (
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"shop/internal/domains/basket"
	"shop/internal/domains/products"
	"shop/pkg/reqTokenHandler"
	"shop/pkg/validation"
)

// basketDisablationMessage is usable when a new basket gets disabled successfully
const basketDisablationMessage string = "requested basket disabled successfully"

// basketEchoHandler is the type which attaches rest api end points to domain functions
type basketEchoHandler struct {
	uC basket.BasketUseCaseInterface
}

// AttachBasketHandlerWithBasketDomain for working with rest Apis
func AttachBasketHandlerWithBasketDomain(engine *echo.Echo, db *gorm.DB) {
	productService := products.NewProductsService(products.NewProductRepository(db))
	basketUseCase := basket.NewUseCase(basket.NewBasketService(basket.NewBasketRepository(db), productService), nil)
	setupBasketHandlerRoutes(engine, &basketEchoHandler{
		uC: basketUseCase,
	})
}

// setupAddressRoutes which are accessible through http URI
func setupBasketHandlerRoutes(engine *echo.Echo, bH *basketEchoHandler) {
	basketRouter := engine.Group("/baskets")

	basketRouter.GET("/actives", bH.GetUserActiveBasket)
	basketRouter.GET("", bH.GetAllUserBaskets)
	basketRouter.POST("", bH.CreateUserBasket)
	basketRouter.DELETE("", bH.DisableActiveBasket)
	basketRouter.POST("/products", bH.AddProductsToBasket)
	basketRouter.PUT("/products", bH.UpdateBasketProducts)
}

// GetUserActiveBasket and return it
func (bH *basketEchoHandler) GetUserActiveBasket(e echo.Context) error {
	bearerToken, err := reqTokenHandler.ExtractBearerToken(e.Request())
	if err != nil {
		return e.JSON(http.StatusForbidden, generateResponse(nil, err))
	}

	ctx := context.Background()

	activeBasket, err := bH.uC.GetUserActiveBasket(ctx, bearerToken)
	if err != nil {
		return e.JSON(http.StatusNotFound, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(activeBasket, nil))
}

// GetAllUserBaskets and return them
func (bH *basketEchoHandler) GetAllUserBaskets(e echo.Context) error {
	bearerToken, err := reqTokenHandler.ExtractBearerToken(e.Request())
	if err != nil {
		return e.JSON(http.StatusForbidden, generateResponse(nil, err))
	}

	ctx := context.Background()

	baskets, err := bH.uC.GetUserBaskets(ctx, bearerToken)
	if err != nil {
		return e.JSON(http.StatusNotFound, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(baskets, nil))
}

// CreateUserBasket and return it
func (bH *basketEchoHandler) CreateUserBasket(e echo.Context) error {
	bearerToken, err := reqTokenHandler.ExtractBearerToken(e.Request())
	if err != nil {
		return e.JSON(http.StatusForbidden, generateResponse(nil, err))
	}

	ctx := context.Background()

	createdBasket, err := bH.uC.CreateBasket(ctx, bearerToken)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(createdBasket, nil))
}

// DisableActiveBasket which already exists in database
func (bH *basketEchoHandler) DisableActiveBasket(e echo.Context) error {
	bearerToken, err := reqTokenHandler.ExtractBearerToken(e.Request())
	if err != nil {
		return e.JSON(http.StatusForbidden, generateResponse(nil, err))
	}

	ctx := context.Background()

	err = bH.uC.DisableActiveBasket(ctx, bearerToken)
	if err != nil {
		return e.JSON(http.StatusNotFound, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(nil, basketDisablationMessage))
}

// AddProductsToBasket and return them
func (bH *basketEchoHandler) AddProductsToBasket(e echo.Context) error {
	addRequest := new(basket.AddProductsToBasketRequest)

	if err := e.Bind(addRequest); err != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, err))
	}

	if errs := validation.Validate(addRequest); errs != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, errs))
	}

	bearerToken, err := reqTokenHandler.ExtractBearerToken(e.Request())
	if err != nil {
		return e.JSON(http.StatusForbidden, generateResponse(nil, err))
	}

	ctx := context.Background()

	activeBasket, err := bH.uC.AddProductsToBasket(ctx, bearerToken, addRequest)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(activeBasket, nil))
}

// UpdateBasketProducts and return newly updated basket
func (bH *basketEchoHandler) UpdateBasketProducts(e echo.Context) error {
	updateRequest := new(basket.EditProductsToBasketRequest)

	if err := e.Bind(updateRequest); err != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, err))
	}

	if errs := validation.Validate(updateRequest); errs != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, errs))
	}

	bearerToken, err := reqTokenHandler.ExtractBearerToken(e.Request())
	if err != nil {
		return e.JSON(http.StatusForbidden, generateResponse(nil, err))
	}

	ctx := context.Background()

	updatedBasket, err := bH.uC.UpdateBasketProductsAmount(ctx, bearerToken, updateRequest)
	if err != nil {
		return e.JSON(http.StatusNotFound, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(updatedBasket, nil))
}
