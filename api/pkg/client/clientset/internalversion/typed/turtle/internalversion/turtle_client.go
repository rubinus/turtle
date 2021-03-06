/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	"git.cloud2go.cn/reedition/turtle/api/pkg/client/clientset/internalversion/scheme"
	rest "k8s.io/client-go/rest"
)

type TurtleInterface interface {
	RESTClient() rest.Interface
	ClustersGetter
	CronJobsGetter
	DaemonSetsGetter
	DataCentersGetter
	DeploymentsGetter
	EventsGetter
	HPAsGetter
	IngressesGetter
	JobsGetter
	NamespacesGetter
	NodesGetter
	PVsGetter
	PVCsGetter
	PodsGetter
	ReplicaSetsGetter
	ServicesGetter
	StatefulSetsGetter
	ZonesGetter
}

// TurtleClient is used to interact with features provided by the turtle group.
type TurtleClient struct {
	restClient rest.Interface
}

func (c *TurtleClient) Clusters() ClusterInterface {
	return newClusters(c)
}

func (c *TurtleClient) CronJobs() CronJobInterface {
	return newCronJobs(c)
}

func (c *TurtleClient) DaemonSets() DaemonSetInterface {
	return newDaemonSets(c)
}

func (c *TurtleClient) DataCenters() DataCenterInterface {
	return newDataCenters(c)
}

func (c *TurtleClient) Deployments() DeploymentInterface {
	return newDeployments(c)
}

func (c *TurtleClient) Events() EventInterface {
	return newEvents(c)
}

func (c *TurtleClient) HPAs() HPAInterface {
	return newHPAs(c)
}

func (c *TurtleClient) Ingresses() IngressInterface {
	return newIngresses(c)
}

func (c *TurtleClient) Jobs() JobInterface {
	return newJobs(c)
}

func (c *TurtleClient) Namespaces() NamespaceInterface {
	return newNamespaces(c)
}

func (c *TurtleClient) Nodes() NodeInterface {
	return newNodes(c)
}

func (c *TurtleClient) PVs() PVInterface {
	return newPVs(c)
}

func (c *TurtleClient) PVCs() PVCInterface {
	return newPVCs(c)
}

func (c *TurtleClient) Pods() PodInterface {
	return newPods(c)
}

func (c *TurtleClient) ReplicaSets() ReplicaSetInterface {
	return newReplicaSets(c)
}

func (c *TurtleClient) Services() ServiceInterface {
	return newServices(c)
}

func (c *TurtleClient) StatefulSets() StatefulSetInterface {
	return newStatefulSets(c)
}

func (c *TurtleClient) Zones() ZoneInterface {
	return newZones(c)
}

// NewForConfig creates a new TurtleClient for the given config.
func NewForConfig(c *rest.Config) (*TurtleClient, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &TurtleClient{client}, nil
}

// NewForConfigOrDie creates a new TurtleClient for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *TurtleClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new TurtleClient for the given RESTClient.
func New(c rest.Interface) *TurtleClient {
	return &TurtleClient{c}
}

func setConfigDefaults(config *rest.Config) error {
	config.APIPath = "/apis"
	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}
	if config.GroupVersion == nil || config.GroupVersion.Group != scheme.Scheme.PrioritizedVersionsForGroup("turtle")[0].Group {
		gv := scheme.Scheme.PrioritizedVersionsForGroup("turtle")[0]
		config.GroupVersion = &gv
	}
	config.NegotiatedSerializer = scheme.Codecs

	if config.QPS == 0 {
		config.QPS = 5
	}
	if config.Burst == 0 {
		config.Burst = 10
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *TurtleClient) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
