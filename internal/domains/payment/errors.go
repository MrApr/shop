package payment

import "errors"

var (
	PaymentNotFound    error = errors.New("no payment found with given credentials")
	Unauthorized       error = errors.New("user is not allowed to do the operation")
	SomethingWentWrong error = errors.New("something went wrong\n Please contact administrator")
)
