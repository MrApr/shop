package api

import (
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"shop/internal/domains/postType"
)

// postTypeEchoHandler is the type which attaches rest api end points to domain functions
type postTypeEchoHandler struct {
	uC postType.PostTypeUseCaseInterface
}

// AttachPostTypeHandlerToPostTypeDomain for working with rest Apis
func AttachPostTypeHandlerToPostTypeDomain(engine *echo.Echo, db *gorm.DB) {
	uC := postType.NewUseCase(postType.NewPostTypeService(postType.NewRepository(db)))

	setupPostTypeHandler(engine, &postTypeEchoHandler{
		uC: uC,
	})
}

// setupPostTypeHandler which are accessible through http URI
func setupPostTypeHandler(engine *echo.Echo, handler *postTypeEchoHandler) {
	postTypeRouter := engine.Group("/post_types")
	postTypeRouter.GET("", handler.GetAllPostTypes)
}

// GetAllPostTypes and return them
func (pH *postTypeEchoHandler) GetAllPostTypes(e echo.Context) error {
	ctx := context.Background()

	postTypes, err := pH.uC.GetAllPostTypes(ctx)
	if err != nil {
		return e.JSON(http.StatusNotFound, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(postTypes, nil))
}
