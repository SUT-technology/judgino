	package slogger

	import (
		"context"
		"io"
		"log/slog"
	)

	func Err(msg string, err error) slog.Attr {
		if err == nil {
			return slog.String(msg, "no-error")
		}
		return slog.String(msg, err.Error())
	}

	// NewJSONLogger creates a new slog.Logger instance which logs to the stdout. The given level must be: "debug", "info",
	// "warn" or "error". It will default to Info level if input is invalid.
	func NewJSONLogger(level string, w io.Writer) *slog.Logger {
		h := &ContextHandler{slog.NewJSONHandler(
			w, &slog.HandlerOptions{
				Level: getLevel(level),
				ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
					switch a.Key {
					case slog.LevelKey:
						lvl := a.Value.Any().(slog.Level)
						label, exist := levelNames[lvl]
						if !exist {
							label = lvl.String()
						}
						a.Value = slog.StringValue(label)
					}

					return a
				},
			}),
		}
		return slog.New(h)
	}

	func getLevel(lvl string) slog.Level {
		switch lvl {
		case "trace":
			return LevelTrace
		case "debug":
			return slog.LevelDebug
		case "info":
			return slog.LevelInfo
		case "warn":
			return slog.LevelWarn
		case "error":
			return slog.LevelError
		case "fatal":
			return LevelFatal
		default:
			return slog.LevelInfo
		}
	}

	type slogAttr string

	const slogAttrs slogAttr = "slog_attrs"

	// ContextHandler embeds slog.Handler, overriding Handle method to log context attributes.
	type ContextHandler struct {
		slog.Handler
	}

	// Handle adds contextual attributes to the Record before calling the underlying handler.
	func (h ContextHandler) Handle(ctx context.Context, r slog.Record) error {
		if attrs, ok := ctx.Value(slogAttrs).([]slog.Attr); ok {
			for _, v := range attrs {
				r.AddAttrs(v)
			}
		}
		return h.Handler.Handle(ctx, r)
	}

	// WithAttrs adds one or more slog attributes to the provided context, so that they will be included in any Log Records
	// created with such context. It relies on the caller to not pass a nil context.
	func WithAttrs(parent context.Context, attr ...slog.Attr) context.Context {
		if len(attr) == 0 {
			return parent
		}

		// if some slog attributes already exist, append to them
		if v, ok := parent.Value(slogAttrs).([]slog.Attr); ok {
			v = append(v, attr...)
			return context.WithValue(parent, slogAttrs, v)
		}

		var v []slog.Attr
		v = append(v, attr...)
		return context.WithValue(parent, slogAttrs, v)
	}
