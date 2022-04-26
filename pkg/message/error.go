package message

import "errors"

var (
	ErrUnnecessaryUpdate = errors.New("the data is up to date")
	// cache
	ErrMaximumRetry = errors.New("reached maximum number of retries")
)
