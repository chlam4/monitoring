package prometheus

import (
	"fmt"
	"github.com/chlam4/monitoring/pkg/model"
	"github.com/chlam4/monitoring/pkg/model/entity"
	"github.com/chlam4/monitoring/pkg/model/property"
	"github.com/chlam4/monitoring/pkg/model/resource"
	prometheusModel "github.com/prometheus/common/model"
	"strings"
)

// Query defines the data type of Prometheus query strings
type Query struct {
	queryString string
	entityId    func(sample *prometheusModel.Sample) (model.EntityId, error)
}

// MetricQueryMap defines for each given MetricKey which Prometheus query to use to collect the metric data
var MetricQueryMap = map[model.MetricKey]Query{
	//
	// TODO: Revisit Node CPU stats more carefully -
	// 	Do we need per-core stats, or user vs system vs iowait vs steal stats?
	//
	{entity.NODE, resource.CPU, property.USED}:    {"100 * (sum(delta(node_cpu{mode!='idle'}[10m])) by (job, instance) / sum(delta(node_cpu[10m])) by (job, instance))", getEntityIdFromNodeExporter},
	{entity.NODE, resource.MEM, property.USED}:    {"node_memory_Active", getEntityIdFromNodeExporter},
	{entity.NODE, resource.MEM, property.CAP}:     {"node_memory_MemTotal", getEntityIdFromNodeExporter},
	{entity.NODE, resource.MEM, property.AVERAGE}: {"avg_over_time(node_memory_Active[10m])", getEntityIdFromNodeExporter},
	{entity.NODE, resource.MEM, property.PEAK}:    {"max_over_time(node_memory_Active[10m])", getEntityIdFromNodeExporter},
	//
	// TODO: Handle multiple interfaces per node, transmit and receive
	//
	{entity.NODE, resource.NETWORK, property.USED}: {"sum(rate(node_network_receive_bytes[10m]) + rate(node_network_transmit_bytes[10m])) by (job, instance)", getEntityIdFromNodeExporter},
	{entity.POD, resource.MEM, property.USED}:      {"sum(container_memory_usage_bytes) by (instance, pod_name)", getEntityIdFromNodeExporter},
}

func getEntityIdFromNodeExporter(sample *prometheusModel.Sample) (model.EntityId, error) {
	instanceName := string(sample.Metric["instance"])
	if instanceName == "" {
		return "", fmt.Errorf("Instance is missing in the returned metric: %v", sample.Metric)
	}
	entityId := model.EntityId(strings.Split(instanceName, ":")[0])
	return entityId, nil

}
