package api

import (
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"shop/internal/domains/products"
	"shop/pkg/reqTokenHandler"
	"shop/pkg/validation"
)

// operationSuccess is a message for successful operations
const operationSuccess string = "Operation was successful"

// likeDislikeEchoHandler is the type which attaches rest api end points to domain functions
type likeDislikeEchoHandler struct {
	uC products.LikeDislikeUseCaseInterface
}

// AttachLikeDislikeToItsDomain for working with rest Apis
func AttachLikeDislikeToItsDomain(engine *echo.Echo, db *gorm.DB) {
	uC := products.NewLikeDislikeUseCase(products.NewLikeDislikeService(products.NewLikeDislikeRepository(db)), nil)
	setupLikeDislikeRoutes(engine, &likeDislikeEchoHandler{
		uC: uC,
	})
}

// setupLikeDislikeRoutes which are accessible through http URI
func setupLikeDislikeRoutes(engine *echo.Echo, handler *likeDislikeEchoHandler) {
	router := engine.Group("products")
	router.POST("/like", handler.Like)
	router.POST("/dislike", handler.Dislike)
}

// Like product for user
func (lH *likeDislikeEchoHandler) Like(e echo.Context) error {
	request := new(products.LikeDislikeRequest)

	if err := e.Bind(request); err != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, err))
	}

	if errs := validation.Validate(request); errs != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, errs))
	}

	ctx := context.Background()

	bearerToken, err := reqTokenHandler.ExtractBearerToken(e.Request())
	if err != nil {
		return e.JSON(http.StatusForbidden, generateResponse(nil, err))
	}

	err = lH.uC.LikeProduct(ctx, bearerToken, request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(nil, operationSuccess))
}

// Dislike product for user
func (lH *likeDislikeEchoHandler) Dislike(e echo.Context) error {
	request := new(products.LikeDislikeRequest)

	if err := e.Bind(request); err != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, err))
	}

	if errs := validation.Validate(request); errs != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, errs))
	}

	ctx := context.Background()

	bearerToken, err := reqTokenHandler.ExtractBearerToken(e.Request())
	if err != nil {
		return e.JSON(http.StatusForbidden, generateResponse(nil, err))
	}

	err = lH.uC.DislikeProduct(ctx, bearerToken, request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(nil, operationSuccess))
}
