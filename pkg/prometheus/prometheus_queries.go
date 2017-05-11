package prometheus

import (
	"github.com/chlam4/monitoring/pkg/metric"
)

type Query string

// MetricQueryMap defines which Prometheus query to use to collect the metric data for each given MetricKey
var MetricQueryMap = map[metric.MetricKey]Query{
	{metric.NODE, metric.MEM, metric.USED}: "node_memory_Active",
}