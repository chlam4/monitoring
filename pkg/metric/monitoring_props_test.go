package metric

import (
	"testing"
	"fmt"
)

func TestMonitoringProps(t *testing.T) {
	repo := MakeTestRepo()
	metricDefs := MakeTestMetricDefs()
	mProps := MakeMonitoringProps(repo, metricDefs)
	fmt.Println(mProps)
	byMetricDef := mProps.ByMetricDef()
	fmt.Println(byMetricDef)
}