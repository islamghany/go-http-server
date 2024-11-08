package middleware

import (
	"fmt"
	"httpserver/internal/web"
	"httpserver/pkg/logger"
	"net/http"
	"time"
)

func Logger(logger *logger.Logger) web.Middleware {
	return func(handler web.Handler) web.Handler {
		h := func(w http.ResponseWriter, r *http.Request) error {
			val := web.GetValues(r.Context())
			path := r.URL.Path
			if r.URL.RawQuery != "" {
				path = fmt.Sprintf("%s?%s", path, r.URL.RawQuery)
			}

			logger.Info(r.Context(), "request started", "method", r.Method, "path", path, "remote_addr", r.RemoteAddr)
			err := handler(w, r)

			logger.Info(r.Context(), "request completed", "method", r.Method, "path", path, "remote_addr", r.RemoteAddr, "status", val.SetStatusCode, "since",
				fmt.Sprintf("%v ms", time.Since(val.Time).Milliseconds()))
			return err

		}
		return h
	}

}
