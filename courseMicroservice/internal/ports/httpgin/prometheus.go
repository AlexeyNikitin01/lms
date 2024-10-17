package httpgin

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Количество HTTP-запросов по методам и статусам.",
		},
		[]string{"method", "status"},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Распределение времени обработки HTTP-запросов.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path", "method"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
}

func prometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start).Seconds()
		path := c.FullPath()
		method := c.Request.Method
		status := http.StatusText(c.Writer.Status())

		httpRequestsTotal.WithLabelValues(method, status).Inc()
		httpRequestDuration.WithLabelValues(path, method).Observe(duration)
	}
}
