# turbo-go-monitoring

This is a GO monitoring library that defines a set of monitoring interfaces, including a metric repository and a
monitoring template.  Any target-specific monitoring client can implement such interfaces, so that upstream processing
such as Turbo DTO building may be coded based on the common interfaces without needing to adapt every time a new type
of target is introduced.

Other than the monitoring interfaces, this library also provides a simple implementation of the metric repository as
well as a monitoring client for [Prometheus](https://prometheus.io/), as one of the first implementations of such
interfaces.

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

This library also provides a monitoring client implementation for [Prometheus](https://prometheus.io/).  The test
function `TestPrometheusMonitor()` illustrates how the Prometheus client can be used to collect metrics.  Steps are:
1. Define a monitoring template to tell Prometheus what metrics to collect.
2. Define the list of entities for Prometheus to monitor, and put them into the metric repository.
3. Instantiate `PrometheusMonitor` - the Prometheus monitoring client and call the `Monitor()` method.
4. The test code dumps out all the collected metrics.

The Prometheus monitoring client is equipped with a `MetricQueryMap` in `prometheus_queries.go`.  The `MetricQueryMap`
defines what query to use for a defined metric.  It currently supports the following metric queries:
* Node-level CPU/memory/network stats from the [Prometheus node exporter](https://github.com/prometheus/node_exporter)
* POD CPU/memory/disk stats from [Kubernetes](https://github.com/kubernetes/kubernetes).

### Test with minikube and kube-prometheus

To try out the test function, one can set up the test environment by installing
[`minikube`](https://github.com/kubernetes/minikube) and then
[`kube-prometheus`](https://github.com/coreos/prometheus-operator/tree/master/contrib/kube-prometheus).  The former
sets up a local mini Kubernetes cluster, while the latter deploys Prometheus components including the node exporter
into the cluster.

Once the test environment is set up, we can run the `TestPrometheusMonitor()`.  If needed, customize the monitoring
template, the list of entities, and the Prometheus server address.