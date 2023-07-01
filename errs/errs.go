package errs

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

func NewNotFoundError(msg string) error {
	return AppError{
		Code:    http.StatusNotFound,
		Message: msg,
	}
}

func NewUnExpectedError() error {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: "unexpected error",
	}
}
