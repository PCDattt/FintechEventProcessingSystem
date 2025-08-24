package prometheus

import "github.com/prometheus/client_golang/prometheus"

var (
	AccountRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests handled, labeled by handler and status code.",
		},
		[]string{"handler", "status"},
	)
)

func Init() {
	prometheus.MustRegister(AccountRequestsTotal)
}