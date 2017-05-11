package metric


// MonitoringProps defines a set of metrics targeted to monitor for each entity
type MonitoringProps map[EntityId][]MetricDef

// MakeMonitorTarget creates a monitor target given a repository and the metric defs
func MakeMonitoringProps(repo Repository, metricDefs []MetricDef) MonitoringProps {

	monitoringProps := make(MonitoringProps)

	for _, metricDef := range metricDefs {
		entities := repo.GetEntityInstances(metricDef.EntityType)
		for _, entity := range entities {
			entityId := EntityId(entity.GetId())
			monitoringProps[entityId] = append(monitoringProps[entityId], metricDef)
		}
	}
	return monitoringProps
}

// ByMetricDef rearranges MonitoringProps by MetricDef
func (byEntityId MonitoringProps) ByMetricDef() map[MetricDef][]EntityId {
	byMetricDef := map[MetricDef][]EntityId{}
	for entityId, metricDefs := range byEntityId {
		for _, metricDef := range metricDefs {
			ids, exists := byMetricDef[metricDef]
			if !exists {
				ids = []EntityId{}
			}
			ids = append(ids, entityId)
			byMetricDef[metricDef] = ids
		}
	}
	return byMetricDef
}