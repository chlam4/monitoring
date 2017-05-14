package setter

import (
	"github.com/chlam4/monitoring/pkg/model"
	"github.com/chlam4/monitoring/pkg/model/entity"
	"github.com/chlam4/monitoring/pkg/model/property"
	"github.com/chlam4/monitoring/pkg/model/resource"
	"testing"
)

var TestDefs = []struct {
	entityType   model.EntityType
	resourceType model.ResourceType
	propType     model.MetricPropType
}{
	{entity.NODE, resource.CPU, property.USED},
	{entity.CONTAINER, resource.CPU, property.PEAK},
	{entity.APP, resource.MEM, property.USED},
	{entity.NODE, resource.MEM, property.AVERAGE},
	{entity.NODE, resource.MEM_PROV, property.CAP},
	{entity.NODE, resource.DISK, property.CAP},
	{entity.APP, resource.CPU_PROV, property.CAP},
}

func TestMakeMetricDefWithDefaultSetter(t *testing.T) {
	for _, testDef := range TestDefs {
		metricDef := MakeMetricDefWithDefaultSetter(testDef.entityType, testDef.resourceType, testDef.propType)
		if metricDef.EntityType != testDef.entityType {
			t.Errorf("Entity type in the metric def %v does not match with input %v", metricDef, testDef.entityType)
		}
		if metricDef.ResourceType != testDef.resourceType {
			t.Errorf("Resource type in the metric def %v does not match with input %v", metricDef, testDef.resourceType)
		}
		if metricDef.PropType != testDef.propType {
			t.Errorf("Property type in the metric def %v does not match with input %v", metricDef, testDef.propType)
		}
		_, ok := metricDef.MetricSetter.(DefaultMetricSetter)
		if !ok {
			t.Errorf("Setter in metric def %v is not the default metric setter", metricDef)
		}
	}
}
