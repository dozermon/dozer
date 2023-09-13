package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type config struct {
	withGoCollector bool
	withDBCollector bool
}

type Registry struct {
	registry *prometheus.Registry
}

func NewRegistry() *Registry {
	return &Registry{registry: prometheus.NewRegistry()}
}

func (r *Registry) Register(collector ...prometheus.Collector) (err error) {
	for _, v := range collector {
		err = r.registry.Register(v)
		if err != nil {
			return
		}
	}
	return
}

func (r *Registry) Serve(group *gin.RouterGroup) {
	group.GET("/", func(ctx *gin.Context) {
		h := promhttp.HandlerFor(
			r.registry,
			promhttp.HandlerOpts{
				EnableOpenMetrics: true,
			})
		h.ServeHTTP(ctx.Writer, ctx.Request)
	})
	return
}
