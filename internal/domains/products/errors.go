package products

import "errors"

var (
	NoCategoriesFound error = errors.New("no categories found in data source")
)
