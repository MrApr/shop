package gateways

import "errors"

var (
	NoTypesFound error = errors.New("no types found in database")
)
