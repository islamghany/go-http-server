package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"path/filepath"
	"runtime"
	"time"
)

// Different types of level for logging.
type Level slog.Level

const (
	LevelDebug = Level(slog.LevelDebug)
	LevelInfo  = Level(slog.LevelInfo)
	LevelWarn  = Level(slog.LevelWarn)
	LevelError = Level(slog.LevelError)
)

// EventFunc the the func type that will be invoked when the current event type happens.
type EventFunc func(ctx context.Context, r Record)

type Events struct {
	Debug EventFunc
	Info  EventFunc
	Warn  EventFunc
	Error EventFunc
}

// Record represent the structs that will be logged.
type Record struct {
	Time    time.Time
	Level   Level
	Message string
	Attrs   map[string]interface{}
}

// toRecord convert the slog.Record to logger.Record.
func toRecord(r slog.Record) Record {

	attrs := make(map[string]interface{}, r.NumAttrs())
	f := func(attr slog.Attr) bool {
		attrs[attr.Key] = attr.Value.Any()
		return true
	}

	r.Attrs(f)

	return Record{
		Time:    r.Time,
		Level:   Level(r.Level),
		Message: r.Message,
		Attrs:   attrs,
	}
}

// TraceIDFunc is the function that will be used to get the trace id from the context.
type TraceIDFunc func(ctx context.Context) string

// Logger is the logger struct that will be used to log the events.
type Logger struct {
	traceIDFunc TraceIDFunc
	handler     slog.Handler
}

func New(w io.Writer, minLevel Level, serviceName string, traceIDFunc TraceIDFunc) *Logger {
	return new(w, minLevel, serviceName, traceIDFunc, Events{})
}

func NewWithEvents(w io.Writer, minLevel Level, serviceName string, traceIDFunc TraceIDFunc, events Events) *Logger {
	return new(w, minLevel, serviceName, traceIDFunc, events)
}

func new(w io.Writer, minLevel Level, serviceName string, traceIDFunc TraceIDFunc, events Events) *Logger {
	// Convert the file name to just the name.ext when logging.
	f := func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.SourceKey {
			if source, ok := a.Value.Any().(*slog.Source); ok {
				v := fmt.Sprintf("%s:%d", filepath.Base(source.File), source.Line)
				return slog.Attr{Key: "file", Value: slog.StringValue(v)}
			}
		}
		return a
	}

	handler := slog.Handler(
		slog.NewJSONHandler(w, &slog.HandlerOptions{
			Level:       slog.Level(minLevel),
			AddSource:   true,
			ReplaceAttr: f,
		}),
	)

	// If any of the events are not nil, then we will create a new log handler.
	if events.Debug != nil || events.Info != nil || events.Warn != nil || events.Error != nil {
		handler = newLogHandler(handler, events)
	}

	// Inject the service name into the handler.
	attrs := []slog.Attr{
		{Key: "service", Value: slog.StringValue(serviceName)},
	}

	handler = handler.WithAttrs(attrs)

	return &Logger{
		traceIDFunc: traceIDFunc,
		handler:     handler,
	}
}

func (l *Logger) write(ctx context.Context, level Level, caller int, msg string, args ...any) {
	slogLevel := slog.Level(level)
	if !l.handler.Enabled(ctx, slogLevel) {
		return
	}

	var pcs [1]uintptr
	runtime.Callers(caller, pcs[:])

	r := slog.NewRecord(time.Now(), slogLevel, msg, pcs[0])
	if l.traceIDFunc != nil {
		args = append(args, slog.String("trace_id", l.traceIDFunc(ctx)))
	}

	r.Add(args...)
	l.handler.Handle(ctx, r)
}

// Debug logs at LevelDebug with the given context.
func (log *Logger) Debug(ctx context.Context, msg string, args ...any) {
	log.write(ctx, LevelDebug, 3, msg, args...)
}

// Info logs at LevelInfo with the given context.
func (log *Logger) Info(ctx context.Context, msg string, args ...any) {
	log.write(ctx, LevelInfo, 3, msg, args...)
}

// Infoc logs the information at the specified call stack position.
func (log *Logger) Infoc(ctx context.Context, caller int, msg string, args ...any) {
	log.write(ctx, LevelInfo, caller, msg, args...)
}

// Error logs at LevelError with the given context.
func (log *Logger) Error(ctx context.Context, msg string, args ...any) {
	log.write(ctx, LevelError, 3, msg, args...)
}

// Warn logs at LevelWarn with the given context.
func (log *Logger) Warn(ctx context.Context, msg string, args ...any) {
	log.write(ctx, LevelWarn, 3, msg, args...)
}
