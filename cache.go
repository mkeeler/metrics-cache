package metricscache

import (
	"sync"

	"github.com/armon/go-metrics"
)

type metric struct {
	key    []string
	val    float32
	labels []metrics.Label
}

type CacheSink struct {
	mu sync.Mutex

	gauges   []metric
	counters []metric
	keys     []metric
	samples  []metric
}

type MetricsCacheReplayer interface {
	SetGaugeWithLabels(key []string, val float32, labels []metrics.Label)
	EmitKey(key []string, val float32)
	IncrCounterWithLabels(key []string, val float32, labels []metrics.Label)
	AddSampleWithLabels(key []string, val float32, labels []metrics.Label)
}

func (c *CacheSink) Replay(r MetricsCacheReplayer) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, m := range c.gauges {
		r.SetGaugeWithLabels(m.key, m.val, m.labels)
	}

	for _, m := range c.counters {
		r.IncrCounterWithLabels(m.key, m.val, m.labels)
	}

	for _, m := range c.keys {
		r.EmitKey(m.key, m.val)
	}

	for _, m := range c.samples {
		r.AddSampleWithLabels(m.key, m.val, m.labels)
	}

	c.gauges = nil
	c.counters = nil
	c.keys = nil
	c.samples = nil
}

func (c *CacheSink) SetGauge(key []string, val float32) {
	c.SetGaugeWithLabels(key, val, nil)
}

func (c *CacheSink) SetGaugeWithLabels(key []string, val float32, labels []metrics.Label) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.gauges = append(c.gauges, metric{key: key, val: val, labels: labels})
}

func (c *CacheSink) EmitKey(key []string, val float32) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.keys = append(c.keys, metric{key: key, val: val})
}

func (c *CacheSink) IncrCounter(key []string, val float32) {
	c.IncrCounterWithLabels(key, val, nil)
}

func (c *CacheSink) IncrCounterWithLabels(key []string, val float32, labels []metrics.Label) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters = append(c.counters, metric{key: key, val: val, labels: labels})
}

func (c *CacheSink) AddSample(key []string, val float32) {
	c.AddSampleWithLabels(key, val, nil)
}

func (c *CacheSink) AddSampleWithLabels(key []string, val float32, labels []metrics.Label) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.samples = append(c.samples, metric{key: key, val: val, labels: labels})
}
