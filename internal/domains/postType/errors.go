package postType

import "errors"

var (
	NotTypesFound error = errors.New("no post types found in db")
)
