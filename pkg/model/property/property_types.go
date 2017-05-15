// The property package contains definitions within the context of metric property.
package property

import "github.com/chlam4/monitoring/pkg/model"

// List of metric property types
const (
	USED    model.MetricPropType = "Used"
	CAP     model.MetricPropType = "Capacity"
	PEAK    model.MetricPropType = "Peak"
	AVERAGE model.MetricPropType = "Average"
)
