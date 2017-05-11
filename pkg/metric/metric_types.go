package metric

type EntityId string
type EntityType string
type NodeIp string
type ResourceType string
type MetricPropType string

type MetricKey struct {
	entityType EntityType
	resourceType ResourceType
	propType MetricPropType
}
