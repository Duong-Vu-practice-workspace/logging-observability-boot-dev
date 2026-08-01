package main

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func Test_requestLogger(t *testing.T) {
	logBuffer := &bytes.Buffer{}
	logger := slog.New(slog.NewTextHandler(logBuffer, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Time(slog.TimeKey, time.Date(2023, 10, 1, 12, 34, 57, 0, time.UTC))
			}
			if a.Key == "duration" {
				return slog.Duration(a.Key, 42*time.Millisecond)
			}
			return a
		},
	}))
	requestLoggerMiddleware := requestLogger(logger)
	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.ReadAll(r.Body)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("ok"))
	})
	loggedHandler := requestLoggerMiddleware(dummyHandler)
	req := httptest.NewRequest("POST", "http://lin.ko/api/stats", strings.NewReader("hello"))
	const requestID = "test-request-id"
	req.Header.Set("X-Request-ID", requestID)
	rr := httptest.NewRecorder()
	loggedHandler.ServeHTTP(rr, req)
	const expectedLogString = `time=2023-10-01T12:34:57.000Z level=INFO msg="Served request" method=POST path=/api/stats client_ip=192.0.2.1:1234 request_id=test-request-id duration=42ms request_body_bytes=5 response_status=201 response_body_bytes=2 username=""` + "\n"
	const expectedStatusCode = http.StatusCreated

	if logBuffer.String() != expectedLogString {
		t.Errorf("expected log %q, got %q", expectedLogString, logBuffer.String())
	}
	if rr.Code != expectedStatusCode {
		t.Errorf("expected status %d, got %d", expectedStatusCode, rr.Code)
	}
	if rr.Header().Get("X-Request-ID") != requestID {
		t.Errorf("expected response header X-Request-ID %q, got %q", requestID, rr.Header().Get("X-Request-ID"))
	}
}

func Test_requestLogger_generatesRequestID(t *testing.T) {
	logBuffer := &bytes.Buffer{}
	logger := slog.New(slog.NewTextHandler(logBuffer, nil))
	requestLoggerMiddleware := requestLogger(logger)
	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	loggedHandler := requestLoggerMiddleware(dummyHandler)
	req := httptest.NewRequest("GET", "http://lin.ko/", nil)
	rr := httptest.NewRecorder()
	loggedHandler.ServeHTTP(rr, req)

	responseID := rr.Header().Get("X-Request-ID")
	if responseID == "" {
		t.Fatal("expected generated X-Request-ID in response header")
	}
	if !strings.Contains(logBuffer.String(), "request_id="+responseID) {
		t.Errorf("expected log to contain request_id=%s, got %q", responseID, logBuffer.String())
	}
}
