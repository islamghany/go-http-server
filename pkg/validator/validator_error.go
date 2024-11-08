package validator

import (
	"encoding/json"
	"errors"
)

// FieldError is used to indicate an error with a specific request field.
type FieldError struct {
	Field string `json:"field"`
	Err   string `json:"error"`
}

type FieldErrors []FieldError

func NewFieldError(field, err string) error {
	return FieldErrors{{Field: field, Err: err}}
}

func (f FieldErrors) Error() string {
	d, err := json.Marshal(f)
	if err != nil {
		return err.Error()
	}
	return string(d)
}

func (f FieldErrors) Fields() map[string]string {
	m := map[string]string{}
	for _, v := range f {
		m[v.Field] = v.Err
	}
	return m
}
func IsFieldErrors(err error) bool {
	var fErr FieldErrors
	return errors.As(err, &fErr)
}

func GetFieldsErrors(err error) FieldErrors {
	var fErr FieldErrors
	if errors.As(err, &fErr) {
		return fErr
	}
	return nil
}
