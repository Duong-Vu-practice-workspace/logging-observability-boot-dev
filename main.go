package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"boot.dev/linko/internal/build"
	"boot.dev/linko/internal/store"
	pkgerr "github.com/pkg/errors"
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

func replaceAttr(groups []string, a slog.Attr) slog.Attr {
	if a.Key == "error" {
		err, ok := a.Value.Any().(error)
		if !ok {
			return a
		}
		return slog.GroupAttrs("error", errorAttrs(err)...)
	}
	return a
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
	if logFile != "" {
		file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to open log file: %w", err)
		}
		bufferedFile := bufio.NewWriterSize(file, 8192)
		closeFunc := func() error {
			if err := bufferedFile.Flush(); err != nil {
				return err
			}
			return file.Close()
		}
		debugHandler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			ReplaceAttr: replaceAttr,
		})
		standardHandler := slog.NewJSONHandler(bufferedFile, &slog.HandlerOptions{
			ReplaceAttr: replaceAttr,
		})
		return slog.New(slog.NewMultiHandler(debugHandler, standardHandler)), closeFunc, nil
	}
	return slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		ReplaceAttr: replaceAttr,
	})), nil, nil
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
