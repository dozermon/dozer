package metrics

import (
	"database/sql"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

type CounterOpts = prometheus.CounterOpts
type HistogramOpts = prometheus.HistogramOpts
type GaugeOpts = prometheus.GaugeOpts
type SummaryOpts = prometheus.SummaryOpts
type ProcessCollectorOpts = collectors.ProcessCollectorOpts

func NewCounter(opts CounterOpts) prometheus.Counter {
	return prometheus.NewCounter(opts)
}

func NewHistogram(opts HistogramOpts) prometheus.Histogram {
	return prometheus.NewHistogram(opts)
}

func NewGauge(ops GaugeOpts) prometheus.Gauge {
	return prometheus.NewGauge(ops)
}

func NewSummary(ops SummaryOpts) prometheus.Summary {
	return prometheus.NewSummary(ops)
}

func NewGoCollector(opts ProcessCollectorOpts) prometheus.Collector {
	return collectors.NewProcessCollector(opts)
}

func NewDBCollector(db *sql.DB, name string) prometheus.Collector {
	return collectors.NewDBStatsCollector(db, name)
}
