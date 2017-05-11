package metric

import (
	"github.com/golang/glog"
)

// A Repository keeps a set of entities and their metrics
type Repository interface {
	GetEntity(entityType EntityType, id EntityId) RepositoryEntity
	GetEntityInstances(entityType EntityType) []RepositoryEntity
	SetEntityInstances([]RepositoryEntity)
}

// RepositoryEntity is an entity in the Repository
// It contains information such as id, type, node ip and all metrics
type RepositoryEntity interface {
	GetId() EntityId
	GetType() EntityType
	GetNodeIp() NodeIp
	GetResourceMetrics() MetricMap
	GetResourceMetric(resourceType ResourceType, propType MetricPropType) (MetricValue, error)
	SetMetricValue(resourceType ResourceType, propType MetricPropType, value MetricValue)
}

func PrintEntity(entity RepositoryEntity) {
	glog.Infof("Entity %s::%s\n", entity.GetType(), entity.GetId())
	resourceMetrics := entity.GetResourceMetrics()
	resourceMetrics.printMetrics()
}

// SimpleMetricRepo is a simple implementation of the metric repository
type SimpleMetricRepo map[EntityType]map[EntityId]RepositoryEntity

// NewSimpleMetricRepo returns a new, empty instance of SimpleMetricRepo
func NewSimpleMetricRepo() Repository {
	 return make(SimpleMetricRepo)
}

func (repo SimpleMetricRepo) GetEntity(entityType EntityType, id EntityId) RepositoryEntity {
	repoEntityMap, exists := repo[entityType]
	if !exists {
		return nil
	}
	return repoEntityMap[id]
}

func (repo SimpleMetricRepo) GetEntityInstances(entityType EntityType) []RepositoryEntity {
	repoEntityMap, exists := repo[entityType]
	if !exists {
		return nil
	}
	entities := make([]RepositoryEntity, len(repoEntityMap))
	for _, entity := range repoEntityMap {
		entities = append(entities, entity)
	}
	return entities
}

func (repo SimpleMetricRepo) SetEntityInstances(repoEntities []RepositoryEntity) {
	for _, repoEntity := range repoEntities {
		id2EntityMap, exists := repo[repoEntity.GetType()]
		if !exists {
			id2EntityMap = make(map[EntityId]RepositoryEntity)
			repo[repoEntity.GetType()] = id2EntityMap
		}
		id2EntityMap[repoEntity.GetId()] = repoEntity
	}
}

// SimpleMetricRepoEntity is a simple implementation of the RepositoryEntity
type SimpleMetricRepoEntity struct {
	entityType EntityType
	entityId   EntityId
	nodeIp     NodeIp
	metricMap  MetricMap
}

func NewSimpleMetricRepoEntity(entityType EntityType, entityId EntityId, nodeIp NodeIp) RepositoryEntity {
	return SimpleMetricRepoEntity{entityId: entityId, entityType: entityType, nodeIp: nodeIp, metricMap: make(MetricMap)}
}

// GetId returns the id of the entity
func (repoEntity SimpleMetricRepoEntity) GetId() EntityId {
	return repoEntity.entityId
}

// GetType returns the type of the entity
func (repoEntity SimpleMetricRepoEntity) GetType() EntityType {
	return repoEntity.entityType
}

// GetNodeIp returns the type of the entity
func (repoEntity SimpleMetricRepoEntity) GetNodeIp() NodeIp {
	return repoEntity.nodeIp
}

// GetResourceMetrics returns the map of metrics for the given repository entity
func (repoEntity SimpleMetricRepoEntity) GetResourceMetrics() MetricMap {
	return repoEntity.metricMap
}

// GetResourceMetric returns the metric value of the given resource type and metric property type
func (repoEntity SimpleMetricRepoEntity) GetResourceMetric(resourceType ResourceType, propType MetricPropType) (MetricValue, error) {
	return repoEntity.metricMap.GetResourceMetric(resourceType, propType)
}

func (repoEntity SimpleMetricRepoEntity) SetMetricValue(resourceType ResourceType, propType MetricPropType, value MetricValue) {
	repoEntity.metricMap.SetResourceMetric(resourceType, propType, value)
}