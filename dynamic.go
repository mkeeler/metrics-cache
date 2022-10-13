package metricscache

import (
	"sync/atomic"

	"github.com/armon/go-metrics"
)

// Go's atomic.Value is kind of nasty in that it panics if the
// concrete type of the value being Stored/Loaded changes. In
// this case the underlying value type SHOULD change. I could
// have used some unsafe pointer code or taken this approach
// which is to embed the interface I care about in a struct
// and Store/Load that struct type.
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
