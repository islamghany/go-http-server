package web

import (
	"context"
	"time"
)

type Values struct {
	TracerID      string
	Time          time.Time
	SetStatusCode int
}

type valuesKey int

const key valuesKey = 1

func SetValues(ctx context.Context, values *Values) context.Context {
	return context.WithValue(ctx, key, values)
}
func GetValues(ctx context.Context) *Values {
	val, ok := ctx.Value(key).(*Values)
	if !ok {
		return &Values{
			TracerID: "0000-0000-0000-0000",
			Time:     time.Now(),
		}
	}
	return val
}
func GetTracerID(ctx context.Context) string {
	val, ok := ctx.Value(key).(*Values)
	if !ok {
		return "0000-0000-0000-0000"
	}
	return val.TracerID
}

func GetTime(ctx context.Context) time.Time {
	val, ok := ctx.Value(key).(*Values)
	if !ok {
		return time.Now()
	}
	return val.Time
}

func GetSetStatusCode(ctx context.Context) int {
	val, ok := ctx.Value(key).(*Values)
	if !ok {
		return 0
	}
	return val.SetStatusCode
}

func SetStatusCode(ctx context.Context, status int) {
	val, ok := ctx.Value(key).(*Values)
	if !ok {
		return
	}
	val.SetStatusCode = status
}
