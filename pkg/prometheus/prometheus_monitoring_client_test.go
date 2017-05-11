package prometheus

import (
	"github.com/chlam4/monitoring/pkg/metric"
	"github.com/chlam4/monitoring/pkg/client"
	"testing"
)

func TestPrometheusMonitor_Monitor(t *testing.T) {
	//
	// What metrics do you want Prometheus to collect?
	//
	metricDefs := []metric.MetricDef{
		metric.MakeMetricDefWithDefaultSetter(metric.NODE, metric.MEM, metric.USED),
	}
	//
	// What entities do you want Prometheus to monitor?
	//
	repoEntities := []metric.RepositoryEntity{
		metric.NewSimpleMetricRepoEntity(metric.NODE, "abc", "192.168.99.100"),
		metric.NewSimpleMetricRepoEntity(metric.NODE, "xyz", "localhost"),
	}
	repo := metric.NewSimpleMetricRepo()
	repo.SetEntityInstances(repoEntities)
	//
	// Construct the monitor target
	//
	monitorTarget := client.MakeMonitorTarget(repo, metricDefs)

	promeServerUrl := "http://localhost:9090"
	promMonitor, err := NewPrometheusMonitor(promeServerUrl)
	if err != nil {
		t.Errorf("Error instantiating a Prometheus Monitor instance: %s", err)
	}
	promMonitor.Monitor(&monitorTarget)
	for _, repoEntity := range repoEntities {
		t.Logf("Metrics collected for (%v, %v) are as follows:\n", repoEntity.GetType(), repoEntity.GetId())
		repoEntity.GetResourceMetrics().PrintMetrics()
	}
}
