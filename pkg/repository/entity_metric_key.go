package repository

import "github.com/chlam4/monitoring/pkg/model"

type EntityMetricKey struct {
	ResourceType model.ResourceType
	PropType     model.MetricPropType
}
