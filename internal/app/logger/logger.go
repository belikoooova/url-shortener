package logger

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

var Log *zap.Logger = zap.NewNop()

func Initialize(level string) error {
	parsedLevel, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return err
	}

	cfg := zap.NewProductionConfig()
	cfg.Level = parsedLevel

	lg, err := cfg.Build()
	if err != nil {
		return err
	}

	Log = lg
	return nil
}

func WithLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		URI := r.RequestURI
		method := r.Method

		responseData := &ResponseData{
			status: 0,
			size:   0,
		}
		lw := LoggingResponseWriter{
			next:         w,
			responseData: responseData,
		}

		next.ServeHTTP(&lw, r)
		duration := time.Since(start)

		Log.Info("got request, send response",
			zap.String("uri", URI),
			zap.String("method", method),
			zap.String("duration", duration.String()),
			zap.Int("status", responseData.status),
			zap.Int("size", responseData.size),
		)
	})
}

type ResponseData struct {
	status int
	size   int
}

type LoggingResponseWriter struct {
	next         http.ResponseWriter
	responseData *ResponseData
}

func (r *LoggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.next.Write(b)
	r.responseData.size += size
	return size, err
}

func (r *LoggingResponseWriter) WriteHeader(statusCode int) {
	r.next.WriteHeader(statusCode)
	r.responseData.status = statusCode
}

func (r *LoggingResponseWriter) Header() http.Header {
	return r.next.Header()
}
