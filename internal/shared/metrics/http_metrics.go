package metrics

import (
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	registerOnce sync.Once

	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total count of HTTP requests.",
		},
		[]string{"method", "path", "status"},
	)
	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request latency in seconds.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path", "status"},
	)
	httpInFlight = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "http_requests_in_flight",
			Help: "Current number of in-flight HTTP requests.",
		},
	)
)

func register() {
	registerOnce.Do(func() {
		prometheus.MustRegister(httpRequestsTotal, httpRequestDuration, httpInFlight)
	})
}

func HTTPMiddleware() fiber.Handler {
	register()

	return func(c *fiber.Ctx) error {
		httpInFlight.Inc()
		start := time.Now()
		err := c.Next()
		duration := time.Since(start).Seconds()
		httpInFlight.Dec()

		method := c.Method()
		path := c.Route().Path
		if path == "" {
			path = c.Path()
		}
		status := strconv.Itoa(c.Response().StatusCode())

		httpRequestsTotal.WithLabelValues(method, path, status).Inc()
		httpRequestDuration.WithLabelValues(method, path, status).Observe(duration)

		return err
	}
}

func Handler() fiber.Handler {
	register()
	return adaptor.HTTPHandler(promhttp.Handler())
}
