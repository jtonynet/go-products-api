package middleware

import (
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/jtonynet/go-products-api/config"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

var (
	startTime = time.Now()

	// TIMER
	processUptime = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "process_uptime_seconds",
			Help: "Total uptime of the process in seconds",
		},
		[]string{},
	)

	// CPU
	systemCPUUsage = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "system_cpu_usage",
			Help: "System CPU usage percent",
		},
		[]string{"host"},
	)

	processCPUUsage = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "process_cpu_usage",
			Help: "Process CPU usage percent",
		},
		[]string{"host"},
	)

	systemLoadAverage1m = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "system_load_average_1m",
			Help: "System load average over the last 1 minute",
		},
		[]string{"host"},
	)

	systemCPUCores = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "system_cpu_cores",
			Help: "Number of CPU cores on the system",
		},
		[]string{"host"},
	)

	// MEMORY
	memoryAlloc = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "memory_alloc_bytes",
			Help: "Memory alloc of the process in seconds",
		},
		[]string{"host"},
	)

	memoryTotalAlloc = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "memory_total_alloc_bytes",
			Help: "Memory total alloc of the process in seconds",
		},
		[]string{"host"},
	)

	processMemoryUsed = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "memory_used_bytes",
			Help: "Memory used by the process in bytes",
		},
		[]string{"host"},
	)

	processMemoryLimit = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "memory_limit_bytes",
			Help: "Memory limit of the process in bytes",
		},
		[]string{"host"},
	)

	// HTTP REQUESTS
	requestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "echo",
			Subsystem: "http",
			Name:      "requests_total",
			Help:      "How many HTTP requests processed, partitioned by status code and HTTP method.",
		},
		[]string{"code", "method", "host", "path"},
	)

	requestsDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "echo",
		Subsystem: "http",
		Name:      "request_duration",
		Help:      "The HTTP request latency bucket.",
	}, []string{"code", "method", "host", "path"})

	maxDurations        = make(map[string]float64)
	requestsMaxDuration = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "echo",
			Subsystem: "http",
			Name:      "requests_seconds_max",
			Help:      "The maximum HTTP request latency.",
		},
		[]string{"code", "method", "host", "path"},
	)
)

func metricsLoop(cfg *config.API) {
	go func() {
		for {
			uptime := time.Since(startTime).Seconds()
			processUptime.WithLabelValues().Set(uptime)

			systemCPUPercent, err := cpu.Percent(0, true)
			if err == nil {
				systemCPUUsage.WithLabelValues(cfg.Host).Set(systemCPUPercent[0])
			}

			processCPUPercent, err := cpu.Percent(0, false)
			if err == nil {
				processCPUUsage.WithLabelValues(cfg.Host).Set(processCPUPercent[0])
			}

			loadAvg, err := load.Avg()
			if err == nil {
				systemLoadAverage1m.WithLabelValues(cfg.Host).Set(loadAvg.Load1)
			}

			cpuInfo, err := cpu.Info()
			if err == nil {
				systemCPUCores.WithLabelValues(cfg.Host).Set(float64(len(cpuInfo)))
			}

			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			memoryAlloc.WithLabelValues(cfg.Host).Set(float64(m.Alloc))
			memoryTotalAlloc.WithLabelValues(cfg.Host).Set(float64(m.TotalAlloc))

			memInfo, err := mem.VirtualMemory()
			if err == nil {
				processMemoryUsed.WithLabelValues(cfg.Host).Set(float64(memInfo.Used))
				processMemoryLimit.WithLabelValues(cfg.Host).Set(float64(memInfo.Total))
			}

			time.Sleep(time.Duration(cfg.MetricsRefreshIntervalInSec) * time.Second)
		}
	}()
}

func InitPrometheus(cfg *config.API) {
	prometheus.Register(processUptime)

	prometheus.Register(systemCPUUsage)
	prometheus.Register(processCPUUsage)
	prometheus.Register(systemLoadAverage1m)
	prometheus.Register(systemCPUCores)

	prometheus.Register(memoryAlloc)
	prometheus.Register(memoryTotalAlloc)

	prometheus.Register(requestsTotal)
	prometheus.Register(requestsDuration)
	prometheus.Register(requestsMaxDuration)

	metricsLoop(cfg)
}

func Prometheus() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			initialRequestTime := time.Now()

			next(c)

			statusCode := strconv.Itoa(c.Response().Status)
			labels := []string{statusCode, c.Request().Method, c.Request().Host, c.Path()}

			requestsTotal.WithLabelValues(labels...).Inc()

			duration := time.Since(initialRequestTime).Seconds()
			requestsDuration.WithLabelValues(labels...).Observe(duration)

			maxDuration, exists := maxDurations[strings.Join(labels, "_")]
			if !exists || duration > maxDuration {
				maxDurations[strings.Join(labels, "_")] = duration
				requestsMaxDuration.WithLabelValues(labels...).Set(duration)
			}

			return nil
		}
	}
}
