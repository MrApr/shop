package comment

import "errors"

var (
	NoCommentsFound     error = errors.New("no products found in db")
	CommentNotFound     error = errors.New("requested comment not found")
	OperationNotAllowed error = errors.New("you are not allowed to do the operation")
)
