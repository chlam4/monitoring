package metric

import "testing"

func TestMetricMap(t *testing.T) {

	testMetrics := []struct {
		resourceType ResourceType
		propType     MetricPropType
		value        float64
	}{
		{CPU, USED, 10.1},
		{CPU, PEAK, 46.7},
		{MEM, USED, 62.3},
		{MEM, AVERAGE, 43.4},
		{MEM_PROV, CAP, 87.9},
		{DISK, CAP, 100.0},
		{CPU_PROV, CAP, 90.5},
	}

	metricMap := &MetricMap{}

	// Add all test metrics into the metric map
	//
	for _, metric := range testMetrics {
		metricMap.SetResourceMetric(metric.resourceType, metric.propType, MetricValue(metric.value))
	}
	//
	// Retrieve the value for each metric and confirm it's the same as entered
	//
	for _, metric := range testMetrics {
		value, err := metricMap.GetResourceMetric(metric.resourceType, metric.propType)
		if err != nil {
			t.Errorf("Error while retrieving metric (%v, %v) from map %v: %s",
				metric.resourceType, metric.propType, metricMap, err)
		}
		if value != MetricValue(metric.value) {
			t.Errorf("Retrieved value %v of metric (%v, %v) from metric map %v is not the same as entered %v",
				value, metric.resourceType, metric.propType, metricMap, metric.value)

		}
	}
	//
	// Try to fetch the value of a non-existing metric
	//
	fakeResourceType := ResourceType("NOT_EXIST_RESOURCE_TYPE")
	fakeMetricPropType := MetricPropType("NOT_EXIST_PROP_TYPE")
	_, err := metricMap.GetResourceMetric(fakeResourceType, fakeMetricPropType)
	if err == nil {
		t.Errorf("Expecting error but getting no error, when retrieving metric value of resource %v and "+
			"property %v from metric map %v", fakeResourceType, fakeMetricPropType, metricMap)
	}
}
