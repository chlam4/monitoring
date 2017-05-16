// The model package contains types and constants in the core of this monitoring library.
package model

// EntityType defines the data type of the entity type.
// Various types of entity are listed in the entity subpackage.
type EntityType string

// EntityId defines the data type of the entity id.
// The entity type and id combo uniquely identify an entity in the system.
type EntityId string

// EntityTypedId is the entity type + id; it uniquely identifies the entity in the system.
type EntityTypedId struct {
	EntityType EntityType
	EntityId   EntityId
}

// ResourceType defines the data type of the resource type.
// Various types of resource are listed in the resource subpackage.
type ResourceType string

// MetricPropType defines the data type of the metric property type.
// Various types of metric property are listed in the property subpackage.
type MetricPropType string

// MetricValue defines the data type of the value of a metric.
type MetricValue float64

// MetricKey defines the type of metric key.
// It is composed of entity type, resource type and metric property type.
type MetricKey struct {
	EntityType   EntityType
	ResourceType ResourceType
	PropType     MetricPropType
}
