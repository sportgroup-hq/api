package logger

import (
	"context"
	"io"
	"log/slog"
)

type Logger struct {
	*slog.TextHandler
}

func NewLogger(w io.Writer, l string) *Logger {
	return &Logger{TextHandler: slog.NewTextHandler(w, &slog.HandlerOptions{
		Level: logLevel(l),
	})}
}

func (l Logger) Handle(ctx context.Context, record slog.Record) error {
	//stringFields := []string{
	//}
	//
	//jsonFields := []string{
	//}
	//
	//for i := range stringFields {
	//	if v := ctx.Value(stringFields[i]); v != nil {
	//		record.Add(stringFields[i], v)
	//	}
	//}
	//
	//for i := range jsonFields {
	//	if v := ctx.Value(jsonFields[i]); v != nil {
	//		j, _ := json.Marshal(v)
	//		record.Add(jsonFields[i], j)
	//	}
	//}

	return l.TextHandler.Handle(ctx, record)
}

func logLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		slog.Warn("unknown log level", "level", level)
		return slog.LevelInfo
	}
}
