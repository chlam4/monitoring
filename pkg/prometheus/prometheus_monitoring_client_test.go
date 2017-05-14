package prometheus

import (
	"github.com/chlam4/monitoring/pkg/client"
	"github.com/chlam4/monitoring/pkg/meta/setter"
	"github.com/chlam4/monitoring/pkg/model/entity"
	"github.com/chlam4/monitoring/pkg/model/property"
	"github.com/chlam4/monitoring/pkg/model/resource"
	"github.com/chlam4/monitoring/pkg/repository"
	"github.com/chlam4/monitoring/pkg/repository/simpleRepo"
	"testing"
)

func TestPrometheusMonitor(t *testing.T) {
	//
	// What metrics do you want Prometheus to collect?
	//
	metricDefs := []setter.MetricDef{
		setter.MakeMetricDefWithDefaultSetter(entity.NODE, resource.CPU, property.USED),
		setter.MakeMetricDefWithDefaultSetter(entity.NODE, resource.MEM, property.USED),
		setter.MakeMetricDefWithDefaultSetter(entity.NODE, resource.MEM, property.CAP),
		setter.MakeMetricDefWithDefaultSetter(entity.NODE, resource.MEM, property.AVERAGE),
		setter.MakeMetricDefWithDefaultSetter(entity.NODE, resource.MEM, property.PEAK),
		setter.MakeMetricDefWithDefaultSetter(entity.NODE, resource.NETWORK, property.USED),
	}
	//
	// What entities do you want Prometheus to monitor?
	//
	repoEntities := []repository.RepositoryEntity{
		simpleRepo.NewSimpleMetricRepoEntity(entity.NODE, "abc", "192.168.99.100"),
		simpleRepo.NewSimpleMetricRepoEntity(entity.NODE, "xyz", "10.10.172.235"),
	}
	repo := simpleRepo.NewSimpleMetricRepo()
	repo.SetEntityInstances(repoEntities)
	//
	// Construct the monitor target
	//
	monitorTarget := client.MakeMonitorTarget(repo, metricDefs)
	//
	// Call Prometheus to collect metrics
	//
	promeServerUrl := "http://192.168.99.100:30900"
	promMonitor, err := NewPrometheusMonitor(promeServerUrl)
	if err != nil {
		t.Errorf("Error instantiating a Prometheus Monitor instance: %s", err)
	}
	promMonitor.Monitor(&monitorTarget)
	//
	// Process the collected metrics
	//
	for _, repoEntity := range repo.GetEntityInstances(entity.NODE) {
		t.Logf("Metrics collected for (%v, %v) are as follows:\n %s", repoEntity.GetType(), repoEntity.GetId(), repoEntity.GetAllMetrics())
	}
}
