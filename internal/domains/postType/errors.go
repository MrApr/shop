package postType

import "errors"

var (
	NotTypesFound        error = errors.New("no post types found in db")
	PostTypeDoesntExists error = errors.New("no post types with given id found in database")
)
