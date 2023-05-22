package address

import "errors"

var (
	NoCitiesFound error = errors.New("no cities found in database")
)
