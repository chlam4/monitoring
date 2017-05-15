package template

import (
	"github.com/chlam4/monitoring/pkg/model"
	"github.com/chlam4/monitoring/pkg/repository"
)


// MonitoringProps defines a set of metrics targeted to monitor for each entity
type MonitoringProps map[model.EntityId]MonitoringTemplate

// MakeMonitoringProps creates a set of monitoring properties given a repository and the metric defs
func MakeMonitoringProps(repo repository.Repository, monitoringTemplate MonitoringTemplate) MonitoringProps {

	monitoringProps := MonitoringProps{}

	for _, metricMeta := range monitoringTemplate {
		entities := repo.GetEntityInstances(metricMeta.MetricKey.EntityType)
		for _, entity := range entities {
			entityId := model.EntityId(entity.GetId())
			monitoringProps[entityId] = append(monitoringProps[entityId], metricMeta)
		}
	}
	return monitoringProps
}

// ByMetricMeta rearranges MonitoringProps by MetricMeta, with value being a map of NodeIp to EntityId
func (byEntityId MonitoringProps) ByMetricMeta(repo repository.Repository) map[MetricMeta]map[model.NodeIp]model.EntityId {
	byMetricMeta := map[MetricMeta]map[model.NodeIp]model.EntityId{}
	for entityId, monTemplate := range byEntityId {
		for _, metricMeta := range monTemplate {
			ip2IdMap, exists := byMetricMeta[metricMeta]
			if !exists {
				ip2IdMap = map[model.NodeIp]model.EntityId{}
			}
			repoEntity := repo.GetEntity(entityId)
			ip2IdMap[repoEntity.GetNodeIp()] = entityId
			byMetricMeta[metricMeta] = ip2IdMap
		}
	}
	return byMetricMeta
}
