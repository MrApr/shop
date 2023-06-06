package discount

import "errors"

var (
	NoDiscountsExists error = errors.New("no discounts exists in data storage")
)
