package simpleRepo

import (
	"fmt"
	"github.com/turbonomic/turbo-go-monitoring/pkg/model"
	"github.com/turbonomic/turbo-go-monitoring/pkg/repository"
)

// SimpleMetricRepo is a simple implementation of the metric repository
type SimpleMetricRepo map[model.EntityType]map[model.EntityId]repository.RepositoryEntity

// NewSimpleMetricRepo returns a new, empty instance of SimpleMetricRepo
func NewSimpleMetricRepo() repository.Repository {
	return SimpleMetricRepo{}
}

func (repo SimpleMetricRepo) GetAllEntities() []repository.RepositoryEntity {
	entities := []repository.RepositoryEntity{}
	for _, id2EntityMap := range repo {
		for _, repoEntity := range id2EntityMap {
			entities = append(entities, repoEntity)
		}
	}
	return entities
}

func (repo SimpleMetricRepo) GetEntity(
	entityType model.EntityType,
	entityId model.EntityId,
) (repository.RepositoryEntity, error) {

	id2EntityMap, exists := repo[entityType]
	if !exists {
		return nil, fmt.Errorf("Entity type %v does not exist in the repository: %v", entityType, repo)
	}
	repoEntity, exists := id2EntityMap[entityId]
	if !exists {
		return nil, fmt.Errorf("Entity %v/%v does not exist in the repository: %v", entityType, entityId, repo)
	}
	return repoEntity, nil
}

func (repo SimpleMetricRepo) GetEntitiesByType(entityType model.EntityType) []repository.RepositoryEntity {
	entities := []repository.RepositoryEntity{}
	for _, repoEntity := range repo[entityType] {
		entities = append(entities, repoEntity)
	}
	return entities
}

func (repo SimpleMetricRepo) SetEntities(inputEntities []repository.RepositoryEntity) {
	for _, inputEntity := range inputEntities {
		id2EntityMap, exists := repo[inputEntity.GetType()]
		if !exists {
			id2EntityMap = map[model.EntityId]repository.RepositoryEntity{}
			repo[inputEntity.GetType()] = id2EntityMap
		}
		id2EntityMap[inputEntity.GetId()] = inputEntity
	}
}

func (repo SimpleMetricRepo) SetMetricValue(
	entityType model.EntityType,
	entityId model.EntityId,
	key repository.EntityMetricKey,
	value model.MetricValue,
) error {
	repoEntity, err := repo.GetEntity(entityType, entityId)
	if err != nil {
		return err
	}
	repoEntity.SetMetricValue(key, value)
	return nil
}
