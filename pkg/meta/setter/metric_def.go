package setter

import (
	"github.com/chlam4/monitoring/pkg/model"
	"github.com/chlam4/monitoring/pkg/model/metric"
	"github.com/chlam4/monitoring/pkg/repository"
)

// MetricDef defines a metric to collect and how it is set in the metric repository
type MetricDef struct {
	EntityType   model.EntityType
	ResourceType model.ResourceType
	PropType model.MetricPropType
	MetricSetter MetricSetter // Setter for the property
}

// MakeMetricDef makes a MetricDef with the given entity type, resource type and metric property type, and the
// default metric setter.
func MakeMetricDefWithDefaultSetter(
	entityType model.EntityType,
	resourceType model.ResourceType,
	propType model.MetricPropType,
) MetricDef {
	setter := DefaultMetricSetter{}
	return MetricDef{EntityType: entityType, ResourceType: resourceType, PropType: propType, MetricSetter: setter}
}

// ToMetricKey() returns the MetricKey corresponding to this MetricDef
func (metricDef *MetricDef) ToMetricKey() metric.MetricKey {
	return metric.MetricKey{EntityType: metricDef.EntityType, ResourceType: metricDef.ResourceType, PropType: metricDef.PropType}
}

// ToEntityMetricKey() returns the EntityMetricKey corresponding to this MetricDef
func (metricDef *MetricDef) ToEntityMetricKey() repository.EntityMetricKey {
	return repository.EntityMetricKey{ResourceType: metricDef.ResourceType, PropType: metricDef.PropType}
}

