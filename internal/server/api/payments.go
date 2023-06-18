package api

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"shop/internal/domains/payment"
	"shop/internal/middleware/auth"
	"shop/pkg/reqTokenHandler"
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
	group.GET("", paymentHandlr.GetUserPayments)
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

	bearerToken, err := reqTokenHandler.ExtractBearerToken(e.Request())
	if err != nil {
		return e.JSON(http.StatusForbidden, generateResponse(nil, err))
	}
	ctx := context.Background()

	requestResult, err := pH.uC.Pay(ctx, bearerToken, paymentIdInt)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, generateResponse(nil, err))
	}
	return e.Redirect(http.StatusMovedPermanently, requestResult.Url) //Todo store payment results in log collector
}

// GetUserPayments and return them
func (pH *PaymentHandler) GetUserPayments(e echo.Context) error {
	request := new(payment.GetUserPaymentsRequest)

	if err := e.Bind(request); err != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, err))
	}

	if errs := validation.Validate(request); errs != nil {
		return e.JSON(http.StatusBadRequest, generateResponse(nil, errs))
	}

	bearerToken, err := reqTokenHandler.ExtractBearerToken(e.Request())
	if err != nil {
		return e.JSON(http.StatusForbidden, generateResponse(nil, err))
	}

	ctx := context.Background()

	payments, err := pH.uC.GetUserPayments(ctx, bearerToken, request)
	if err != nil {
		return e.JSON(http.StatusNotFound, generateResponse(nil, err))
	}

	return e.JSON(http.StatusOK, generateResponse(payments, nil))
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

	bearerToken, err := reqTokenHandler.ExtractBearerToken(e.Request())
	if err != nil {
		return e.JSON(http.StatusForbidden, generateResponse(nil, err))
	}

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
