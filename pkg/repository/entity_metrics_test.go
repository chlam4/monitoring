package repository

import (
	"github.com/turbonomic/turbo-go-monitoring/pkg/model"
	"github.com/turbonomic/turbo-go-monitoring/pkg/model/property"
	"github.com/turbonomic/turbo-go-monitoring/pkg/model/resource"
	"testing"
)

var TestMetrics = []struct {
	resourceType model.ResourceType
	propType     model.MetricPropType
	value        float64
}{
	{resource.CPU, property.USED, 10.1},
	{resource.CPU, property.PEAK, 46.7},
	{resource.MEM, property.USED, 62.3},
	{resource.MEM, property.AVERAGE, 43.4},
	{resource.MEM_PROV, property.CAP, 87.9},
	{resource.DISK, property.CAP, 100.0},
	{resource.CPU_PROV, property.CAP, 90.5},
}

func TestMetricMap(t *testing.T) {

	metricMap := &EntityMetricMap{}

	// Add all test metrics into the metric map
	//
	for _, testMetric := range TestMetrics {
		metricMap.SetMetricValue(testMetric.resourceType, testMetric.propType, model.MetricValue(testMetric.value))
	}
	//
	// Retrieve the value for each metric and confirm it's the same as entered
	//
	for _, testMetric := range TestMetrics {
		value, err := metricMap.GetMetricValue(testMetric.resourceType, testMetric.propType)
		if err != nil {
			t.Errorf("Error while retrieving metric (%v, %v) from map %v: %s",
				testMetric.resourceType, testMetric.propType, metricMap, err)
		}
		if value != model.MetricValue(testMetric.value) {
			t.Errorf("Retrieved value %v of metric (%v, %v) from metric map %v is not the same as entered %v",
				value, testMetric.resourceType, testMetric.propType, metricMap, testMetric.value)

		}
	}
	//
	// Try to fetch the value of a non-existing metric
	//
	fakeResourceType := model.ResourceType("NOT_EXIST_RESOURCE_TYPE")
	fakeMetricPropType := model.MetricPropType("NOT_EXIST_PROP_TYPE")
	_, err := metricMap.GetMetricValue(fakeResourceType, fakeMetricPropType)
	if err == nil {
		t.Errorf("Expecting error but getting no error, when retrieving metric value of resource %v and "+
			"property %v from metric map %v", fakeResourceType, fakeMetricPropType, metricMap)
	}
}
