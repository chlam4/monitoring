package template

import (
	"github.com/chlam4/monitoring/pkg/model"
	"github.com/chlam4/monitoring/pkg/model/metric"
	"github.com/chlam4/monitoring/pkg/repository"
)

// MonitoringTemplate defines a metric to collect and how its value is set in the metric repository
type MonitoringTemplate struct {
	MetricKey    metric.MetricKey
	MetricSetter MetricSetter // Setter for the property
}

// MakeMonitoringTemplateWithDefaultSetter makes a MonitoringTemplate with given entity type, resource type and metric
// property type, and the default metric setter.
func MakeMonitoringTemplateWithDefaultSetter(
	entityType model.EntityType,
	resourceType model.ResourceType,
	propType model.MetricPropType,
) MonitoringTemplate {
	metricKey := metric.MetricKey{EntityType: entityType, ResourceType: resourceType, PropType: propType}
	setter := DefaultMetricSetter{}
	return MonitoringTemplate{MetricKey: metricKey, MetricSetter: setter}
}

// ToMetricKey() returns the MetricKey corresponding to this MetricDef
func (monTemplate *MonitoringTemplate) ToMetricKey() metric.MetricKey {
	return monTemplate.MetricKey
}

// ToEntityMetricKey() returns the EntityMetricKey corresponding to this MetricDef
func (monTemplate *MonitoringTemplate) ToEntityMetricKey() repository.EntityMetricKey {
	return repository.EntityMetricKey{ResourceType: monTemplate.MetricKey.ResourceType, PropType: monTemplate.MetricKey.PropType}
}
