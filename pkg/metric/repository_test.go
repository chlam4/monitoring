package metric

import "testing"

var TestEntities = []struct {
	entityType EntityType
	entityId   EntityId
	nodeIp     NodeIp
}{
	{NODE, "foo", "1.2.3.4"},
	{NODE, "bar", "192.168.99.100"},
	{CONTAINER, "123", "10.10.172.236"},
	{APP, "xyz", "127.0.0.1"},
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

func TestSimpleMetricRepoEntity_GetSetMetricValue(t *testing.T) {
	//
	// Pick one set of test data to construct a repo entity
	//
	test0 := TestEntities[0]
	repoEntity := NewSimpleMetricRepoEntity(test0.entityType, test0.entityId, test0.nodeIp)
	//
	// Add all test metrics into the repository entity
	//
	for _, metric := range TestMetrics {
		repoEntity.SetMetricValue(metric.resourceType, metric.propType, MetricValue(metric.value))
	}
	//
	// Retrieve the value for each metric and confirm it's the same as entered
	//
	for _, metric := range TestMetrics {
		value, err := repoEntity.GetResourceMetric(metric.resourceType, metric.propType)
		if err != nil {
			t.Errorf("Error while retrieving metric (%v, %v) from repo entity %v: %s",
				metric.resourceType, metric.propType, repoEntity, err)
		}
		if value != MetricValue(metric.value) {
			t.Errorf("Retrieved value %v of metric (%v, %v) from repo entity %v is not the same as entered %v",
				value, metric.resourceType, metric.propType, repoEntity, metric.value)
		}
	}
}

func TestSimpleMetricRepo(t *testing.T) {
	//
	// Construct a list of repo entities based on the test data
	//
	repoEntities := []RepositoryEntity{}
	for _, testEntity := range TestEntities {
		repoEntity := NewSimpleMetricRepoEntity(testEntity.entityType, testEntity.entityId, testEntity.nodeIp)
		repoEntities = append(repoEntities, repoEntity)
	}
	//
	// Construct a repo and add those repo entities to the repo
	//
	repo := NewSimpleMetricRepo()
	repo.SetEntityInstances(repoEntities)
	//
	// Check GetEntity result
	//
	for _, testEntity := range TestEntities {
		repoEntity := repo.GetEntity(testEntity.entityType, testEntity.entityId)
		if repoEntity == nil {
			t.Errorf("No repo entity for type %v and id %v found in repo %v", testEntity.entityType, testEntity.entityId, repo)
		} else if repoEntity.GetType() != testEntity.entityType {
			t.Errorf("Retrieved type %v from repo %v for entity type %v and id %v is not the same as entered %v",
				repoEntity.GetType(), repo, testEntity.entityType, testEntity.entityId, testEntity.entityType)
		} else if repoEntity.GetId() != testEntity.entityId {
			t.Errorf("Retrieved id %v from repo %v for entity type %v and id %v is not the same as entered %v",
				repoEntity.GetId(), repo, testEntity.entityType, testEntity.entityId, testEntity.entityId)
		} else if repoEntity.GetNodeIp() != testEntity.nodeIp {
			t.Errorf("Retrieved node ip %v from repo %v for entity type %v and id %v is not the same as entered %v",
				repoEntity.GetNodeIp(), repo, testEntity.entityType, testEntity.entityId, testEntity.nodeIp)
		}
	}
}
