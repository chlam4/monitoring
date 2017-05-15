package simpleRepo

import (
	"github.com/chlam4/monitoring/pkg/model"
	"github.com/chlam4/monitoring/pkg/repository"
	"github.com/golang/glog"
	"fmt"
)

// SimpleMetricRepo is a simple implementation of the metric repository
type SimpleMetricRepo map[model.EntityId]repository.RepositoryEntity

// NewSimpleMetricRepo returns a new, empty instance of SimpleMetricRepo
func NewSimpleMetricRepo() repository.Repository {
	return SimpleMetricRepo{}
}

func (repo SimpleMetricRepo) GetAllEntityInstances() []repository.RepositoryEntity {
	entities := []repository.RepositoryEntity{}
	for _, repoEntity := range repo {
		entities = append(entities, repoEntity)
	}
	return entities
}

func (repo SimpleMetricRepo) GetEntity(id model.EntityId) repository.RepositoryEntity {
	return repo[id]
}

func (repo SimpleMetricRepo) GetEntityInstances(entityType model.EntityType) []repository.RepositoryEntity {
	entities := []repository.RepositoryEntity{}
	for _, repoEntity := range repo {
		if repoEntity.GetType() == entityType {
			entities = append(entities, repoEntity)
		}
	}
	return entities
}

func (repo SimpleMetricRepo) SetEntityInstances(repoEntities []repository.RepositoryEntity) {
	for _, repoEntity := range repoEntities {
		repo[repoEntity.GetId()] = repoEntity
	}
}

func (repo SimpleMetricRepo) SetMetricValue(
	entityId model.EntityId,
	key repository.EntityMetricKey,
	value model.MetricValue,
) error {
	repoEntity := repo.GetEntity(entityId)
	if repoEntity == nil {
		glog.V(4).Infof("Repository: %v", repo)
		return fmt.Errorf("Repository entity is not found with id %v", entityId)
	}
	repoEntity.SetMetricValue(key, value)
	return nil
}
