package setter

import (
	"github.com/chlam4/monitoring/pkg/model/metric"
	"github.com/chlam4/monitoring/pkg/repository"
)

// The MetricSetter interface defines what a metric setter does -
// it sets the input value in the given metric repository entity
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
