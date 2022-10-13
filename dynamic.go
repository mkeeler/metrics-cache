package metricscache

import (
	"sync/atomic"

	"github.com/armon/go-metrics"
)

type dynamicValue struct {
	metrics.MetricSink
}

type DynamicSink struct {
	sink atomic.Value
}

func NewDynamicSink(sink metrics.MetricSink) *DynamicSink {
	d := &DynamicSink{}
	d.ReplaceSink(sink)
	return d
}

func (d *DynamicSink) ReplaceSink(sink metrics.MetricSink) {
	if sink == nil {
		sink = &metrics.BlackholeSink{}
	}
	d.sink.Store(&dynamicValue{sink})
}

func (d *DynamicSink) SetGauge(key []string, val float32) {
	d.sink.Load().(*dynamicValue).SetGauge(key, val)
}

func (d *DynamicSink) SetGaugeWithLabels(key []string, val float32, labels []metrics.Label) {
	d.sink.Load().(*dynamicValue).SetGaugeWithLabels(key, val, labels)
}

func (d *DynamicSink) EmitKey(key []string, val float32) {
	d.sink.Load().(*dynamicValue).EmitKey(key, val)
}

func (d *DynamicSink) IncrCounter(key []string, val float32) {
	d.sink.Load().(*dynamicValue).IncrCounter(key, val)
}

func (d *DynamicSink) IncrCounterWithLabels(key []string, val float32, labels []metrics.Label) {
	d.sink.Load().(*dynamicValue).IncrCounterWithLabels(key, val, labels)
}

func (d *DynamicSink) AddSample(key []string, val float32) {
	d.sink.Load().(*dynamicValue).AddSample(key, val)
}

func (d *DynamicSink) AddSampleWithLabels(key []string, val float32, labels []metrics.Label) {
	d.sink.Load().(*dynamicValue).AddSampleWithLabels(key, val, labels)
}
