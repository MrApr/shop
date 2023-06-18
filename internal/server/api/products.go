package api

import (
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"shop/internal/domains/products"
	"shop/pkg/validation"
	"strconv"
)

// productEchoHandler is the type which attaches rest api end points to domain functions
type productEchoHandler struct {
	uC products.ProductUseCaseInterface
}

// AttachProductHandlerToProductDomain for working with rest Apis
func AttachProductHandlerToProductDomain(engine *echo.Echo, db *gorm.DB) {
	uC := products.NewProductUseCase(products.NewProductsService(products.NewProductRepository(db)))
	setupProductRoutes(engine, &productEchoHandler{
		uC: uC,
	})
}

// setupProductRoutes which are accessible through http URI
func setupProductRoutes(engine *echo.Echo, handler *productEchoHandler) {
	router := engine.Group("products")
	router.GET("", handler.GetAllProducts)
}

// GetAllProducts and return them
func (pH *productEchoHandler) GetAllProducts(e echo.Context) error {
	request := new(products.GetAllProductsRequest)

	if err := e.Bind(request); err != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, err))
	}

	if errs := validation.Validate(request); errs != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, errs))
	}

	ctx := context.Background()

	productsList, err := pH.uC.GetAllProducts(ctx, request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(productsList, nil))
}

// GetProduct and return it
func (pH *productEchoHandler) GetProduct(e echo.Context) error {
	productId := e.Param("id")
	productIdInt, err := strconv.Atoi(productId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, wrongProductId))
	}

	ctx := context.Background()

	singleProduct, err := pH.uC.GetProduct(ctx, productIdInt)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(singleProduct, nil))
}
