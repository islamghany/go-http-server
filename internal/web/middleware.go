package web

type Middleware func(handler Handler) Handler

func wrapMiddleware(handler Handler, mw []Middleware) Handler {
	for i := len(mw) - 1; i >= 0; i-- {
		mwFunc := mw[i]
		if mwFunc != nil {
			handler = mwFunc(handler)
		}
	}
	return handler
}
