package repository

import (
	"github.com/chlam4/monitoring/pkg/model"
)

// RepositoryEntity defines a set of interfaces to access a repository entity in the metric Repository.
// It has per-entity info such as id, type, and node ip, and interfaces to get and set metric values.
type RepositoryEntity interface {
	// GetId() returns the associated entity id
	GetId() model.EntityId

	// GetType() returns the type of the associated entity
	GetType() model.EntityType

	// GetTypedId() returns the type and id of the associated entity
	GetTypedId() model.EntityTypedId

	// GetAllMetrics() returns all metrics of the entity in the form of EntityMetricMap
	GetAllMetrics() EntityMetricMap

	// GetMetricValue() returns the metric value associated with the given key in this entity
	GetMetricValue(metricKey EntityMetricKey) (model.MetricValue, error)

	// SetMetricValue() sets the given metric key-value in this entity
	SetMetricValue(metricKey EntityMetricKey, value model.MetricValue)
}
