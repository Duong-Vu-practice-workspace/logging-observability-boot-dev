package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"boot.dev/linko/internal/store"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	httpPort := flag.Int("port", 8899, "port to listen on")
	dataDir := flag.String("data", "./data", "directory to store data")
	flag.Parse()

	status := run(ctx, cancel, *httpPort, *dataDir)
	cancel()
	os.Exit(status)
}
func initializeLogger(logFile string) (*log.Logger, error) {
	if logFile != "" {
		file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
		if err != nil {
			return nil, fmt.Errorf("failed to open log file: %w", err)
		}
		bufferedFile := bufio.NewWriterSize(file, 8192)
		multiWriter := io.MultiWriter(os.Stderr, bufferedFile)
		return log.New(multiWriter, "", log.LstdFlags), nil
	}
	return log.New(os.Stderr, "", log.LstdFlags), nil
}
func run(ctx context.Context, cancel context.CancelFunc, httpPort int, dataDir string) int {
	logFile := os.Getenv("LINKO_LOG_FILE")
	logger, err := initializeLogger(logFile)

	st, err := store.New(dataDir, logger)

	if err != nil {
		logger.Printf("failed to create store: %v", err)
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
		s.logger.Printf("failed to shutdown server: %v", err)
		return 1
	}
	if serverErr != nil {
		s.logger.Printf("server error: %v\n", serverErr)
		return 1
	}
	return 0
}
