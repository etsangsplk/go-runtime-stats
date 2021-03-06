package publisher

import (
	"reflect"
	"testing"

	"github.com/jtaczanowski/go-graphite-client"
	"github.com/jtaczanowski/go-runtime-stats/collector"
)

func TestShouldPrepareDataToSend(t *testing.T) {
	testCollector := collector.NewCollector()
	graphiteClient := graphite.NewClient("localhost", 2003, "prefix", "udp")
	testPublisher := NewPublisher(testCollector, graphiteClient)

	excepted := map[string]float64{
		"cpu.count":                        0,
		"cpu.goroutines_number":            0,
		"cpu.cgo_calls_number_delta":       0,
		"cpu.cgo_calls_number_total":       0,
		"mem.general.lookups_number_delta": 0,
		"mem.general.lookups_number_total": 0,
		"mem.general.mallocs_number_total": 0,
		"mem.general.frees_number_total":   0,
		"mem.stack.sys_bytes":              0,
		"mem.stack.mspan_sys_bytes":        0,
		"mem.stack.mcache_inuse_bytes":     0,
		"mem.stack.mcache_sys_bytes":       0,
		"mem.general.mallocs_number_delta": 0,
		"mem.general.frees_number_delta":   0,
		"mem.heap.sys_bytes":               0,
		"mem.heap.inuse_bytes":             0,
		"mem.othersys_bytes":               0,
		"mem.stack.mspan_inuse_bytes":      0,
		"mem.general.total_bytes":          0,
		"mem.heap.alloc_bytes":             0,
		"mem.heap.idle_bytes":              0,
		"mem.heap.released_bytes":          0,
		"mem.general.alloc_bytes":          0,
		"mem.general.sys_bytes":            0,
		"mem.heap.objects_number":          0,
		"mem.stack.inuse_bytes":            0,
		"gc.pause_ns":                      0,
		"gc.number_total":                  0,
		"gc.next_bytes":                    0,
		"gc.time_from_last_gc_s":           0,
		"gc.pause_ns_total_delta":          0,
		"gc.pause_ns_total":                0,
		"gc.cpu_fraction_total":            0,
		"gc.sys_bytes":                     0,
		"gc.between_period_s":              0,
		"gc.pause_last_ns":                 0,
		"gc.number_delta":                  0,
	}

	data := testPublisher.prepareDataToSend()

	for _, metric := range data {
		var found bool
		for _, exceptedMetric := range excepted {
			eq := reflect.DeepEqual(metric, exceptedMetric)
			if eq {
				found = true
			}
		}
		if found {
			continue
		} else {
			t.Errorf("Metric %v not found in excepted metrics", metric)
		}
	}
}
