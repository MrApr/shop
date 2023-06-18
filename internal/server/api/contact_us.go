package api

import (
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"shop/internal/domains/contact_us"
	"shop/pkg/validation"
)

// contactUsEchoHandler is the type which attaches rest api end points to domain functions
type contactUsEchoHandler struct {
	uC contact_us.ContactUsUseCaseInterface
}

// AttachContactUsHandlerToContactUsDomain for working with rest Apis
func AttachContactUsHandlerToContactUsDomain(engine *echo.Echo, db *gorm.DB) {
	useCase := contact_us.NewUseCase(contact_us.NewContactUsService(contact_us.NewContactUsRepository(db)))
	setupContactUsRoutes(engine, &contactUsEchoHandler{
		uC: useCase,
	})
}

// setupContactUsRoutes which are accessible through http URI
func setupContactUsRoutes(engine *echo.Echo, handler *contactUsEchoHandler) {
	contactRoutes := engine.Group("/contact_us")
	contactRoutes.POST("", handler.SubmitContactUs)
}

// SubmitContactUs and store it in db
func (cH *contactUsEchoHandler) SubmitContactUs(e echo.Context) error {
	newContactRequest := new(contact_us.CreateContactUsRequest)

	if err := e.Bind(newContactRequest); err != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, err))
	}

	if errs := validation.Validate(newContactRequest); errs != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, errs))
	}

	ctx := context.Background()

	contactUs, err := cH.uC.CreateContactUs(ctx, newContactRequest)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(contactUs, nil))
}
