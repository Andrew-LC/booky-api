package middleware

import (
    "net/http"
    "strconv"
    "time"
    "bookmark-api/metrics"
)

type statusWriter struct {
    http.ResponseWriter
    status int
}

func (sw *statusWriter) WriteHeader(code int) {
    sw.status = code
    sw.ResponseWriter.WriteHeader(code)
}

func MetricsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/metrics" {
            next.ServeHTTP(w, r)
            return
        }
        start := time.Now()
        sw := &statusWriter{ResponseWriter: w, status: 200}
        next.ServeHTTP(sw, r)
        metrics.HttpRequestsTotal.WithLabelValues(
            r.Method,
            r.URL.Path,
            strconv.Itoa(sw.status),
        ).Inc()
        metrics.HttpRequestDuration.WithLabelValues(
            r.Method,
            r.URL.Path,
        ).Observe(time.Since(start).Seconds())
    })
}
