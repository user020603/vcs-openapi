package middlewares

import (
    "net/http"

    "github.com/google/uuid"
)

func RequestIDMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        requestID := r.Header.Get("X-Request-ID")
        if requestID == "" {
            requestID = generateRequestID()
        }

        w.Header().Set("X-Request-ID", requestID)

        ctx := r.Context()
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func generateRequestID() string {
    return uuid.New().String()
}