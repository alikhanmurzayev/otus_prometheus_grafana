package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"strconv"
	"time"
)

var (
	MetricsRequestLatency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "app_request_latency_seconds",
		Help: "Application Request Latency",
	}, []string{"method", "status"})

	MetricsRequestCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "app_request_count",
		Help: "Application Request Count",
	}, []string{"method", "status"})

	Metrics5xxErrorCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "app_5xx_error_count",
		Help: "Application 5xx Error Count",
	}, []string{"method", "status"})

	metrics = []prometheus.Collector{
		MetricsRequestLatency,
		MetricsRequestCount,
		Metrics5xxErrorCount,
	}
)

func init() {
	for _, metric := range metrics {
		if err := prometheus.Register(metric); err != nil {
			panic(fmt.Errorf("could not register metric: %w", err))
		}
	}
}

func Instrumenting(c *fiber.Ctx) error {
	startTime := time.Now()
	err := c.Next()
	labels := prometheus.Labels{
		"method": string(c.Request().Header.Method()),
		"status": strconv.Itoa(c.Response().StatusCode()),
	}
	MetricsRequestCount.With(labels).Inc()
	MetricsRequestLatency.With(labels).Observe(time.Now().Sub(startTime).Seconds())
	if c.Response().StatusCode() >= http.StatusInternalServerError {
		Metrics5xxErrorCount.With(labels).Inc()
	}
	return err
}
