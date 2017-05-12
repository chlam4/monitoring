package metric

type MetricKey struct {
	EntityType   EntityType
	ResourceType ResourceType
	PropType     MetricPropType
}

// MetricDef defines a metric to collect and how it is set in the metric repository
type MetricDef struct {
	EntityType   EntityType
	ResourceType ResourceType
	PropType     MetricPropType
	metricSetter MetricSetter // Setter for the property
}

// ToMetricKey returns the corresponding Metric Key for the given MetricDef
func (metricDef *MetricDef) ToMetricKey() MetricKey {
	metricKey := MetricKey{EntityType: metricDef.EntityType, ResourceType: metricDef.ResourceType, PropType: metricDef.PropType}
	return metricKey
}

// MakeMetricDef makes a MetricDef with the given entity type, resource type and metric property type, and the
// default metric setter.
func MakeMetricDefWithDefaultSetter(entityType EntityType, resourceType ResourceType, propType MetricPropType) MetricDef {
	setter := DefaultMetricSetter{entityType: entityType, resourceType: resourceType, propType: propType}
	return MetricDef{EntityType: entityType, ResourceType: resourceType, PropType: propType, metricSetter: setter}
}