package template

import (
	"github.com/chlam4/monitoring/pkg/model"
	"github.com/chlam4/monitoring/pkg/model/metric"
	"github.com/chlam4/monitoring/pkg/repository"
)

// MonitoringTemplate defines a set of metric meta data to drive monitoring
type MonitoringTemplate []MetricMeta

// MetricMeta is the meta data of a metric, including the key of the metric and a metric setter
type MetricMeta struct {
	MetricKey    metric.MetricKey
	MetricSetter MetricSetter // Setter for the property
}

// The MetricSetter interface defines what a metric setter does -
// it defines how the input metric value is processed before setting the corresponding value in the repo entity.
type MetricSetter interface {
	SetMetricValue(entity repository.RepositoryEntity, key repository.EntityMetricKey, value metric.MetricValue)
}

// DefaultMetricSetter is a default implementation of a MetricSetter that just sets the value
// with the given key in the repo entity
type DefaultMetricSetter struct{}

func (setter DefaultMetricSetter) SetMetricValue(
	repoEntity repository.RepositoryEntity,
	key repository.EntityMetricKey,
	value metric.MetricValue,
) {
	repoEntity.SetMetricValue(key, value)
}

// MakeMetricMetaWithDefaultSetter makes a MetricMeta with given entity type, resource type and metric
// property type, and the default metric setter.
func MakeMetricMetaWithDefaultSetter(
	entityType model.EntityType,
	resourceType model.ResourceType,
	propType model.MetricPropType,
) MetricMeta {
	metricKey := metric.MetricKey{EntityType: entityType, ResourceType: resourceType, PropType: propType}
	setter := DefaultMetricSetter{}
	return MetricMeta{MetricKey: metricKey, MetricSetter: setter}
}
