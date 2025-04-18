package middlewares

import (
    "net/http"
)

type customResponseWriter struct {
    http.ResponseWriter
    statusCode int
}

func (crw *customResponseWriter) WriteHeader(statusCode int) {
    crw.statusCode = statusCode
    crw.ResponseWriter.WriteHeader(statusCode)
}