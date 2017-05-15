// The entity package contains definitions within the context of entity.
package entity

import "github.com/chlam4/monitoring/pkg/model"

// List of entity types
const (
	NODE      model.EntityType = "Node"
	CONTAINER model.EntityType = "Pod"
	APP       model.EntityType = "App"
)
