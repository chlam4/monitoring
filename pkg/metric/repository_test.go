package metric

import "testing"

func TestSimpleMetricRepoEntity_GetId_GetType(t *testing.T) {

	testIdsTypes := []struct {
		entityType EntityType
		entityId   EntityId
	}{
		{NODE, "abc"},
		{CONTAINER, "123"},
		{APP, "xyz"},
	}

	for _, idType := range testIdsTypes {
		repoEntity := NewSimpleMetricRepoEntity(idType.entityType, idType.entityId)
		if repoEntity.GetType() != idType.entityType {
			t.Errorf("Retrieved type %v from repo entity %v is not the same as input %v",
				repoEntity.GetType(), repoEntity, idType.entityType)
		}
	}
}

func TestSimpleMetricRepoEntity_GetSetMetricValue(t *testing.T) {

	repoEntity := NewSimpleMetricRepoEntity(NODE, "abc")

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
