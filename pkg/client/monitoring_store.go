package client

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/chlam4/monitoring/pkg/metric"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

// =======================================================================
type MONITOR_NAME string

const (
	DEFAULT_MESOS    MONITOR_NAME = "DEFAULT_MESOS"
	PROMETHEUS_MESOS MONITOR_NAME = "PROMETHEUS_MESOS"
)

type ENTITY_ID string

// MonitoringProperty list for a MesosEntity
type EntityMonitoringProps struct {
	entityId string
	propMap  map[PropKey]*MonitoringProperty
}

// Key struct using resource and metric type to organize the monitoring property for an entity instance
type PropKey string

var (
	CPU_CAP       PropKey = NewPropKey(metric.CPU, metric.CAP)
	CPU_USED      PropKey = NewPropKey(metric.CPU, metric.USED)
	MEM_CAP       PropKey = NewPropKey(metric.MEM, metric.CAP)
	MEM_USED      PropKey = NewPropKey(metric.MEM, metric.USED)
	CPU_PROV_CAP  PropKey = NewPropKey(metric.CPU_PROV, metric.CAP)
	CPU_PROV_USED PropKey = NewPropKey(metric.CPU_PROV, metric.USED)
	MEM_PROV_CAP  PropKey = NewPropKey(metric.MEM_PROV, metric.CAP)
	MEM_PROV_USED PropKey = NewPropKey(metric.MEM_PROV, metric.USED)
)

func NewPropKey(resourceType metric.ResourceType, metricType metric.MetricPropType) PropKey {
	propKey := PropKey(fmt.Sprintf(string(resourceType) + "-" + string(metricType)))
	return propKey
}

// ====================================================================================================================
// Metadata for the metric to monitored
type MetricDef struct {
	entityType   metric.EntityType
	resourceType metric.ResourceType
	metricType   metric.MetricPropType
	metricSetter MetricSetter // Setter for the property
	// TODO: monitorSpec - spec used to poll for the property
}

//
type MonitoringProperty struct {
	metricDef *MetricDef
	id        string
}

// Object responsible for setting the value for a metric property
type MetricSetter interface {
	SetName(name string)
	SetMetricValue(entity metric.RepositoryEntity, value *float64)
}

// Object that will fetch values for the given monitoring properties for all the entities in the repository
// by connecting to the target
type Monitor interface {
	GetSourceName() MONITOR_NAME
	Monitor(target *MonitorTarget) (error)
}

type MonitorTarget struct {
	targetId        string
	config          interface{}
	repository      metric.Repository
	//rawStatsCache 	*RawStatsCache
	monitoringProps map[ENTITY_ID]*EntityMonitoringProps
}

// MetricStore is responsible for collecting values for different metrics for various resources belonging
// to the entities in the given probe repository.
// It is configured with a set of Monitors responsible for collecting the data values for the metrics.
// Applications invoke the GetMetrics() to trigger the data collection.
type MetricsMetadataStore interface {
	GetMetricDefs() map[metric.EntityType]map[metric.ResourceType]map[metric.MetricPropType]*MetricDef
}

// =========================== MetricSetter Implementation ============================================

type DefaultMetricSetter struct {
	entityType   metric.EntityType
	resourceType metric.ResourceType
	metricType   metric.MetricPropType
	name         string
}

func (setter *DefaultMetricSetter) SetMetricValue(entity metric.RepositoryEntity, value *float64) {
	//fmt.Printf("Setter : %s %+v\n", &setter, setter)
	if convertEntityType(setter.entityType) != entity.GetType() {
		glog.Errorf("Invalid entity type %s, required %s", entity.GetType(), setter.entityType)
	}
	var entityMetrics metric.MetricMap
	entityMetrics = entity.GetResourceMetrics()
	if entityMetrics == nil {
		glog.Errorf("Nil entity metrics for %s::%s", entity.GetType(), entity.GetId())
	}
	entityMetrics.SetResourceMetric(setter.resourceType, setter.metricType, value)
}


func (setter *DefaultMetricSetter) SetName(name string) {
	setter.name = name
}
// ============================== MetricStore Implementation =========================================

type MetricDefMap map[metric.EntityType]map[metric.ResourceType]map[metric.MetricPropType]*MetricDef

func NewMetricDefMap() MetricDefMap {
	return make(map[metric.EntityType]map[metric.ResourceType]map[metric.MetricPropType]*MetricDef)
}
func (mdm MetricDefMap) Put(et metric.EntityType, rt metric.ResourceType, mt metric.MetricPropType, md *MetricDef) {
	resourceMap, ok := mdm[et]
	if !ok {
		mdm[et] = make(map[metric.ResourceType]map[metric.MetricPropType]*MetricDef)
	}
	resourceMap = mdm[et]

	metricMap, ok := resourceMap[rt]
	if !ok {
		resourceMap[rt] = make(map[metric.MetricPropType]*MetricDef)
	}
	metricMap = resourceMap[rt]

	metricMap[mt] = md
}

func (mdm MetricDefMap) Get(et metric.EntityType, rt metric.ResourceType, mt metric.MetricPropType) *MetricDef {
	resourceMap, ok := mdm[et]
	if !ok {
		return nil
	}
	resourceMap = mdm[et]

	metricMap, ok := resourceMap[rt]
	if !ok {
		return nil
	}
	return metricMap[mt]
}


// Implementation for the MetricsStore for Mesos target
type MesosMetricsMetadataStore struct {
	metricDefMap map[metric.EntityType]map[metric.ResourceType]map[metric.MetricPropType]*MetricDef
}

func NewMesosMetricsMetadataStore() *MesosMetricsMetadataStore {
	mc := &MesosMetricsMetadataStore{}

	// TODO: parse from a file
	// read from a config file a table of entity type, resource type, metric type to be discovered
	var mdMap map[metric.EntityType]map[metric.ResourceType]map[metric.MetricPropType]*MetricDef
	mdMap = make(map[metric.EntityType]map[metric.ResourceType]map[metric.MetricPropType]*MetricDef)

	mdMap[metric.NODE] = make(map[metric.ResourceType]map[metric.MetricPropType]*MetricDef)
	resourceMap := mdMap[metric.NODE]
	addDefaultMetricDef(metric.NODE, metric.CPU, metric.CAP, resourceMap)
	addDefaultMetricDef(metric.NODE, metric.MEM, metric.CAP, resourceMap)
	addDefaultMetricDef(metric.NODE, metric.CPU, metric.USED, resourceMap)
	addDefaultMetricDef(metric.NODE, metric.MEM, metric.USED, resourceMap)
	addDefaultMetricDef(metric.NODE, metric.CPU_PROV, metric.CAP, resourceMap)
	addDefaultMetricDef(metric.NODE, metric.CPU_PROV, metric.USED, resourceMap)
	addDefaultMetricDef(metric.NODE, metric.MEM_PROV, metric.CAP, resourceMap)
	addDefaultMetricDef(metric.NODE, metric.MEM_PROV, metric.USED, resourceMap)

	mdMap[metric.CONTAINER] = make(map[metric.ResourceType]map[metric.MetricPropType]*MetricDef)
	resourceMap = mdMap[metric.CONTAINER]
	addDefaultMetricDef(metric.CONTAINER, metric.CPU, metric.CAP, resourceMap)
	addDefaultMetricDef(metric.CONTAINER, metric.CPU, metric.USED, resourceMap)
	addDefaultMetricDef(metric.CONTAINER, metric.MEM, metric.CAP, resourceMap)
	addDefaultMetricDef(metric.CONTAINER, metric.MEM, metric.USED, resourceMap)
	addDefaultMetricDef(metric.CONTAINER, metric.CPU_PROV, metric.CAP, resourceMap)
	addDefaultMetricDef(metric.CONTAINER, metric.CPU_PROV, metric.USED, resourceMap)
	addDefaultMetricDef(metric.CONTAINER, metric.MEM_PROV, metric.CAP, resourceMap)
	addDefaultMetricDef(metric.CONTAINER, metric.MEM_PROV, metric.USED, resourceMap)

	mdMap[metric.APP] = make(map[metric.ResourceType]map[metric.MetricPropType]*MetricDef)
	resourceMap = mdMap[metric.APP]
	addDefaultMetricDef(metric.APP, metric.CPU, metric.CAP, resourceMap)
	addDefaultMetricDef(metric.APP, metric.CPU, metric.USED, resourceMap)
	addDefaultMetricDef(metric.APP, metric.MEM, metric.CAP, resourceMap)
	addDefaultMetricDef(metric.APP, metric.MEM, metric.USED, resourceMap)
	addDefaultMetricDef(metric.APP, metric.CPU_PROV, metric.CAP, resourceMap)
	addDefaultMetricDef(metric.APP, metric.CPU_PROV, metric.USED, resourceMap)
	addDefaultMetricDef(metric.APP, metric.MEM_PROV, metric.CAP, resourceMap)
	addDefaultMetricDef(metric.APP, metric.MEM_PROV, metric.USED, resourceMap)

	mc.metricDefMap = mdMap

	return mc
}

func addDefaultMetricDef(entityType metric.EntityType, resourceType metric.ResourceType, metricType metric.MetricPropType,
	resourceMap map[metric.ResourceType]map[metric.MetricPropType]*MetricDef) {
	metricMap, ok := resourceMap[resourceType]
	if !ok {
		resourceMap[resourceType] = make(map[metric.MetricPropType]*MetricDef)
	}
	metricMap = resourceMap[resourceType]

	metricSetter := &DefaultMetricSetter{
		entityType:   entityType,
		resourceType: resourceType,
		metricType:   metricType,}
	metricSetter.SetName(fmt.Sprintf("%s:%s:%s:%s", &metricSetter, entityType, resourceType, metricType))

	metricDef := &MetricDef{
		entityType:   entityType,
		resourceType: resourceType,
		metricType:   metricType,
		metricSetter: metricSetter,
	}
	metricMap[metricType] = metricDef
}

func (metricStore *MesosMetricsMetadataStore) GetMetricDefs() map[metric.EntityType]map[metric.ResourceType]map[metric.MetricPropType]*MetricDef {
	return metricStore.metricDefMap
}

func createMonitoringProps(repository metric.Repository, mdMap map[metric.EntityType]map[metric.ResourceType]map[metric.MetricPropType]*MetricDef) map[ENTITY_ID]*EntityMonitoringProps {
	// Create monitoring property for the repository entities using the MetricDef configured for each entity type
	var entityPropsMap map[ENTITY_ID]*EntityMonitoringProps
	entityPropsMap = make(map[ENTITY_ID]*EntityMonitoringProps)

	for entityType, resourceMap := range mdMap {
		// entity instances
		entityList := repository.GetEntityInstances(convertEntityType(entityType))
		for _, entity := range entityList {
			entityId := entity.GetId()
			// monitoring properties of an entity instance
			entityProps := &EntityMonitoringProps{
				entityId: entityId,
				propMap:  make(map[PropKey]*MonitoringProperty),
			}
			entityPropsMap[ENTITY_ID(entityId)] = entityProps
			tempMap := entityProps.propMap
			for _, metricMap := range resourceMap {
				for _, metricDef := range metricMap {
					glog.V(4).Infof("MetricDef %s --->  %s::%s::%s\n", entityId, metricDef.entityType, metricDef.resourceType, metricDef.metricType)
					prop := &MonitoringProperty{
						metricDef: metricDef,
						id:        entityId,
					}
					tempMap[NewPropKey(metricDef.resourceType, metricDef.metricType)] = prop
				}
			}
		}
	}
	return entityPropsMap
}

func convertEntityType(entityType metric.EntityType) proto.EntityDTO_EntityType {
	if entityType == metric.NODE {
		return proto.EntityDTO_VIRTUAL_MACHINE
	} else if entityType == metric.APP {
		return proto.EntityDTO_APPLICATION
	} else if entityType == metric.CONTAINER {
		return proto.EntityDTO_CONTAINER
	}
	return proto.EntityDTO_UNKNOWN
}
