package basket

import "errors"

var (
	ProductIsFinished         error = errors.New("stack amount for requested product is empty")
	InsufficientProductAmount error = errors.New("product doesn't have enough and sufficient inventory")
	NoActiveBasket            error = errors.New("user doesnt have any active basket")
)
