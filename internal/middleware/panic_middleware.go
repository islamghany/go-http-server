package middleware

import (
	"fmt"
	"httpserver/internal/web"
	"net/http"
	"runtime/debug"
)

func Panic() web.Middleware {
	return func(handler web.Handler) web.Handler {
		return func(w http.ResponseWriter, r *http.Request) (err error) {
			defer func() {
				if recErr := recover(); recErr != nil {
					trace := debug.Stack()
					err = fmt.Errorf("panic [%v] trace[%s]", recErr, trace)
					// fmt.Println("err", err)
				}
			}()
			return handler(w, r)
		}
	}
}
