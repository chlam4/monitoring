package metric


// MonitoringProps defines a set of metrics targeted to monitor for each entity
type MonitoringProps map[EntityId][]MetricDef

// MakeMonitoringProps creates a set of monitoring properties given a repository and the metric defs
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

// ByMetricDef rearranges MonitoringProps by MetricDef, with value being a map of NodeIp to EntityId
func (byEntityId MonitoringProps) ByMetricDef(repo Repository) map[MetricDef]map[NodeIp]EntityId {
	byMetricDef := map[MetricDef]map[NodeIp]EntityId{}
	for entityId, metricDefs := range byEntityId {
		for _, metricDef := range metricDefs {
			ip2IdMap, exists := byMetricDef[metricDef]
			if !exists {
				ip2IdMap = map[NodeIp]EntityId{}
			}
			repoEntity := repo.GetEntity(metricDef.EntityType, entityId)
			ip2IdMap[repoEntity.GetNodeIp()] = entityId
			byMetricDef[metricDef] = ip2IdMap
		}
	}
	return byMetricDef
}
