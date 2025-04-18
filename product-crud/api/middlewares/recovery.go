package middlewares

import (
    "net/http"
    "product-crud/pkg/logger"
)

func RecoveryMiddleware(logger *logger.Logger) func(next http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            defer func() {
                if err := recover(); err != nil {
                    logger.Error("Recover from panic",
                        "error", err,
                        "path", r.URL.Path,
                        "method", r.Method,
                    )

                    http.Error(w, "Internal server error", http.StatusInternalServerError)
                }
            }()
            
            next.ServeHTTP(w, r) // Added this line - was missing in original
        })
    }
}