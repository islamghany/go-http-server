package middleware

import (
	"fmt"
	"httpserver/internal/validator"
	"httpserver/internal/web"
	"log"
	"net/http"
)

func Error() web.Middleware {
	return func(handler web.Handler) web.Handler {
		h := func(w http.ResponseWriter, r *http.Request) error {
			if err := handler(w, r); err != nil {
				log.Println(err)
				var er web.ErrorDocument
				var status int
				switch {
				// check if he error is of type Error
				case web.IsError(err):
					respErr := web.GetError(err)
					// check if the error is a validation error
					if validator.IsFieldErrors(respErr.Err) {
						fieldsErros := validator.GetFieldsErrors(respErr.Err)
						er = web.ErrorDocument{
							Error:  "validation error",
							Fields: fieldsErros.Fields(),
						}
						status = http.StatusBadRequest
						break
					}
					er = web.ErrorDocument{
						Error: respErr.Error(),
					}
					status = respErr.Status
					break
				// if the error is not of type Error then it is an internal server error
				default:
					er = web.ErrorDocument{
						Error: http.StatusText(http.StatusInternalServerError),
					}
					status = http.StatusInternalServerError
				}
				if err = web.Response(w, r, status, er); err != nil {
					return fmt.Errorf("sending error: %w", err)
				}
			}
			return nil
		}
		return h
	}

}
