package template

import (
	"github.com/turbonomic/turbo-go-monitoring/pkg/model"
	"github.com/turbonomic/turbo-go-monitoring/pkg/model/entity"
	"github.com/turbonomic/turbo-go-monitoring/pkg/repository"
	"github.com/turbonomic/turbo-go-monitoring/pkg/repository/simpleRepo"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

var TestEntities = []struct {
	entityType model.EntityType
	entityId   model.EntityId
}{
	{entity.NODE, "1.2.3.4"},
	{entity.NODE, "192.168.99.100"},
	{entity.POD, "abc"},
	{entity.APP, "xyz"},
}

func MakeTestMonTemplate() MonitoringTemplate {
	monTemplate := MonitoringTemplate{}
	for _, testMetricMeta := range TestMetricDefs {
		metricMeta := MakeMetricMetaWithDefaultSetter(
			testMetricMeta.entityType, testMetricMeta.resourceType, testMetricMeta.propType)
		monTemplate = append(monTemplate, metricMeta)
	}
	return monTemplate
}

func TestMonitoringProps(t *testing.T) {
	repo := MakeTestRepo()
	metricDefs := MakeTestMonTemplate()
	mProps := MakeMonitoringProps(repo, metricDefs)
	spew.Dump(mProps)
	byMetricDef := mProps.ByMetricMeta(repo)
	spew.Dump(byMetricDef)
}

func MakeTestRepo() repository.Repository {
	//
	// Construct a list of repo entities based on the test data
	//
	repoEntities := []repository.RepositoryEntity{}
	for _, testEntity := range TestEntities {
		repoEntity := simpleRepo.NewSimpleMetricRepoEntity(testEntity.entityType, testEntity.entityId)
		repoEntities = append(repoEntities, repoEntity)
	}
	//
	// Construct a repo and add those repo entities to the repo
	//
	repo := simpleRepo.NewSimpleMetricRepo()
	repo.SetEntities(repoEntities)

	return repo
}
