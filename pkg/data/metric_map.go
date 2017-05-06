package data

import (
	"github.com/golang/glog"
	"fmt"
)

// A MetricMap is a 2-layer map of metric values, indexed by the resource type and the metric property type
// For example, all metrics of an entity can be stored in such a map.
type MetricMap map[ResourceType]map[MetricPropType]*Metric

type Metric struct {
	value *float64
}

// SetResourceMetric sets the metric value in the MetricMap for the given resource type and the metric property type
func (resourceMetrics MetricMap) SetResourceMetric(resourceType ResourceType, metricType MetricPropType, value *float64) {
	resourceMap, exists := resourceMetrics[resourceType]
	if !exists {
		resourceMap = make(map[MetricPropType]*Metric)
		resourceMetrics[resourceType] = resourceMap
	}
	metric, ok := resourceMap[metricType]
	if !ok {
		metric = &Metric{}
	}
	metric.value = value
	resourceMap[metricType] = metric
}

// GetResourceMetric retrieves the metric value from the MetricMap for the given resource type and the metric property type
func (resourceMetrics MetricMap) GetResourceMetric(resourceType ResourceType, metricType MetricPropType) (*Metric, error) {
	resourceMap, exists := resourceMetrics[resourceType]
	if !exists {
		glog.V(4).Infof("Cannot find metrics for resource %s\n", resourceType)
		return nil, fmt.Errorf("missing metrics for resource %s", resourceType)
	}
	metric, exists := resourceMap[metricType]
	if !exists {
		glog.V(4).Infof("Cannot find metrics for type %s\n", metricType)
		return nil, fmt.Errorf("missing metrics for type %s:%s", resourceType, metricType)
	}
	return metric, nil
}

// printMetrics prints out all metrics to the log
func (resourceMetrics MetricMap) printMetrics() {
	for rt, resourceMap := range resourceMetrics {
		for mkey, metric := range resourceMap {
			if (metric != nil) {
				glog.Infof("\t\t%s::%s : %f\n", rt, mkey, *metric.value)
			} else {
				glog.Infof("\t\t%s::%s : %f\n", rt, mkey, metric.value)

			}
		}
	}
}
