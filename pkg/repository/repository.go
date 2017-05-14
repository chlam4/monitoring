package repository

import (
	"github.com/chlam4/monitoring/pkg/model"
	"github.com/chlam4/monitoring/pkg/model/metric"
)

// A Repository defines a set of interfaces to access its entities and their metrics
type Repository interface {
	GetEntity(id model.EntityId) RepositoryEntity
	GetAllEntityInstances() []RepositoryEntity
	GetEntityInstances(entityType model.EntityType) []RepositoryEntity
	SetEntityInstances([]RepositoryEntity)
	SetMetricValue(entityId model.EntityId, metricKey metric.MetricKey, value metric.MetricValue)
}
