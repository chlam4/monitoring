package simpleRepo

import (
	"testing"
	"github.com/turbonomic/turbo-go-monitoring/pkg/model/entity"
	"github.com/turbonomic/turbo-go-monitoring/pkg/model"
)

var TestEntities = []struct {
	entityType model.EntityType
	entityId   model.EntityId
}{
	{entity.NODE, "1.2.3.4"},
	{entity.NODE, "192.168.99.100"},
	{entity.POD, "abc"},
	{entity.APP, "xyz"},
}

func TestSimpleMetricRepoEntity_GetId_GetType(t *testing.T) {

	for _, testEntity := range TestEntities {
		repoEntity := NewSimpleMetricRepoEntity(testEntity.entityType, testEntity.entityId)
		if repoEntity.GetType() != testEntity.entityType {
			t.Errorf("Retrieved type %v from repo entity %v is not the same as input %v",
				repoEntity.GetType(), repoEntity, testEntity.entityType)
		}
		if repoEntity.GetId() != testEntity.entityId {
			t.Errorf("Retrieved id %v from repo entity %v is not the same as input %v",
				repoEntity.GetId(), repoEntity, testEntity.entityId)
		}
	}
}
