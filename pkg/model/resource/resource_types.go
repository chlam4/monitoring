// The resource package contains definitions within the context of resource
package resource

import "github.com/turbonomic/turbo-go-monitoring/pkg/model"

// List of resource types
const (
	CPU      model.ResourceType = "CPU"
	MEM      model.ResourceType = "MEM"
	DISK     model.ResourceType = "DISK"
	MEM_PROV model.ResourceType = "MEM_PROV"
	CPU_PROV model.ResourceType = "CPU_PROV"
	NETWORK  model.ResourceType = "NETWORK"
)
