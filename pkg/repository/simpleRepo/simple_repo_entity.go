package simpleRepo

import (
	"github.com/chlam4/monitoring/pkg/model"
	"github.com/chlam4/monitoring/pkg/model/metric"
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

// GetId returns the id of the entity
func (repoEntity SimpleMetricRepoEntity) GetId() model.EntityId {
	return repoEntity.entityId
}

// GetType returns the type of the entity
func (repoEntity SimpleMetricRepoEntity) GetType() model.EntityType {
	return repoEntity.entityType
}

// GetNodeIp returns the type of the entity
func (repoEntity SimpleMetricRepoEntity) GetNodeIp() model.NodeIp {
	return repoEntity.nodeIp
}

func (repoEntity SimpleMetricRepoEntity) GetAllMetrics() repository.EntityMetricMap {
	return repoEntity.metricMap
}

// GetMetricValue returns the metric value of the given resource type and metric property type
func (repoEntity SimpleMetricRepoEntity) GetMetricValue(metricKey repository.EntityMetricKey) (metric.MetricValue, error) {
	return repoEntity.metricMap.GetMetricValue(metricKey.ResourceType, metricKey.PropType)
}

func (repoEntity SimpleMetricRepoEntity) SetMetricValue(
	key repository.EntityMetricKey,
	value metric.MetricValue,
) {
	repoEntity.metricMap.SetMetricValue(key.ResourceType, key.PropType, value)
}
