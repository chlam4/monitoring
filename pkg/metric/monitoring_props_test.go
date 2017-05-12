package metric

import (
	"testing"
	"github.com/davecgh/go-spew/spew"
)

func TestMonitoringProps(t *testing.T) {
	repo := MakeTestRepo()
	metricDefs := MakeTestMetricDefs()
	mProps := MakeMonitoringProps(repo, metricDefs)
	spew.Dump(mProps)
	byMetricDef := mProps.ByMetricDef(repo)
	spew.Dump(byMetricDef)
}