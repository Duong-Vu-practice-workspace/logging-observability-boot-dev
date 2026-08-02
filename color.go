package main

import (
	"context"
	"fmt"
	"log/slog"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorGreen  = "\033[32m"
)

// colorize wraps s in the ANSI color code for the given log level, making
// levels and messages easier to spot in a terminal during development.
func colorize(level slog.Level, s string) string {
	switch level {
	case slog.LevelDebug:
		return fmt.Sprintf("%s%s%s", colorBlue, s, colorReset)
	case slog.LevelInfo:
		return fmt.Sprintf("%s%s%s", colorGreen, s, colorReset)
	case slog.LevelWarn:
		return fmt.Sprintf("%s%s%s", colorYellow, s, colorReset)
	case slog.LevelError:
		return fmt.Sprintf("%s%s%s", colorRed, s, colorReset)
	default:
		return s
	}
}

// colorHandler wraps a slog.Handler and colorizes the log message before
// delegating, producing development-friendly console output.
type colorHandler struct {
	slog.Handler
}

func (h *colorHandler) Handle(ctx context.Context, r slog.Record) error {
	r.Message = colorize(r.Level, r.Message)
	return h.Handler.Handle(ctx, r)
}
