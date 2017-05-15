package template

import (
	"github.com/chlam4/monitoring/pkg/model"
	"github.com/chlam4/monitoring/pkg/model/entity"
	"github.com/chlam4/monitoring/pkg/repository"
	"github.com/chlam4/monitoring/pkg/repository/simpleRepo"
	"github.com/davecgh/go-spew/spew"
	"testing"
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

func MakeTestMonTemplates() []MonitoringTemplate {
	monTemplates := []MonitoringTemplate{}
	for _, testMonTemplate := range TestMonTemplates {
		metricDef := MakeMonitoringTemplateWithDefaultSetter(
			testMonTemplate.entityType, testMonTemplate.resourceType, testMonTemplate.propType)
		monTemplates = append(monTemplates, metricDef)
	}
	return monTemplates
}

func TestMonitoringProps(t *testing.T) {
	repo := MakeTestRepo()
	metricDefs := MakeTestMonTemplates()
	mProps := MakeMonitoringProps(repo, metricDefs)
	spew.Dump(mProps)
	byMetricDef := mProps.ByMonTemplate(repo)
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
