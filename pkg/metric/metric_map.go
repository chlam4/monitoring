package metric

import (
	"fmt"
	"github.com/golang/glog"
	"bytes"
)

// A MetricMap is a 2-layer map of metric values, indexed by the resource type and the metric property type
// For example, all metrics of an entity can be stored in such a map.
type MetricMap map[ResourceType]map[MetricPropType]MetricValue
// MetricValue is a float64
type MetricValue float64

// SetResourceMetric sets the metric value in the MetricMap for the given resource type and the metric property type
func (resourceMetrics MetricMap) SetResourceMetric(resourceType ResourceType, propType MetricPropType, value MetricValue) {
	resourceMap, exists := resourceMetrics[resourceType]
	if !exists {
		resourceMap = make(map[MetricPropType]MetricValue)
		resourceMetrics[resourceType] = resourceMap
	}
	resourceMap[propType] = value
}

// GetResourceMetric retrieves the metric value from the MetricMap for the given resource type and the metric property type
func (resourceMetrics MetricMap) GetResourceMetric(resourceType ResourceType, metricType MetricPropType) (MetricValue, error) {
	resourceMap, exists := resourceMetrics[resourceType]
	if !exists {
		glog.V(4).Infof("Cannot find metrics for resource %s\n", resourceType)
		return MetricValue(0), fmt.Errorf("missing metrics for resource %s", resourceType)
	}
	value, exists := resourceMap[metricType]
	if !exists {
		glog.V(4).Infof("Cannot find metrics for type %s\n", metricType)
		return MetricValue(0), fmt.Errorf("missing metrics for type %s:%s", resourceType, metricType)
	}
	return value, nil
}

func (resourceMetrics MetricMap) String() string {
	var buffer bytes.Buffer
	for resourceType, resourceMap := range resourceMetrics {
		for prop, value := range resourceMap {
			line := fmt.Sprintf("\t\t%s::%s : %f\n", resourceType, prop, value)
			buffer.WriteString(line)
		}
	}
	return buffer.String()
}
