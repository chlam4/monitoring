package client

import (
	"github.com/chlam4/monitoring/pkg/metric"
)

// Monitor defines the monitoring interface
// Object that will fetch values for the given monitoring properties for all the entities in the repository
// by connecting to the target
type Monitor interface {
	GetMonitoringType() MONITORING_TYPE
	Monitor(target *MonitorTarget) error
}

// MonitorTarget defines
type MonitorTarget struct {
	targetId        string
	config          interface{}
	Repository      metric.Repository // metric repository to store the metric values
	MonitoringProps metric.MonitoringProps   // meta data that defines what metrics to collect for what entities
}

// MakeMonitorTarget creates a monitor target given a repository and the metric defs
func MakeMonitorTarget(repo metric.Repository, metricDefs []metric.MetricDef) MonitorTarget {

	monitoringProps := metric.MakeMonitoringProps(repo, metricDefs)
	return MonitorTarget{Repository: repo, MonitoringProps: monitoringProps}
}
