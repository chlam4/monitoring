package metric

import "testing"

var TestDefs = []struct {
	entityType   EntityType
	resourceType ResourceType
	propType     MetricPropType
}{
	{NODE, CPU, USED},
	{CONTAINER, CPU, PEAK},
	{APP, MEM, USED},
	{NODE, MEM, AVERAGE},
	{NODE, MEM_PROV, CAP},
	{NODE, DISK, CAP},
	{APP, CPU_PROV, CAP},
}

func TestMakeMetricDefWithDefaultSetter(t *testing.T) {
	for _, testDef := range TestDefs {
		metricDef := MakeMetricDefWithDefaultSetter(testDef.entityType, testDef.resourceType, testDef.propType)
		if metricDef.EntityType != testDef.entityType {
			t.Errorf("Entity type in the metric def %v does not match with input %v", metricDef, testDef.entityType)
		}
		if metricDef.resourceType != testDef.resourceType {
			t.Errorf("Resource type in the metric def %v does not match with input %v", metricDef, testDef.resourceType)
		}
		if metricDef.propType != testDef.propType {
			t.Errorf("Property type in the metric def %v does not match with input %v", metricDef, testDef.propType)
		}
		_, ok := metricDef.metricSetter.(DefaultMetricSetter)
		if !ok {
			t.Errorf("Setter in metric def %v is not the default metric setter", metricDef)
		}
	}
}
