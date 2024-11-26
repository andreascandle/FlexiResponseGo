package observability

import (
	"net/http"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total HTTP requests",
		},
		[]string{"method", "path", "status_code", "host", "protocol"},
	)

	responseDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_response_duration_seconds",
			Help:    "Histogram of response durations for HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path", "status_code", "host", "protocol"},
	)
)

func init() {
	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(responseDuration)
}

// sanitizePath prevents label explosion by normalizing paths.
func sanitizePath(path string) string {
	// Example: Replace dynamic segments like IDs with a placeholder
	segments := strings.Split(path, "/")
	for i, segment := range segments {
		if isDynamic(segment) {
			segments[i] = "{param}"
		}
	}
	return strings.Join(segments, "/")
}

func isDynamic(segment string) bool {
	// Check if the segment is a numeric ID or UUID, customize as needed.
	return strings.ContainsAny(segment, "0123456789")
}

// MetricsMiddleware collects HTTP request metrics.
func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		rr := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(rr, r)

		duration := time.Since(startTime).Seconds()
		sanitizedPath := sanitizePath(r.URL.Path)

		requestCounter.WithLabelValues(
			r.Method, sanitizedPath, http.StatusText(rr.statusCode), r.Host, r.Proto,
		).Inc()

		responseDuration.WithLabelValues(
			r.Method, sanitizedPath, http.StatusText(rr.statusCode), r.Host, r.Proto,
		).Observe(duration)
	})
}

// responseRecorder captures status codes for HTTP responses.
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rr *responseRecorder) WriteHeader(code int) {
	rr.statusCode = code
	rr.ResponseWriter.WriteHeader(code)
}

// HTTPHandlerForMetrics exposes Prometheus metrics at /metrics endpoint.
func HTTPHandlerForMetrics() http.Handler {
	return promhttp.Handler()
}
