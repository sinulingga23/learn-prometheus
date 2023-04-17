package monitoring

import "github.com/prometheus/client_golang/prometheus"

var (
	TotalRequestApi = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "learn_prometheus",
		Name:      "total_request_api",
		Help:      "Show total request of api",
	}, []string{"service_name", "http_status", "http_method"})
	DurationRequestAPi = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "learn_prometheus",
		Name:      "duration_request_api",
		Buckets:   []float64{0.1, 0.15, 0.2, 0.25, 0.3},
	}, []string{"service_name", "http_status", "http_method"})
)
