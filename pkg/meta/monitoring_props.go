package meta

import (
	"github.com/chlam4/monitoring/pkg/model"
	"github.com/chlam4/monitoring/pkg/meta/setter"
	"github.com/chlam4/monitoring/pkg/repository"
)


// MonitoringProps defines a set of metrics targeted to monitor for each entity
type MonitoringProps map[model.EntityId][]setter.MetricDef

// MakeMonitoringProps creates a set of monitoring properties given a repository and the metric defs
func MakeMonitoringProps(repo repository.Repository, metricDefs []setter.MetricDef) MonitoringProps {

	monitoringProps := MonitoringProps{}

	for _, metricDef := range metricDefs {
		entities := repo.GetEntityInstances(metricDef.EntityType)
		for _, entity := range entities {
			entityId := model.EntityId(entity.GetId())
			monitoringProps[entityId] = append(monitoringProps[entityId], metricDef)
		}
	}
	return monitoringProps
}

// ByMetricDef rearranges MonitoringProps by MetricDef, with value being a map of NodeIp to EntityId
func (byEntityId MonitoringProps) ByMetricDef(repo repository.Repository) map[setter.MetricDef]map[model.NodeIp]model.EntityId {
	byMetricDef := map[setter.MetricDef]map[model.NodeIp]model.EntityId{}
	for entityId, metricDefs := range byEntityId {
		for _, metricDef := range metricDefs {
			ip2IdMap, exists := byMetricDef[metricDef]
			if !exists {
				ip2IdMap = map[model.NodeIp]model.EntityId{}
			}
			repoEntity := repo.GetEntity(entityId)
			ip2IdMap[repoEntity.GetNodeIp()] = entityId
			byMetricDef[metricDef] = ip2IdMap
		}
	}
	return byMetricDef
}
