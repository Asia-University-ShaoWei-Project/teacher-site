package message

import "errors"

var (
	ErrUnnecessaryUpdate  = errors.New("the data is up to date")
	ErrExistUserId        = errors.New("the user is already existed")
	ErrExistTeacherDomain = errors.New("the teacher domain is already existed")
	// cache
	ErrMaximumRetry = errors.New("reached maximum number of retries")
)
