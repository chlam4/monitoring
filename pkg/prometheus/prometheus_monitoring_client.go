package prometheus

import (
	"context"
	"github.com/chlam4/monitoring/pkg/client"
	"github.com/chlam4/monitoring/pkg/model"
	"github.com/chlam4/monitoring/pkg/repository"
	"github.com/golang/glog"
	prometheusHttpClient "github.com/prometheus/client_golang/api"
	prometheus "github.com/prometheus/client_golang/api/prometheus/v1"
	prometheusModel "github.com/prometheus/common/model"
	"time"
)

// PrometheusMonitor is a monitoring client that talks to Prometheus server and collect metrics as defined.
type PrometheusMonitor struct {
	PrometheusApi prometheus.API
}

// NewPrometheusMonitor() returns a PrometheusMonitor instance
func NewPrometheusMonitor(address string) (*PrometheusMonitor, error) {
	//
	// Create a Prometheus client
	//
	promConfig := prometheusHttpClient.Config{Address: address}
	promHttpClient, err := prometheusHttpClient.NewClient(promConfig)
	if err != nil {
		return nil, err
	}
	//
	// Instantiate a PrometheusMonitor object with the Prometheus client
	return &PrometheusMonitor{
		PrometheusApi: prometheus.NewAPI(promHttpClient),
	}, nil
}

func (monitor *PrometheusMonitor) GetSourceName() client.MONITORING_TYPE {
	return client.PROMETHEUS
}

func (monitor *PrometheusMonitor) Monitor(target *client.MonitorTarget) error {

	repo := target.Repository                             // metric repository
	monProps := target.MonitoringProps.ByMetricMeta(repo) // template to drive metric collection
	for metricMeta := range monProps {
		//
		// Locate the Prometheus query corresponding to the metric definition
		//
		key := metricMeta.MetricKey
		query, exists := MetricQueryMap[key]
		if !exists {
			glog.Warningf("Unsupported metric query: metric key = %v, supported queries = %v", key, MetricQueryMap)
			continue
		}
		//
		// Send a query to Prometheus for each required metric
		//
		value, err := monitor.PrometheusApi.Query(context.Background(), string(query.queryString), time.Now())
		if err != nil {
			glog.Errorf("Error querying Prometheus with query %v: %s", query, err)
			continue
		}
		glog.V(3).Infof("Metric sample collected for %v with query %v: %v", metricMeta, query, value)
		switch value.Type() {
		case prometheusModel.ValVector:
			for _, sample := range value.(prometheusModel.Vector) {
				entityId, err := query.entityId(sample)
				if err != nil {
					glog.Errorf("No entity id is found in Prometheus metric sample %v: %s", sample, err)
					continue
				}
				entityMetricKey := repository.EntityMetricKey{
					ResourceType: metricMeta.MetricKey.ResourceType,
					PropType:     metricMeta.MetricKey.PropType,
				}
				repo.SetMetricValue(entityId, entityMetricKey, model.MetricValue(sample.Value))
			}
		default:
			glog.Warningf("Unexpected/unsupported data type returned from Prometheus query %v: %v", query, value)
		}
	}
	return nil
}
