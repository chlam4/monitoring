package prometheus

import (
	"github.com/chlam4/monitoring/pkg/model/entity"
	"github.com/chlam4/monitoring/pkg/model/property"
	"github.com/chlam4/monitoring/pkg/model/resource"
	"github.com/chlam4/monitoring/pkg/model"
)

// Query defines the data type of Prometheus query strings
type Query string

// MetricQueryMap defines for each given MetricKey which Prometheus query to use to collect the metric data
var MetricQueryMap = map[model.MetricKey]Query{
	//
	// TODO: Revisit Node CPU stats more carefully -
	// 	Do we need per-core stats, or user vs system vs iowait vs steal stats?
	//
	{entity.NODE, resource.CPU, property.USED}:    "100 * (sum(delta(node_cpu{mode!='idle'}[10m])) by (job, instance) / sum(delta(node_cpu[10m])) by (job, instance))",
	{entity.NODE, resource.MEM, property.USED}:    "node_memory_Active",
	{entity.NODE, resource.MEM, property.CAP}:     "node_memory_MemTotal",
	{entity.NODE, resource.MEM, property.AVERAGE}: "avg_over_time(node_memory_Active[10m])",
	{entity.NODE, resource.MEM, property.PEAK}:    "max_over_time(node_memory_Active[10m])",
	//
	// TODO: Handle multiple interfaces per node, transmit and receive
	//
	{entity.NODE, resource.NETWORK, property.USED}: "sum(rate(node_network_receive_bytes[10m]) + rate(node_network_transmit_bytes[10m])) by (job, instance)",
	{entity.POD, resource.MEM, property.USED}: "sum(container_memory_usage_bytes) by (instance, pod_name)",
}
