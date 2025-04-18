package middlewares

import (
    "net/http"
    "product-crud/pkg/logger"
    "time"
)

func LoggingMiddleware(logger *logger.Logger) func(next http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            start := time.Now()

            crw := &customResponseWriter{
                ResponseWriter: w,
                statusCode:     http.StatusOK,
            }

            next.ServeHTTP(crw, r)

            duration := time.Since(start)
            logger.Info("HTTP Request",
                "method", r.Method,
                "path", r.URL.Path,
                "status", crw.statusCode,
                "duration", duration,
                "user_agent", r.UserAgent(),
                "remote_addr", r.RemoteAddr,
                "request_id", r.Header.Get("X-Request-ID"))
        })
    }
}