package metric


// MetricDef defines a metric to collect and how it is set in the metric repository
type MetricDef struct {
	EntityType   EntityType
	resourceType ResourceType
	propType     MetricPropType
	metricSetter MetricSetter // Setter for the property
}

// MakeMetricDef makes a MetricDef with the given entity type, resource type and metric property type, and the
// default metric setter.
func MakeMetricDefWithDefaultSetter(entityType EntityType, resourceType ResourceType, propType MetricPropType) MetricDef {
	setter := DefaultMetricSetter{entityType: entityType, resourceType: resourceType, propType: propType}
	return MetricDef{EntityType: entityType, resourceType: resourceType, propType: propType, metricSetter: &setter}
}
