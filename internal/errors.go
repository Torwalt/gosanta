package awards

import (
	"errors"
	"fmt"
)

type ErrorType int

const (
	DuplicateError ErrorType = iota
	DoesNotExistError
	NoAward
	NotPhishingAction
	Unknown
)

type Error struct {
	Code ErrorType
	Err  error
}

func (e *Error) Error() string {
	return e.Err.Error()
}

func ExtendError(err error, msg string) *Error {
	var awardErr *Error

	if errors.As(err, &awardErr) != true {
		return &Error{Code: Unknown, Err: fmt.Errorf("%v: %v", msg, err)}
	}

	return &Error{Code: awardErr.Code, Err: fmt.Errorf("%v: %v", msg, err)}
}
