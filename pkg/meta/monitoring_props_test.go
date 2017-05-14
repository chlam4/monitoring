package meta

import (
	"testing"
	"github.com/davecgh/go-spew/spew"
	"github.com/chlam4/monitoring/pkg/repository"
	"github.com/chlam4/monitoring/pkg/model/entity"
	"github.com/chlam4/monitoring/pkg/model"
	"github.com/chlam4/monitoring/pkg/meta/setter"
	"github.com/chlam4/monitoring/pkg/repository/simpleRepo"
	"github.com/chlam4/monitoring/pkg/model/resource"
	"github.com/chlam4/monitoring/pkg/model/property"
)

var TestEntities = []struct {
	entityType model.EntityType
	entityId   model.EntityId
	nodeIp     model.NodeIp
}{
	{entity.NODE, "foo", "1.2.3.4"},
	{entity.NODE, "bar", "192.168.99.100"},
	{entity.CONTAINER, "123", "10.10.172.236"},
	{entity.APP, "xyz", "127.0.0.1"},
}

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

func MakeTestMetricDefs() []setter.MetricDef {
	metricDefs := []setter.MetricDef{}
	for _, testDef := range TestDefs {
		metricDef := setter.MakeMetricDefWithDefaultSetter(testDef.entityType, testDef.resourceType, testDef.propType)
		metricDefs = append(metricDefs, metricDef)
	}
	return metricDefs
}

func TestMonitoringProps(t *testing.T) {
	repo := MakeTestRepo()
	metricDefs := MakeTestMetricDefs()
	mProps := MakeMonitoringProps(repo, metricDefs)
	spew.Dump(mProps)
	byMetricDef := mProps.ByMetricDef(repo)
	spew.Dump(byMetricDef)
}

func MakeTestRepo() repository.Repository {
	//
	// Construct a list of repo entities based on the test data
	//
	repoEntities := []repository.RepositoryEntity{}
	for _, testEntity := range TestEntities {
		repoEntity := simpleRepo.NewSimpleMetricRepoEntity(testEntity.entityType, testEntity.entityId, testEntity.nodeIp)
		repoEntities = append(repoEntities, repoEntity)
	}
	//
	// Construct a repo and add those repo entities to the repo
	//
	repo := simpleRepo.NewSimpleMetricRepo()
	repo.SetEntityInstances(repoEntities)

	return repo
}