package discount

import "errors"

var (
	NoDiscountsExists error = errors.New("no discounts exists in data storage")
	DiscountNotFound  error = errors.New("required discount does not exists")
)
