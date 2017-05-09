package metric

import (
	"github.com/golang/glog"
)

// The MetricSetter interface defines what a metric setter does -
// it sets the input value in the given metric repository entity
type MetricSetter interface {
	SetName(name string)
	SetMetricValue(entity RepositoryEntity, value MetricValue)
}

// DefaultMetricSetter is a default implementation of a MetricSetter
type DefaultMetricSetter struct {
	entityType   EntityType
	resourceType ResourceType
	propType     MetricPropType
	name         string
}

func (setter DefaultMetricSetter) SetMetricValue(entity RepositoryEntity, value MetricValue) {
	if setter.entityType != entity.GetType() {
		glog.Errorf("Invalid entity type %s, required %s", entity.GetType(), setter.entityType)
	}
	entityMetrics := entity.GetResourceMetrics()
	if entityMetrics == nil {
		glog.Errorf("Nil entity metrics for %s::%s", entity.GetType(), entity.GetId())
	}
	entityMetrics.SetResourceMetric(setter.resourceType, setter.propType, value)
}

func (setter DefaultMetricSetter) SetName(name string) {
	setter.name = name
}
