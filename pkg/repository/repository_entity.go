package repository

import (
	"github.com/chlam4/monitoring/pkg/model"
	"github.com/chlam4/monitoring/pkg/model/metric"
)

// RepositoryEntity defines a set of interfaces to access a repository entity in the metric Repository.
// It has per-entity info such as id, type, and node ip, and interfaces to get and set metric values.
type RepositoryEntity interface {
	GetId() model.EntityId
	GetType() model.EntityType
	GetNodeIp() model.NodeIp
	GetAllMetrics() EntityMetricMap
	GetMetricValue(metricKey EntityMetricKey) (metric.MetricValue, error)
	SetMetricValue(metricKey EntityMetricKey, value metric.MetricValue)
}
