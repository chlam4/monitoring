package simpleRepo

import (
	"github.com/chlam4/monitoring/pkg/model"
	"github.com/chlam4/monitoring/pkg/repository"
)

// SimpleMetricRepoEntity is a simple implementation of the RepositoryEntity
type SimpleMetricRepoEntity struct {
	entityType model.EntityType
	entityId   model.EntityId
	nodeIp     model.NodeIp
	metricMap  repository.EntityMetricMap
}

func NewSimpleMetricRepoEntity(
	entityType model.EntityType,
	entityId model.EntityId,
	nodeIp model.NodeIp,
) repository.RepositoryEntity {
	return SimpleMetricRepoEntity{entityId: entityId, entityType: entityType, nodeIp: nodeIp, metricMap: make(repository.EntityMetricMap)}
}

func (repoEntity SimpleMetricRepoEntity) GetId() model.EntityId {
	return repoEntity.entityId
}

func (repoEntity SimpleMetricRepoEntity) GetType() model.EntityType {
	return repoEntity.entityType
}

func (repoEntity SimpleMetricRepoEntity) GetNodeIp() model.NodeIp {
	return repoEntity.nodeIp
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
