package gateways

import "errors"

var (
	NoTypesFound    error = errors.New("no types found in database")
	NoGatewaysFound error = errors.New("no gateways found with requested credentials in db")
)
