package metric

import "testing"

func TestMetricMap(t *testing.T) {

	testMetrics := []struct {
		resourceType ResourceType
		propType     MetricPropType
		value        float64
	}{
		{CPU, USED, 10.1},
		{CPU_PROV, USED, 46.7},
		{CPU_PROV, CAP, 90.5},
		{MEM, USED, 43.4},
	}

	metricMap := &MetricMap{}

	// Add all test metrics into the metric map
	for _, metric := range testMetrics {
		metricMap.SetResourceMetric(metric.resourceType, metric.propType, metric.value)
	}

	// Retrieve the value for each metric and confirm it's the same as entered
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
}
