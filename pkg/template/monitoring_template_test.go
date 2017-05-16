package template

import (
	"github.com/turbonomic/turbo-go-monitoring/pkg/model"
	"github.com/turbonomic/turbo-go-monitoring/pkg/model/entity"
	"github.com/turbonomic/turbo-go-monitoring/pkg/model/property"
	"github.com/turbonomic/turbo-go-monitoring/pkg/model/resource"
	"testing"
)

var TestMetricDefs = []struct {
	entityType   model.EntityType
	resourceType model.ResourceType
	propType     model.MetricPropType
}{
	{entity.NODE, resource.CPU, property.USED},
	{entity.POD, resource.CPU, property.PEAK},
	{entity.APP, resource.MEM, property.USED},
	{entity.NODE, resource.MEM, property.AVERAGE},
	{entity.NODE, resource.MEM_PROV, property.CAP},
	{entity.NODE, resource.DISK, property.CAP},
	{entity.APP, resource.CPU_PROV, property.CAP},
}

func TestMakeMetricMetaWithDefaultSetter(t *testing.T) {
	for _, testMetricDef := range TestMetricDefs {
		metricMeta := MakeMetricMetaWithDefaultSetter(testMetricDef.entityType, testMetricDef.resourceType, testMetricDef.propType)
		if metricMeta.MetricKey.EntityType != testMetricDef.entityType {
			t.Errorf("Entity type in the metric def %v does not match with input %v", metricMeta, testMetricDef.entityType)
		}
		if metricMeta.MetricKey.ResourceType != testMetricDef.resourceType {
			t.Errorf("Resource type in the metric def %v does not match with input %v", metricMeta, testMetricDef.resourceType)
		}
		if metricMeta.MetricKey.PropType != testMetricDef.propType {
			t.Errorf("Property type in the metric def %v does not match with input %v", metricMeta, testMetricDef.propType)
		}
		_, ok := metricMeta.MetricSetter.(DefaultMetricSetter)
		if !ok {
			t.Errorf("Setter in metric def %v is not the default metric setter", metricMeta)
		}
	}
}
