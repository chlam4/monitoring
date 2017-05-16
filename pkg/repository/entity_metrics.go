package repository

import (
	"bytes"
	"fmt"
	"github.com/turbonomic/turbo-go-monitoring/pkg/model"
	"github.com/golang/glog"
)

// A MetricMap is a 2-layer map of metric values, indexed by the resource type and the metric property type
// For example, all metrics of an entity can be stored in such a map.
type EntityMetricMap map[model.ResourceType]map[model.MetricPropType]model.MetricValue

type EntityMetricKey struct {
	ResourceType model.ResourceType
	PropType     model.MetricPropType
}

// SetMetricValue sets the metric value in the MetricMap for the given resource type and the metric property type
func (metricMap EntityMetricMap) SetMetricValue(
	resourceType model.ResourceType,
	propType model.MetricPropType,
	value model.MetricValue,
) {
	resourceMap, exists := metricMap[resourceType]
	if !exists {
		resourceMap = make(map[model.MetricPropType]model.MetricValue)
		metricMap[resourceType] = resourceMap
	}
	resourceMap[propType] = value
}

// GetMetricValue retrieves the metric value from the MetricMap for the given resource type and the metric property type
func (metricMap EntityMetricMap) GetMetricValue(
	resourceType model.ResourceType,
	propType model.MetricPropType,
) (model.MetricValue, error) {
	resourceMap, exists := metricMap[resourceType]
	if !exists {
		glog.V(4).Infof("Cannot find metrics for resource %s\n", resourceType)
		return model.MetricValue(0), fmt.Errorf("missing metrics for resource %s", resourceType)
	}
	value, exists := resourceMap[propType]
	if !exists {
		glog.V(4).Infof("Cannot find metrics for type %s\n", propType)
		return model.MetricValue(0), fmt.Errorf("missing metrics for type %s:%s", resourceType, propType)
	}
	return value, nil
}

func (metricMap EntityMetricMap) String() string {
	var buffer bytes.Buffer
	for resourceType, resourceMap := range metricMap {
		for prop, value := range resourceMap {
			line := fmt.Sprintf("\t\t%s::%s : %f\n", resourceType, prop, value)
			buffer.WriteString(line)
		}
	}
	return buffer.String()
}
