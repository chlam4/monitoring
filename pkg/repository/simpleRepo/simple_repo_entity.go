package simpleRepo

import (
	"github.com/turbonomic/turbo-go-monitoring/pkg/model"
	"github.com/turbonomic/turbo-go-monitoring/pkg/repository"
)

// SimpleMetricRepoEntity is a simple implementation of the RepositoryEntity
type SimpleMetricRepoEntity struct {
	entityType model.EntityType
	entityId   model.EntityId
	metricMap  repository.EntityMetricMap
}

func NewSimpleMetricRepoEntity(
	entityType model.EntityType,
	entityId model.EntityId,
) repository.RepositoryEntity {
	return SimpleMetricRepoEntity{entityId: entityId, entityType: entityType, metricMap: make(repository.EntityMetricMap)}
}

func (repoEntity SimpleMetricRepoEntity) GetId() model.EntityId {
	return repoEntity.entityId
}

func (repoEntity SimpleMetricRepoEntity) GetType() model.EntityType {
	return repoEntity.entityType
}

func (repoEntity SimpleMetricRepoEntity) GetTypedId() model.EntityTypedId {
	return model.EntityTypedId{EntityType: repoEntity.entityType, EntityId: repoEntity.entityId}
}

func (repoEntity SimpleMetricRepoEntity) GetAllMetrics() repository.EntityMetricMap {
	return repoEntity.metricMap
}

func (repoEntity SimpleMetricRepoEntity) GetMetricValue(metricKey repository.EntityMetricKey) (model.MetricValue, error) {
	return repoEntity.metricMap.GetMetricValue(metricKey.ResourceType, metricKey.PropType)
}

func (repoEntity SimpleMetricRepoEntity) SetMetricValue(
	key repository.EntityMetricKey,
	value model.MetricValue,
) {
	repoEntity.metricMap.SetMetricValue(key.ResourceType, key.PropType, value)
}
