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
	{entity.NODE, resource.CPU, property.USED}:    {"100 * (sum(delta(node_cpu{mode!='idle'}[10m])) by (job, instance) / sum(delta(node_cpu[10m])) by (job, instance))", fromInstanceName},
	{entity.NODE, resource.MEM, property.USED}:    {"node_memory_Active", fromInstanceName},
	{entity.NODE, resource.MEM, property.CAP}:     {"node_memory_MemTotal", fromInstanceName},
	{entity.NODE, resource.MEM, property.AVERAGE}: {"avg_over_time(node_memory_Active[10m])", fromInstanceName},
	{entity.NODE, resource.MEM, property.PEAK}:    {"max_over_time(node_memory_Active[10m])", fromInstanceName},
	//
	// TODO: Handle multiple interfaces per node, transmit and receive
	//
	{entity.NODE, resource.NETWORK, property.USED}: {"sum(rate(node_network_receive_bytes[10m]) + rate(node_network_transmit_bytes[10m])) by (job, instance)", fromInstanceName},
	{entity.POD, resource.MEM, property.USED}:      {"sum(container_memory_usage_bytes{image!=\"\"}) by (pod_name)", fromPodName},
}

// fromInstanceName() extracts the entity id from the instance field of the given Prometheus metric sample.
func fromInstanceName(sample *prometheusModel.Sample) (model.EntityId, error) {
	instanceName := string(sample.Metric["instance"])
	if instanceName == "" {
		return "", fmt.Errorf("Instance is missing in the returned metric: %v", sample.Metric)
	}
	entityId := model.EntityId(strings.Split(instanceName, ":")[0])
	return entityId, nil

}

// fromPodName() extracts the entity id from the pod name field of the given Prometheus metric sample.
func fromPodName(sample *prometheusModel.Sample) (model.EntityId, error) {
	podName := string(sample.Metric["pod_name"])
	if podName == "" {
		return "", fmt.Errorf("Pod name is missing in the returned metric: %v", sample.Metric)
	}
	entityId := model.EntityId(podName)
	return entityId, nil

}
