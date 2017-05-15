package template

import (
	"github.com/chlam4/monitoring/pkg/model"
	"github.com/chlam4/monitoring/pkg/model/entity"
	"github.com/chlam4/monitoring/pkg/model/property"
	"github.com/chlam4/monitoring/pkg/model/resource"
	"testing"
)

var TestMonTemplates = []struct {
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
	for _, testMonTemplate := range TestMonTemplates {
		monTemplate := MakeMonitoringTemplateWithDefaultSetter(testMonTemplate.entityType, testMonTemplate.resourceType, testMonTemplate.propType)
		if monTemplate.MetricKey.EntityType != testMonTemplate.entityType {
			t.Errorf("Entity type in the metric def %v does not match with input %v", monTemplate, testMonTemplate.entityType)
		}
		if monTemplate.MetricKey.ResourceType != testMonTemplate.resourceType {
			t.Errorf("Resource type in the metric def %v does not match with input %v", monTemplate, testMonTemplate.resourceType)
		}
		if monTemplate.MetricKey.PropType != testMonTemplate.propType {
			t.Errorf("Property type in the metric def %v does not match with input %v", monTemplate, testMonTemplate.propType)
		}
		_, ok := monTemplate.MetricSetter.(DefaultMetricSetter)
		if !ok {
			t.Errorf("Setter in metric def %v is not the default metric setter", monTemplate)
		}
	}
}
