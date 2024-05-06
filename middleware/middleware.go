package middleware

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	HttpDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_response_duration_seconds",
		Help:    "Duration of HTTP responses in seconds",
		Buckets: prometheus.DefBuckets,
	}, []string{"path", "method"})
)

func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)

		// Record the duration in the histogram
		HttpDuration.WithLabelValues(r.URL.Path, r.Method).Observe(duration.Seconds())
	})
}
