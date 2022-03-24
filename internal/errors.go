package awards

type ErrorType int

const (
	DuplicateError ErrorType = iota
	DoesNotExistError
)

type Error struct {
	Code ErrorType
	Err  error
}

func (e *Error) Error() string {
	return e.Err.Error()
}
