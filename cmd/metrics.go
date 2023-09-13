package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"math/rand"
	"time"
	
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	requestDurations := prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "A histogram of the HTTP request durations in seconds.",
		Buckets: prometheus.ExponentialBuckets(0.1, 1.5, 5),
	})
	
	epochCount := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "epoch",
		Help: "Show chain epoch",
	})
	
	// Create non-global registry.
	registry := prometheus.NewRegistry()
	
	// Add go runtime metrics and process collectors.
	registry.MustRegister(
		collectors.NewGoCollector(),
		//collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		//requestDurations,
		epochCount,
	)
	
	go func() {
		for {
			// Record fictional latency.
			now := time.Now()
			requestDurations.(prometheus.ExemplarObserver).ObserveWithExemplar(
				time.Since(now).Seconds(), prometheus.Labels{"dummyID": fmt.Sprint(rand.Intn(100000))},
			)
			epochCount.Add(1)
			time.Sleep(600 * time.Millisecond)
		}
	}()
	
	r := gin.Default()
	r.GET("/metrics", func(ctx *gin.Context) {
		h := promhttp.HandlerFor(
			registry,
			promhttp.HandlerOpts{
				EnableOpenMetrics: true,
			})
		h.ServeHTTP(ctx.Writer, ctx.Request)
		return
	})
	
	// Expose /metrics HTTP endpoint using the created custom registry.
	//http.Handle(
	//	"/metrics", promhttp.HandlerFor(
	//		registry,
	//		promhttp.HandlerOpts{
	//			EnableOpenMetrics: true,
	//		}),
	//)
	// To test: curl -H 'Accept: application/openmetrics-text' localhost:8080/metrics
	//log.Fatalln(http.ListenAndServe(":8080", nil))
	
	r.Run(":8080")
}
