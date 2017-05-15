package client

// MONITORING_TYPE defines the data type of the type of monitoring
type MONITORING_TYPE string

// List of monitoring types
const (
	DEFAULT_MESOS MONITORING_TYPE = "DEFAULT_MESOS"
	PROMETHEUS    MONITORING_TYPE = "PROMETHEUS"
)
