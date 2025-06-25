package monitoring

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	RequestsTotal  *prometheus.CounterVec
	RequestLatency *prometheus.HistogramVec
}

func NewMetrics() *Metrics {
	metrics := &Metrics{
		RequestsTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_requests_total",
				Help: "Total number of HTTP requests",
			},
			[]string{"method", "endpoint", "status"},
		),
		RequestLatency: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "http_request_duration_seconds",
				Help:    "HTTP request latency",
				Buckets: prometheus.DefBuckets,
			},
			[]string{"method", "endpoint"},
		),
	}

	// register metrics
	prometheus.MustRegister(
		metrics.RequestsTotal,
		metrics.RequestLatency,
	)

	return metrics
}
