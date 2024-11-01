package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"httpserver/internal/validator"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

const (
	MAX_BODY_SIZE = int64(1024 * 1024 * 2) // 2MB
)

// Decode decodes the request body into dst and validates it.
func Decode[T any](w http.ResponseWriter, r *http.Request, dst T) error {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_BODY_SIZE)
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(dst); err != nil {
		return fmt.Errorf("decoding request: %w", err)
	}
	// Validate the request body
	return validator.Check(dst)
}

func GetParamNumber(r *http.Request, key string) (int, error) {
	param := r.PathValue(key)
	if param == "" {
		return 0, fmt.Errorf("missing parameter %s", key)
	}
	return strconv.Atoi(param)
}

func GetParamUUID(r *http.Request, key string) (uuid.UUID, error) {
	param := r.PathValue(key)
	if param == "" {
		return uuid.UUID{}, fmt.Errorf("missing parameter %s", key)
	}
	uid, err := uuid.Parse(param)
	if err != nil {
		return uuid.UUID{}, errors.New("url parameter id is not a valid uuid")
	}
	return uid, nil
}

func GetParamString(r *http.Request, key string) (string, error) {
	param := r.PathValue(key)
	if param == "" {
		return "", fmt.Errorf("missing parameter %s", key)
	}
	return param, nil
}
