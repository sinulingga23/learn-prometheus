package monitoring

import "github.com/prometheus/client_golang/prometheus"

var (
	TotalRequestApi = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "learn_prometheus",
		Name:      "total_request_api",
		Help:      "Show total request of api",
	}, []string{"service_name", "http_status"})
)
