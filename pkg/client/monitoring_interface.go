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
	repository      metric.Repository // metric repository to store the metric values
	monitoringProps MonitoringProps   // meta data that defines what metrics to collect for what entities
}

// MonitoringProps defines a set of metrics targeted to monitor for each entity
type MonitoringProps map[metric.EntityId][]metric.MetricDef

// CreateMonitoringProps creates monitoring property for the repository entities using the MetricDef of each entity type
func CreateMonitoringProps(repo metric.Repository, metricDefs []metric.MetricDef) *MonitoringProps {

	monitoringProps := make(MonitoringProps)

	for _, metricDef := range metricDefs {
		entities := repo.GetEntityInstances(metricDef.EntityType)
		for _, entity := range entities {
			entityId := metric.EntityId(entity.GetId())
			monitoringProps[entityId] = append(monitoringProps[entityId], metricDef)
		}
	}
	return monitoringProps
}
