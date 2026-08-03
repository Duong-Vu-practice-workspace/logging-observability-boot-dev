package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net/url"
	"os"
	"strings"
	"os/signal"
	"syscall"
	"time"

	"boot.dev/linko/internal/build"
	"boot.dev/linko/internal/store"
	pkgerr "github.com/pkg/errors"
	"gopkg.in/natefinch/lumberjack.v2"
)

type stackTracer interface {
	error
	StackTrace() pkgerr.StackTrace
}

type multiError interface {
	error
	Unwrap() []error
}

type closeFunc func() error

func errorAttrs(err error) []slog.Attr {
	var me multiError
	if errors.As(err, &me) {
		var attrs []slog.Attr
		for i, e := range me.Unwrap() {
			attrs = append(attrs, slog.GroupAttrs(fmt.Sprintf("error_%d", i+1), collectAttrs(e)...))
		}
		return attrs
	}
	return collectAttrs(err)
}

func collectAttrs(err error) []slog.Attr {
	attrs := []slog.Attr{
		{Key: "message", Value: slog.StringValue(err.Error())},
	}
	attrs = append(attrs, Attrs(err)...)
	var se stackTracer
	if errors.As(err, &se) {
		attrs = append(attrs, slog.Attr{Key: "stack_trace", Value: slog.StringValue(fmt.Sprintf("%+v", se.StackTrace()))})
	}
	return attrs
}


// sensitiveKeys are attribute keys whose values should never reach a log.
var sensitiveKeys = []string{"password", "key", "apikey", "secret", "pin", "creditcardno"}

// redactValue replaces a string value with [REDACTED] if it carries a
// suspected secret: a sensitive key name, or a URL that embeds credentials.
func redactValue(v string) string {
	if strings.Contains(v, "://") {
		if stripped := stripURLCredentials(v); stripped != v {
			return stripped
		}
	}
	if containsAny(strings.ToLower(v), sensitiveKeys) {
		return "[REDACTED]"
	}
	return v
}

func stripURLCredentials(raw string) string {
	u, err := url.Parse(raw)
	if err != nil {
		return raw
	}
	if u.Hostname() == "" || (u.User == nil || u.User.Username() == "") {
		return raw
	}
	u.User = url.User("xxxxx")
	return u.String()
}

func containsAny(haystack string, needles []string) bool {
	for _, n := range needles {
		if strings.Contains(haystack, n) {
			return true
		}
	}
	return false
}

// sensitiveFilter is the last-resort filter. It runs last so redaction always
// wins over any earlier transform that expanded a value.
func sensitiveFilter(groups []string, a slog.Attr) slog.Attr {
	if a.Value.Kind() != slog.KindString {
		return a
	}
	if _, ok := a.Value.Any().(string); !ok {
		return a
	}
	v := a.Value.String()
	if a.Key == "username" || a.Key == "user" || containsAny(a.Key, sensitiveKeys) {
		return slog.String(a.Key, "[REDACTED]")
	}
	if redacted := redactValue(v); redacted != v {
		return slog.String(a.Key, redacted)
	}
	return a
}

func replaceAttr(groups []string, a slog.Attr) slog.Attr {
	if a.Key == "error" {
		err, ok := a.Value.Any().(error)
		if !ok {
			return a
		}
		return slog.GroupAttrs("error", errorAttrs(err)...)
	}
	return sensitiveFilter(groups, a)
}
func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	httpPort := flag.Int("port", 8899, "port to listen on")
	dataDir := flag.String("data", "./data", "directory to store data")
	flag.Parse()

	status := run(ctx, cancel, *httpPort, *dataDir)
	cancel()
	os.Exit(status)
}
func initializeLogger(logFile string) (*slog.Logger, closeFunc, error) {
	debugHandler := &colorHandler{Handler: slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			a = replaceAttr(groups, a)
			if a.Key == slog.LevelKey {
				if level, ok := a.Value.Any().(slog.Level); ok {
					a.Value = slog.StringValue(colorize(level, level.String()))
				}
			}
			return a
		},
	})}
	if logFile != "" {
		rotating := &lumberjack.Logger{
			Filename:   logFile,
			MaxSize:    1,
			MaxAge:     28,
			MaxBackups: 10,
			LocalTime:  false,
			Compress:   true,
		}
		bufferedFile := bufio.NewWriterSize(rotating, 8192)
		closeFunc := func() error {
			if err := bufferedFile.Flush(); err != nil {
				return err
			}
			return rotating.Close()
		}
		standardHandler := slog.NewJSONHandler(bufferedFile, &slog.HandlerOptions{
			ReplaceAttr: replaceAttr,
		})
		return slog.New(slog.NewMultiHandler(debugHandler, standardHandler)), closeFunc, nil
	}
	return slog.New(debugHandler), nil, nil
}
func run(ctx context.Context, cancel context.CancelFunc, httpPort int, dataDir string) int {
	logFile := os.Getenv("LINKO_LOG_FILE")
	logger, closeFunc, err := initializeLogger(logFile)
	if closeFunc != nil {
		defer closeFunc()
	}
	logger = logger.With(
		slog.String("git_sha", build.GitSHA),
		slog.String("build_time", build.BuildTime),
		slog.String("env", os.Getenv("ENV")),
	)
	hostname, err := os.Hostname()
	if err == nil {
		logger = logger.With(slog.String("hostname", hostname))
	}
	st, err := store.New(dataDir, logger)

	if err != nil {
		logger.Error("failed to create store",
			"error", pkgerr.WithStack(err))
		return 1
	}
	s := newServer(*st, httpPort, cancel, logger)
	var serverErr error
	go func() {
		serverErr = s.start()
	}()

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.shutdown(shutdownCtx); err != nil {
		s.logger.Error("failed to shutdown server",
			"error", pkgerr.WithStack(err))
		return 1
	}
	if serverErr != nil {
		s.logger.Error("server error",
			"error", pkgerr.WithStack(serverErr))
		return 1
	}
	return 0
}
