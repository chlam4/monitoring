package setter

import (
	"github.com/chlam4/monitoring/pkg/model"
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
