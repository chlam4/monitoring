package metric

import "github.com/chlam4/monitoring/pkg/model"

type MetricKey struct {
	ResourceType model.ResourceType
	PropType     model.MetricPropType
}
