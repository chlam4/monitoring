package prometheus

import (
	"github.com/chlam4/monitoring/pkg/metric"
)

type Query string

// MetricQueryMap defines for each given MetricKey which Prometheus query to use to collect the metric data
var MetricQueryMap = map[metric.MetricKey]Query{
	//
	// TODO: Revisit Node CPU stats more carefully -
	// 	Do we need per-core stats, or user vs system vs iowait vs steal stats?
	//
	{metric.NODE, metric.CPU, metric.USED}:    "100 * (sum(delta(node_cpu{mode!='idle'}[10m])) by (job, instance) / sum(delta(node_cpu[10m])) by (job, instance))",
	{metric.NODE, metric.MEM, metric.USED}:    "node_memory_Active",
	{metric.NODE, metric.MEM, metric.CAP}:     "node_memory_MemTotal",
	{metric.NODE, metric.MEM, metric.AVERAGE}: "avg_over_time(node_memory_Active[10m])",
	{metric.NODE, metric.MEM, metric.PEAK}:    "max_over_time(node_memory_Active[10m])",
	//
	// TODO: Handle multiple interfaces per node, transmit and receive
	//
	{metric.NODE, metric.NETWORK, metric.USED}: "sum(rate(node_network_receive_bytes[10m]) + rate(node_network_transmit_bytes[10m])) by (job, instance)",
}
