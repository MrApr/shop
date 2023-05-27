package products

import "errors"

var (
	NoCategoriesFound error = errors.New("no categories found in data source")
	NoTypesFound      error = errors.New("no types found with provided credentials")
	NoProductsFound   error = errors.New("no products found in database")
	ProductNotFound   error = errors.New("required product didnt found in database")
)
