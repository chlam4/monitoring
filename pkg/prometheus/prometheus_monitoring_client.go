package prometheus


import (
	prometheusHttpClient "github.com/prometheus/client_golang/api"
	prometheus "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/chlam4/monitoring/pkg/client"
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
	return client.PROMETHEUS_MESOS
}

func (monitor *PrometheusMonitor) Monitor(target *client.MonitorTarget) (error) {
	return nil
}
