package template

import (
	"github.com/chlam4/monitoring/pkg/model"
	"github.com/chlam4/monitoring/pkg/repository"
)


// MonitoringProps defines a set of metrics targeted to monitor for each entity
type MonitoringProps map[model.EntityId][]MonitoringTemplate

// MakeMonitoringProps creates a set of monitoring properties given a repository and the metric defs
func MakeMonitoringProps(repo repository.Repository, monitoringTemplates []MonitoringTemplate) MonitoringProps {

	monitoringProps := MonitoringProps{}

	for _, monitoringTemplate := range monitoringTemplates {
		entities := repo.GetEntityInstances(monitoringTemplate.MetricKey.EntityType)
		for _, entity := range entities {
			entityId := model.EntityId(entity.GetId())
			monitoringProps[entityId] = append(monitoringProps[entityId], monitoringTemplate)
		}
	}
	return monitoringProps
}

// ByMonTemplate rearranges MonitoringProps by MonitoringTemplate, with value being a map of NodeIp to EntityId
func (byEntityId MonitoringProps) ByMonTemplate(repo repository.Repository) map[MonitoringTemplate]map[model.NodeIp]model.EntityId {
	byMetricDef := map[MonitoringTemplate]map[model.NodeIp]model.EntityId{}
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
