package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	//HTTPReqDuration metric:http_request_duration_seconds
	HTTPReqDuration *prometheus.HistogramVec
	//HTTPReqTotal metric:http_request_total
	HTTPReqTotal *prometheus.CounterVec
	// TaskRunning metric:task_running
	TaskRunning *prometheus.GaugeVec
)

func init() {
	// 监控接口请求耗时
	// 指标类型是 Histogram
	HTTPReqDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "http request latencies in seconds",
		Buckets: nil,
	}, []string{"method", "path"})
	// "method"、"path" 是 label

	// 监控接口请求次数
	// 指标类型是 Counter
	HTTPReqTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "total number of http requests",
	}, []string{"method", "path", "status"})
	// "method"、"path"、"status" 是 label

	prometheus.MustRegister(
		HTTPReqDuration,
		HTTPReqTotal,
	)
}

//Metric metric middleware
func Metrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		duration := float64(time.Since(start)) / float64(time.Second)

		path := c.Request.URL.Path

		// 请求数加1
		HTTPReqTotal.With(prometheus.Labels{
			"method": c.Request.Method,
			"path":   path,
			"status": strconv.Itoa(c.Writer.Status()),
		}).Inc()

		//  记录本次请求处理时间
		HTTPReqDuration.With(prometheus.Labels{
			"method": c.Request.Method,
			"path":   path,
		}).Observe(duration)
	}
}

func Monitor(c *gin.Context) {
	h := promhttp.Handler()
	h.ServeHTTP(c.Writer, c.Request)
}
