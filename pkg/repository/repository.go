// The repository package contains definitions about metric repository.
// A metric repository is composed of a list of repository entities, each of which contains metrics related to the entity.
package repository

import (
	"github.com/chlam4/monitoring/pkg/model"
)

// A Repository defines a set of interfaces to access its entities and their metrics
type Repository interface {
	// GetEntity() returns the RepositoryEntity associated with the given entity type and id
	GetEntity(entityType model.EntityType, entityId model.EntityId) (RepositoryEntity, error)

	// GetAllEntities() returns the list of all RepositoryEntity's in the repository
	GetAllEntities() []RepositoryEntity

	// GetEntitiesByType() returns the list of RepositoryEntity's matching the given entity type
	GetEntitiesByType(entityType model.EntityType) []RepositoryEntity

	// SetEntities() updates the repository with the given set of RepositoryEntity's
	SetEntities([]RepositoryEntity)

	// SetMetricValue() sets the value of the given metric in the repository
	SetMetricValue(
		entityType model.EntityType,
		entityId model.EntityId,
		metricKey EntityMetricKey,
		value model.MetricValue,
	) error
}
