package client

import (
	"github.com/chlam4/monitoring/pkg/meta"
	"github.com/chlam4/monitoring/pkg/meta/setter"
	"github.com/chlam4/monitoring/pkg/repository"
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
	Repository      repository.Repository // metric repository to store the metric values
	MonitoringProps meta.MonitoringProps  // meta data that defines what metrics to collect for what entities
}

// MakeMonitorTarget creates a monitor target given a repository and the metric defs
func MakeMonitorTarget(repo repository.Repository, metricDefs []setter.MetricDef) MonitorTarget {

	monitoringProps := meta.MakeMonitoringProps(repo, metricDefs)
	return MonitorTarget{Repository: repo, MonitoringProps: monitoringProps}
}
