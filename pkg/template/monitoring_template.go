package template

import (
	"github.com/turbonomic/turbo-go-monitoring/pkg/model"
	"github.com/turbonomic/turbo-go-monitoring/pkg/repository"
)

// MonitoringTemplate defines a set of metric meta data to drive monitoring
type MonitoringTemplate []MetricMeta

// MetricMeta is the meta data of a metric, including the key of the metric and a metric setter
type MetricMeta struct {
	MetricKey    model.MetricKey
	MetricSetter MetricSetter // Setter for the property
}

// The MetricSetter interface defines what a metric setter does -
// it defines how the input metric value is processed before setting the corresponding value in the repo entity.
type MetricSetter interface {
	SetMetricValue(entity repository.RepositoryEntity, key repository.EntityMetricKey, value model.MetricValue)
}

// DefaultMetricSetter is a default implementation of a MetricSetter that just sets the value
// with the given key in the repo entity
type DefaultMetricSetter struct{}

func (setter DefaultMetricSetter) SetMetricValue(
	repoEntity repository.RepositoryEntity,
	key repository.EntityMetricKey,
	value model.MetricValue,
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
	metricKey := model.MetricKey{EntityType: entityType, ResourceType: resourceType, PropType: propType}
	setter := DefaultMetricSetter{}
	return MetricMeta{MetricKey: metricKey, MetricSetter: setter}
}
