package metric

import (
	"github.com/golang/glog"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	"fmt"
)

// Interface for Repository Entity
type RepositoryEntity interface {
	GetId() string
	GetType() proto.EntityDTO_EntityType
	GetResourceMetrics() MetricMap
	GetResourceMetric(resourceType ResourceType, metricType MetricPropType) (*MetricValue, error)
}

// Interface for a Repository
type Repository interface {
	GetEntity(entityType proto.EntityDTO_EntityType, id string) RepositoryEntity
	GetEntityInstances(entityType proto.EntityDTO_EntityType) []RepositoryEntity
}

func PrintEntity(entity RepositoryEntity) {
	glog.Infof("Entity %s::%s\n", entity.GetType(), entity.GetId())
	fmt.Printf("Entity %s::%s\n", entity.GetType(), entity.GetId())
	resourceMetrics := entity.GetResourceMetrics()
	resourceMetrics.printMetrics()
}

//func PrintRepository(repository Repository) {
//	PrintEntity(repository.GetAgentEntity())
//	taskEntities := repository.GetTaskEntities()
//	for _, taskEntity := range taskEntities {
//		PrintEntity(taskEntity)
//	}
//	containerEntities := repository.GetContainerEntities()
//	for _, containerEntity := range containerEntities {
//		PrintEntity(containerEntity)
//	}
//}
