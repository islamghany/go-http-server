package web

import (
	"encoding/json"
	"net/http"
)

// Response sends a JSON response with the given status code and data.
func Response[T any](w http.ResponseWriter, r *http.Request, statusCode int, data T) error {
	// set the statusCode
	SetStatusCode(r.Context(), statusCode)
	w.WriteHeader(statusCode)
	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}
