package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/armon/go-metrics"
	metricscache "github.com/mkeeler/metrics-cache"
)

func main() {
	deferredSink := metricscache.NewDeferredSink()
	conf := metrics.DefaultConfig("deferred")
	conf.EnableRuntimeMetrics = false
	conf.EnableHostname = false
	conf.EnableHostnameLabel = true
	m, _ := metrics.New(conf, deferredSink)
	m.SetGauge([]string{"gauge"}, 1)
	m.MeasureSince([]string{"sample"}, time.Now().Add(-50*time.Millisecond))
	m.IncrCounter([]string{"counter"}, 1)

	time.Sleep(time.Second)

	inmem := metrics.NewInmemSink(10*time.Second, 60*time.Second)
	deferredSink.Configure(inmem)

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "   ")
	values, _ := inmem.DisplayMetrics(nil, nil)
	fmt.Printf("Cached Startup Values\n")
	_ = enc.Encode(values)

	m.IncrCounter([]string{"counter"}, 2)
	m.MeasureSince([]string{"sample"}, time.Now().Add(-200*time.Millisecond))
	m.SetGauge([]string{"gauge"}, 100)

	fmt.Printf("Values After More Modifications\n")
	values, _ = inmem.DisplayMetrics(nil, nil)
	_ = enc.Encode(values)
}
