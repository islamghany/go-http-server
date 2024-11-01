package web

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

type WebApp struct {
	*http.ServeMux
	mw []Middleware
}

func NewApp(mw ...Middleware) *WebApp {
	return &WebApp{
		ServeMux: http.NewServeMux(),
		mw:       mw,
	}
}

func (a *WebApp) Handle(method, group, path string, handler Handler, mw ...Middleware) {
	handler = wrapMiddleware(handler, mw)
	handler = wrapMiddleware(handler, a.mw)
	a.handle(method, group, path, handler)
}

func (a *WebApp) handle(method, group, path string, handler Handler) {
	h := func(w http.ResponseWriter, r *http.Request) {
		ctx := SetValues(r.Context(), &Values{
			TracerID: uuid.New().String(),
			Time:     time.Now().UTC(),
		})
		if err := handler(w, r.WithContext(ctx)); err != nil {
			log.Println(err)
		}
	}

	pattern := a.parsePathForMux(method, group, path)
	a.ServeMux.HandleFunc(pattern, h)

}

func (a *WebApp) parsePathForMux(method, group, path string) string {
	// convert /api/v1/users/:id to  /api/v1/users/{id}
	// convert /api/v1/users/:id/:name to  /api/v1/users/{id}/{name}

	segments := strings.Split(path, "/")
	p := ""
	for i, segment := range segments {
		if segment != "" {
			if segment[0] == ':' {
				segment = "{" + segment[1:] + "}"
			}
			if i == 0 {
				p = segment
			} else {
				p = p + "/" + segment
			}
		}
	}
	if group != "" {
		p = "/" + group + p
	}
	return method + " " + p
}
