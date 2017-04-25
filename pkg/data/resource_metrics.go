package data

const (
	CPU_MULTIPLIER float64 = 2000
	KB_MULTIPLIER float64 = 1024
	DEFAULT_VAL float64 = 0.0
)

type MetricPropType string

const (
	USED    MetricPropType = "Used"
	CAP     MetricPropType = "Capacity"
	PEAK    MetricPropType = "Peak"
	AVERAGE MetricPropType = "Average"
)

type ResourceType string

const (
	CPU      ResourceType = "CPU"
	MEM      ResourceType = "MEM"
	DISK     ResourceType = "DISK"
	MEM_PROV ResourceType = "MEM_PROV"
	CPU_PROV ResourceType = "CPU_PROV"
)

type EntityType string

const (
	NODE EntityType = "Node"
	CONTAINER  EntityType = "Pod"
	APP  EntityType = "App"
)
