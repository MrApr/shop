package zarinpal

import (
	"errors"
	"fmt"
	"github.com/sinabakh/go-zarinpal-checkout"
	"log"
	"os"
	"shop/internal/domains/payment"
	"shop/pkg/advancedError"
	"shop/pkg/paymentHandler"
)

// defines env key which fetches zarinpal settings
const (
	gatewayName               string = "ZARINPAL"
	ZarinPalReturnURLEnvKey   string = "ZARINPAL_RETURN_URL"
	zarinPalSuccessStatusCode int    = 100
)

var PaymentUpdateFails error = fmt.Errorf("%s\n%s", "Payment update failed", "Please contact administrator")
var ENVAreNotSet error = errors.New("Environment variables are not set completely")

// ZarinpalPGW defines set of methods which implements payment gateway payment contract
type ZarinpalPGW struct {
	repo       payment.PaymentRepoContract
	key        string
	zarinPalPG *zarinpal.Zarinpal
}

// NewZarinpal creates and returns zarinpal implementation
func NewZarinpal(repo payment.PaymentRepoContract, isSandBox bool) payment.PaymentPGWServiceContract {
	merchantId := paymentHandler.GetGatewayToken(gatewayName)
	if merchantId == "" {
		log.Fatalf("%s", "Required enviroment variable for zarinpal is not set properly")
	}

	zp, err := zarinpal.NewZarinpal(merchantId, isSandBox)
	if err != nil {
		log.Fatalf("%s", "Zarinpal Handler creation failed!")
	}

	return &ZarinpalPGW{
		repo:       repo,
		zarinPalPG: zp,
	}
}

// Pay operates payment
func (zp *ZarinpalPGW) Pay(paymentId int) (*payment.RequestPaymentResponse, error) {
	requestedPayment, err := zp.repo.GetPayment(paymentId)
	if err != nil {
		return nil, err
	}

	returnUrl, err := zp.createReturnUrl(requestedPayment.Id)
	if err != nil {
		return nil, advancedError.New(err, "Cannot create requestedPayment request")
	}

	paymentAmountInt := int(requestedPayment.TotalPrice)
	url, auth, status, err := zp.zarinPalPG.NewPaymentRequest(paymentAmountInt, returnUrl, "New requestedPayment for our products", "", "")
	if err != nil || status != zarinPalSuccessStatusCode {
		return nil, advancedError.New(err, "Cannot create requestedPayment request")
	}

	requestedPayment, err = zp.repo.UpdatePaymentTrace(requestedPayment, auth)
	if err != nil {
		return nil, advancedError.New(err, "Cannot update payment authority")
	}

	return &payment.RequestPaymentResponse{
		Url:             url,
		Key:             auth,
		OperationStatus: status,
	}, nil
}

// Verify done payment by user
func (zp *ZarinpalPGW) Verify(paymentId int, Authority string) (*payment.Payment, error) {
	requestedPayment, err := zp.repo.GetPayment(paymentId)
	if err != nil {
		return nil, err
	}

	paymentAmountInt := int(requestedPayment.TotalPrice)
	isVerified, refId, statusCode, err := zp.zarinPalPG.PaymentVerification(paymentAmountInt, Authority)
	if !isVerified || statusCode != zarinPalSuccessStatusCode && err != nil {
		return nil, advancedError.New(err, "Cannot verify requestedPayment")
	}

	isVerifiedStr := zp.getPaymentStatusStr(isVerified)
	requestedPayment, err = zp.repo.UpdatePaymentRefStatus(requestedPayment, refId, isVerifiedStr)
	if err != nil {
		return nil, PaymentUpdateFails
	}

	return requestedPayment, nil
}

// createReturnUrl and return it in order to pass it to zarinpal
func (zp *ZarinpalPGW) createReturnUrl(paymentId int) (string, error) {
	returnUrl := os.Getenv(ZarinPalReturnURLEnvKey)
	if returnUrl == "" {
		return "", ENVAreNotSet
	}
	returnUrl = fmt.Sprintf("%s/?payment_id=%d", returnUrl, paymentId)
	return returnUrl, nil
}

// getPaymentStatusStr based on given bool
func (zp *ZarinpalPGW) getPaymentStatusStr(status bool) string {
	if status {
		return payment.PaymentSuccessStatus
	}

	return payment.PaymentFailureStatus
}
