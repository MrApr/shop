package payment

import "errors"

var (
	PaymentNotFound error = errors.New("no payment found with given credentials")
)
