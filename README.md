# turbo-go-monitoring

This is a GO monitoring library that defines a set of monitoring interfaces, including
[metric repository](pkg/repository/repository.go) and [monitoring template](pkg/template/monitoring_template.go)
interfaces.  Any target-specific monitoring client can implement such interfaces, so that upstream processing such as
Turbo DTO building may be coded based on the common interfaces without needing to adapt every time a new type of target
is introduced.

Other than the monitoring interfaces, this library also provides a [simple implementation](pkg/repository/simpleRepo)
of the metric repository, as well as a monitoring client for [Prometheus](https://prometheus.io/), as one of the first
implementations of such interfaces.

## Metric Repository

In this library, a metric repository is organized by entity.  Each entity has a set of metrics.  Each metric is a
key-value pair with key being the combination of the resource and the metric property, and value being a float64.

For example, Node '1.2.3.4' has memory usage of 3GB is represented in this library's model as follows:
* EntityType: Node
* EntityId: 1.2.3.4
* ResourceType: MEM
* MetricPropertyType: Used
* MetricValue: 3GB

This library defines a set of interfaces to manage/access the metric repository, including get/set metric values.

## Monitoring Template

Monitoring template is a set of metric meta data used to drive the metric collection.  Metric meta data is composed of
entity type, resource type, metric property type, and a metric setter that defines how the value is set in the metric
repository.  This library provides a default metric setter that simply puts the value into the repository, though
other use cases may exist to have a custom setter.

## Prometheus Monitoring Client

This library also provides a monitoring client implementation for [Prometheus](https://prometheus.io/).  The
[Prometheus monitoring client](pkg/prometheus) is equipped with a
[`MetricQueryMap`](pkg/prometheus/prometheus_queries.go) that defines what query to use for a defined metric and how
the entity id is obtained from the query result.  Currently, the following metric queries are supported:
* Node-level CPU/memory/network stats from the [Prometheus node exporter](https://github.com/prometheus/node_exporter)
* POD CPU/memory/disk stats from [Kubernetes](https://github.com/kubernetes/kubernetes).

To add support for more metric queries, one can simply add more entries to the map and implement new get-entity-id
functions as necessary.

To use this monitoring client, please refer to the test function
[`TestPrometheusMonitor()`](pkg/prometheus/prometheus_monitoring_client_test.go).  The major steps are:
1. Define a monitoring template to tell Prometheus what metrics to collect.
2. Define the list of entities for Prometheus to monitor.  The entity id must match what Prometheus is able to get from
the corresponding query results.
3. Instantiate `PrometheusMonitor` - the Prometheus monitoring client and call the `Monitor()` method.
4. The test code dumps out all the collected metrics.


### Test with minikube and kube-prometheus

To try out the test function, one way is to set up the test environment by installing
[`minikube`](https://github.com/kubernetes/minikube) and then
[`kube-prometheus`](https://github.com/coreos/prometheus-operator/tree/master/contrib/kube-prometheus).  The former
sets up a local mini Kubernetes cluster, while the latter deploys Prometheus components including the node exporter
into the cluster.

Once the test environment is set up, please run the `TestPrometheusMonitor()`.  If needed, customize the Prometheus
server address, the monitoring template, and the list of entities.

### Test with Prometheus server and exporters

Another way to test is to install [Prometheus server](https://prometheus.io/docs/introduction/install/) and selected
exporters such as the [node exporter](https://github.com/prometheus/node_exporter) yourself.
