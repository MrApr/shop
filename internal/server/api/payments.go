package api

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"shop/internal/domains/payment"
	"shop/internal/middleware/auth"
	"shop/pkg/validation"
	"strconv"
)

// PaymentHandler is a struct which attaches payment use case with it's desired restfull API
type PaymentHandler struct {
	uC payment.PaymentUseCaseContract
}

var (
	invalidPaymentUrl      error = errors.New("RequestedPaymentUrl is invalid")
	paymentIdConversionErr error = errors.New("internal server error, please contact administration")
)

// AttachPaymentWithRestUrls attaches payment operations with api service
func AttachPaymentWithRestUrls(echoEngine *echo.Echo, db *gorm.DB) {
	paymentRepo := payment.NewPaymentRepo(db)
	paymentUc := payment.NewPaymentUseCase(payment.NewPaymentStorageService(paymentRepo, nil, nil, nil, nil), db, nil)

	paymentHandlr := &PaymentHandler{
		uC: paymentUc,
	}

	setupPaymentDomains(echoEngine, paymentHandlr)
}

// setupPaymentDomains and attaches them to desired urls and function methods
func setupPaymentDomains(echoEngine *echo.Echo, paymentHandlr *PaymentHandler) {
	group := echoEngine.Group("/payments")

	group.Use(auth.ValidateJWT)

	group.GET("/:id", paymentHandlr.Pay)
}

// Pay operates created payment for user
func (pH *PaymentHandler) Pay(e echo.Context) error {
	paymentId := e.Param("id")
	if paymentId == "" {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, invalidPaymentUrl))
	}

	paymentIdInt, err := strconv.Atoi(paymentId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, paymentIdConversionErr))
	}

	bearerToken := e.Request().Header.Get("Authorization")
	ctx := context.Background()

	requestResult, err := pH.uC.Pay(ctx, bearerToken, paymentIdInt)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}
	return e.Redirect(http.StatusMovedPermanently, requestResult.Url) //Todo store payment results in log collector
}

// Verify user payment that is done
func (pH *PaymentHandler) Verify(e echo.Context) error {
	paymentRequest := new(payment.PaymentVerifyRequest)

	if err := e.Bind(paymentRequest); err != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, err))
	}

	if errs := validation.Validate(paymentRequest); errs != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, errs))
	}

	authorityCode := e.QueryParam("Authority")
	if authorityCode == "" {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, invalidPaymentUrl))
	}

	bearerToken := e.Request().Header.Get("Authorization")
	ctx := context.Background()

	verifiedPayment, err := pH.uC.Verify(ctx, bearerToken, paymentRequest)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}

	if verifiedPayment.Status != payment.PaymentSuccessStatus {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, paymentIdConversionErr))
	}

	return e.JSON(http.StatusOK, generateResponse(verifiedPayment, nil))
}
