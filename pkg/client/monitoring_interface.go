// The client package defines the monitoring interface any client would implement to leverage the metric repository
// and the monitoring template defined in this library.
package client

import (
	"github.com/chlam4/monitoring/pkg/template"
	"github.com/chlam4/monitoring/pkg/repository"
)

// Monitor defines the monitoring interface
type Monitor interface {
	// GetMonitoringType() returns the monitoring type
	GetMonitoringType() MONITORING_TYPE

	// Monitor() performs monitoring and collect metrics as defined
	Monitor(target *MonitorTarget) error
}

// MonitorTarget abstracts the arguments for the monitoring interface.
type MonitorTarget struct {
	targetId        string
	config          interface{}
	Repository      repository.Repository    // metric repository to store the metric values
	MonitoringProps template.MonitoringProps // meta data that defines what metrics to collect for what entities
}

// MakeMonitorTarget creates a monitor target based on the given a repository and the monitoring template
func MakeMonitorTarget(repo repository.Repository, monTemplate template.MonitoringTemplate) MonitorTarget {

	monitoringProps := template.MakeMonitoringProps(repo, monTemplate)
	return MonitorTarget{Repository: repo, MonitoringProps: monitoringProps}
}
