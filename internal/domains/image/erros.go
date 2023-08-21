package image

import "errors"

var (
	FileDoesntExists error = errors.New("Requested file does not exists")
)
