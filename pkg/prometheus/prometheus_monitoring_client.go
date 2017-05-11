package prometheus


import (
	prometheusHttpClient "github.com/prometheus/client_golang/api"
	prometheus "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/chlam4/monitoring/pkg/client"
	"github.com/golang/glog"
	"context"
	"time"
	"github.com/prometheus/common/model"
)

type PrometheusMonitor struct {
	PrometheusApi prometheus.API
}

func NewPrometheusMonitor(address string) (*PrometheusMonitor, error) {
	//
	// Create a Prometheus client
	//
	promConfig := prometheusHttpClient.Config{Address: address}
	promHttpClient, err := prometheusHttpClient.NewClient(promConfig)
	if err != nil {
		return nil, err
	}

	return &PrometheusMonitor{
		PrometheusApi: prometheus.NewAPI(promHttpClient),
	}, nil
}

func (monitor *PrometheusMonitor) GetSourceName() client.MONITORING_TYPE {
	return client.PROMETHEUS
}

func (monitor *PrometheusMonitor) Monitor(target *client.MonitorTarget) (error) {
	monProps := target.MonitoringProps.ByMetricDef()

	for metricDef := range monProps {
		metricKey := metricDef.ToMetricKey()
		query, exists := MetricQueryMap[metricKey]
		if !exists {
			glog.Warningf("Unsupported metric query: metric key = %v, supported queries = %v", metricKey, MetricQueryMap)
			continue
		}
		value, err := monitor.PrometheusApi.Query(context.Background(), string(query), time.Now())
		if err != nil {
			glog.Errorf("Error querying Prometheus with query %v: %s", query, err)
		}
		for _, metric := range value.(model.Vector) {
			glog.Infof("Metric collected: %v", metric)
		}
	}
	return nil
}
