package comment

import "errors"

var (
	NoCommentsFound error = errors.New("no products found in db")
	CommentNotFound error = errors.New("requested comment not found")
)
