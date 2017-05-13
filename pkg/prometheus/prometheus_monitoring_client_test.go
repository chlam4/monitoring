package prometheus

import (
	"github.com/chlam4/monitoring/pkg/client"
	"github.com/chlam4/monitoring/pkg/metric"
	"testing"
)

func TestPrometheusMonitor(t *testing.T) {
	//
	// What metrics do you want Prometheus to collect?
	//
	metricDefs := []metric.MetricDef{
		metric.MakeMetricDefWithDefaultSetter(metric.NODE, metric.CPU, metric.USED),
		metric.MakeMetricDefWithDefaultSetter(metric.NODE, metric.MEM, metric.USED),
		metric.MakeMetricDefWithDefaultSetter(metric.NODE, metric.MEM, metric.CAP),
		metric.MakeMetricDefWithDefaultSetter(metric.NODE, metric.MEM, metric.AVERAGE),
		metric.MakeMetricDefWithDefaultSetter(metric.NODE, metric.MEM, metric.PEAK),
		metric.MakeMetricDefWithDefaultSetter(metric.NODE, metric.NETWORK, metric.USED),
	}
	//
	// What entities do you want Prometheus to monitor?
	//
	repoEntities := []metric.RepositoryEntity{
		metric.NewSimpleMetricRepoEntity(metric.NODE, "abc", "192.168.99.100"),
		metric.NewSimpleMetricRepoEntity(metric.NODE, "xyz", "10.10.172.235"),
	}
	repo := metric.NewSimpleMetricRepo()
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
	for _, repoEntity := range repo.GetEntityInstances(metric.NODE) {
		t.Logf("Metrics collected for (%v, %v) are as follows:\n %s", repoEntity.GetType(), repoEntity.GetId(), repoEntity.GetResourceMetrics())
	}
}
