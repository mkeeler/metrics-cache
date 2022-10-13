package metricscache

import (
	"sync"

	"github.com/armon/go-metrics"
)

type DeferredSink struct {
	configured sync.Once
	cache      *CacheSink
	dynamic    *DynamicSink
}

func NewDeferredSink() *DeferredSink {
	cache := &CacheSink{}
	return &DeferredSink{cache: cache, dynamic: NewDynamicSink(cache)}
}

func (d *DeferredSink) Configure(sink metrics.MetricSink) {
	d.configured.Do(func() {
		d.dynamic.ReplaceSink(sink)
		d.cache.Replay(sink)
	})
}

func (d *DeferredSink) SetGauge(key []string, val float32) {
	d.dynamic.SetGauge(key, val)
}

func (d *DeferredSink) SetGaugeWithLabels(key []string, val float32, labels []metrics.Label) {
	d.dynamic.SetGaugeWithLabels(key, val, labels)
}

func (d *DeferredSink) EmitKey(key []string, val float32) {
	d.dynamic.EmitKey(key, val)
}

func (d *DeferredSink) IncrCounter(key []string, val float32) {
	d.dynamic.IncrCounter(key, val)
}

func (d *DeferredSink) IncrCounterWithLabels(key []string, val float32, labels []metrics.Label) {
	d.dynamic.IncrCounterWithLabels(key, val, labels)
}

func (d *DeferredSink) AddSample(key []string, val float32) {
	d.dynamic.AddSample(key, val)
}

func (d *DeferredSink) AddSampleWithLabels(key []string, val float32, labels []metrics.Label) {
	d.dynamic.AddSampleWithLabels(key, val, labels)
}
