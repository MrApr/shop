package api

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"shop/internal/domains/comment"
	"shop/internal/middleware/auth"
	"shop/pkg/reqTokenHandler"
	"shop/pkg/validation"
	"strconv"
)

var (
	wrongProductId             error  = errors.New("provided product id is not valid")
	commentDeletedSuccessFully string = "requested comment deleted successfully"
)

// CommentEchoHandler is the type which attaches rest api end points to domain functions
type CommentEchoHandler struct {
	uC comment.CommentUseCaseInterface
}

// AttachCommentHandlerWithCommentDomain for working with rest Apis
func AttachCommentHandlerWithCommentDomain(engine *echo.Echo, db *gorm.DB) {
	uC := comment.NewCommentUseCase(comment.NewCommentService(comment.NewCommentRepository(db)), nil)
	setupCommentHandlerRoutes(engine, &CommentEchoHandler{
		uC: uC,
	})
}

// setupCommentHandlerRoutes which are accessible through http URI
func setupCommentHandlerRoutes(engine *echo.Echo, cH *CommentEchoHandler) {
	commentRoutes := engine.Group("/comments")
	commentRoutes.GET("/products/:id", cH.GetAllProductComments)
	commentRoutes.Use(auth.ValidateJWT)
	commentRoutes.POST("", cH.CreateProductComment)
	commentRoutes.DELETE("/:id", cH.DeleteComment)
}

// GetAllProductComments and return them
func (cH *CommentEchoHandler) GetAllProductComments(e echo.Context) error {
	productId := e.Param("id")
	productIdInt, err := strconv.Atoi(productId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, wrongProductId))
	}

	ctx := context.Background()

	comments, err := cH.uC.GetAllActiveComments(ctx, productIdInt)
	if err != nil {
		return e.JSON(http.StatusNotFound, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(comments, nil))
}

// CreateProductComment and insert it in db
func (cH *CommentEchoHandler) CreateProductComment(e echo.Context) error {
	createRequest := new(comment.CreateCommentRequest)

	if err := e.Bind(createRequest); err != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, err))
	}

	if errs := validation.Validate(createRequest); errs != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, errs))
	}

	bearerToken, err := reqTokenHandler.ExtractBearerToken(e.Request())
	if err != nil {
		return e.JSON(http.StatusForbidden, generateResponse(nil, err))
	}

	ctx := context.Background()

	createdCm, err := cH.uC.CreateComment(ctx, bearerToken, createRequest)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(createdCm, nil))
}

// DeleteComment which already exists in database
func (cH *CommentEchoHandler) DeleteComment(e echo.Context) error {
	cmId := e.Param("id")
	cmIdInt, err := strconv.Atoi(cmId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, wrongProductId))
	}

	bearerToken, err := reqTokenHandler.ExtractBearerToken(e.Request())
	if err != nil {
		return e.JSON(http.StatusForbidden, generateResponse(nil, err))
	}

	ctx := context.Background()

	err = cH.uC.DeleteComment(ctx, bearerToken, cmIdInt)
	if err != nil {
		return e.JSON(http.StatusNotFound, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(nil, commentDeletedSuccessFully))
}
