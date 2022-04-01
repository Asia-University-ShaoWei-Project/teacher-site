package message

import "errors"

var (
	ErrQueryNotFound = errors.New("not found the info data")
	ErrDataEmpty     = errors.New("the data is empty")
	ErrAuthInvalid   = errors.New("authorization invalided")
)
