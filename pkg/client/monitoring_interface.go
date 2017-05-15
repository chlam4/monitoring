package client

import (
	"github.com/chlam4/monitoring/pkg/template"
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
	Repository      repository.Repository    // metric repository to store the metric values
	MonitoringProps template.MonitoringProps // meta data that defines what metrics to collect for what entities
}

// MakeMonitorTarget creates a monitor target given a repository and the metric defs
func MakeMonitorTarget(repo repository.Repository, monTemplates []template.MonitoringTemplate) MonitorTarget {

	monitoringProps := template.MakeMonitoringProps(repo, monTemplates)
	return MonitorTarget{Repository: repo, MonitoringProps: monitoringProps}
}
