package template

import (
	"github.com/turbonomic/turbo-go-monitoring/pkg/model"
	"github.com/turbonomic/turbo-go-monitoring/pkg/repository"
)

// MonitoringProps defines a set of metrics targeted to monitor for each entity
type MonitoringProps map[model.EntityTypedId]MonitoringTemplate

// MakeMonitoringProps creates a set of monitoring properties given a repository and the metric defs
func MakeMonitoringProps(repo repository.Repository, monitoringTemplate MonitoringTemplate) MonitoringProps {

	monitoringProps := MonitoringProps{}

	for _, metricMeta := range monitoringTemplate {
		entities := repo.GetEntitiesByType(metricMeta.MetricKey.EntityType)
		for _, entity := range entities {
			entityTypedId := entity.GetTypedId()
			monitoringProps[entityTypedId] = append(monitoringProps[entityTypedId], metricMeta)
		}
	}
	return monitoringProps
}

// ByMetricMeta rearranges MonitoringProps by MetricMeta, with value being a list of EntityId's
func (byEntityTypedId MonitoringProps) ByMetricMeta(repo repository.Repository) map[MetricMeta][]model.EntityTypedId {
	byMetricMeta := map[MetricMeta][]model.EntityTypedId{}
	for entityTypedId, monTemplate := range byEntityTypedId {
		for _, metricMeta := range monTemplate {
			typedIds, exists := byMetricMeta[metricMeta]
			if !exists {
				typedIds = []model.EntityTypedId{}
			}
			typedIds = append(typedIds, entityTypedId)
			byMetricMeta[metricMeta] = typedIds
		}
	}
	return byMetricMeta
}
