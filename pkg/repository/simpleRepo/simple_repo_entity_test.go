package simpleRepo

import (
	"testing"
	"github.com/chlam4/monitoring/pkg/model/entity"
	"github.com/chlam4/monitoring/pkg/model"
)

var TestEntities = []struct {
	entityType model.EntityType
	entityId   model.EntityId
	nodeIp     model.NodeIp
}{
	{entity.NODE, "foo", "1.2.3.4"},
	{entity.NODE, "bar", "192.168.99.100"},
	{entity.POD, "123", "10.10.172.236"},
	{entity.APP, "xyz", "127.0.0.1"},
}

func TestSimpleMetricRepoEntity_GetId_GetType(t *testing.T) {

	for _, testEntity := range TestEntities {
		repoEntity := NewSimpleMetricRepoEntity(testEntity.entityType, testEntity.entityId, testEntity.nodeIp)
		if repoEntity.GetType() != testEntity.entityType {
			t.Errorf("Retrieved type %v from repo entity %v is not the same as input %v",
				repoEntity.GetType(), repoEntity, testEntity.entityType)
		}
		if repoEntity.GetId() != testEntity.entityId {
			t.Errorf("Retrieved id %v from repo entity %v is not the same as input %v",
				repoEntity.GetId(), repoEntity, testEntity.entityId)
		}
		if repoEntity.GetNodeIp() != testEntity.nodeIp {
			t.Errorf("Retrieved node ip %v from repo entity %v is not the same as input %v",
				repoEntity.GetNodeIp(), repoEntity, testEntity.nodeIp)
		}
	}
}
