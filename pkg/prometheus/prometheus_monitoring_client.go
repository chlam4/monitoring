package prometheus

import (
	"context"
	"github.com/chlam4/monitoring/pkg/client"
	"github.com/golang/glog"
	prometheusHttpClient "github.com/prometheus/client_golang/api"
	prometheus "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"time"
	"strings"
	"github.com/chlam4/monitoring/pkg/metric"
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

func (monitor *PrometheusMonitor) Monitor(target *client.MonitorTarget) error {
	//
	// metric repository to fill
	//
	repo := target.Repository
	//
	// Send a query to Prometheus for each required metric
	//
	monProps := target.MonitoringProps.ByMetricDef(repo)
	for metricDef, ip2IdMap := range monProps {
		metricKey := metricDef.ToMetricKey()
		query, exists := MetricQueryMap[metricKey]
		if !exists {
			glog.Warningf("Unsupported metric query: metric key = %v, supported queries = %v", metricKey, MetricQueryMap)
			continue
		}
		value, err := monitor.PrometheusApi.Query(context.Background(), string(query), time.Now())
		if err != nil {
			glog.Errorf("Error querying Prometheus with query %v: %s", query, err)
			continue
		}
		switch value.Type() {
		case model.ValVector:
			for _, sample := range value.(model.Vector) {
				glog.V(3).Infof("Metric sample collected: %v", sample)
				instanceName := string(sample.Metric["instance"])
				nodeIp := strings.Split(instanceName, ":")[0]
				entityId, exists := ip2IdMap[metric.NodeIp(nodeIp)]
				if !exists {
					glog.Warningf("No entity found for IP %v in metric sample %v", nodeIp, sample)
				} else {
					repo.SetMetricValue(entityId, metricDef, metric.MetricValue(sample.Value))
				}
			}
		default:
			glog.Warningf("Unexpected/unsupported data type returned from Prometheus query %v: %v", query, value)
		}
	}
	return nil
}
