package web

import "errors"

// ErrorDocument the error which returned to the client when there is an error
type ErrorDocument struct {
	Error  string            `json:"error"`
	Fields map[string]string `json:"fields,omitempty"`
}

// Error is used to passed the error in the app
type Error struct {
	Err    error
	Status int
}

func NewError(err error, status int) error {
	return &Error{
		err,
		status,
	}
}

func (err *Error) Error() string {
	return err.Err.Error()
}

func IsError(err error) bool {
	var er *Error
	return errors.As(err, &er)
}

func GetError(err error) *Error {
	var er *Error
	if errors.As(err, &er) {
		return er
	}
	return nil
}
