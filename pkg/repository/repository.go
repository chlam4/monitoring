package repository

import (
	"github.com/chlam4/monitoring/pkg/model"
)

// A Repository defines a set of interfaces to access its entities and their metrics
type Repository interface {
	GetEntity(id model.EntityId) RepositoryEntity
	GetAllEntityInstances() []RepositoryEntity
	GetEntityInstances(entityType model.EntityType) []RepositoryEntity
	SetEntityInstances([]RepositoryEntity)
	SetMetricValue(entityId model.EntityId, metricKey EntityMetricKey, value model.MetricValue)
}
