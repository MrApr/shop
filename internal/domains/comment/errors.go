package comment

import "errors"

var (
	NoProductsFound error = errors.New("no products found in db")
	CommentNotFound error = errors.New("requested comment not found")
)
