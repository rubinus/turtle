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

package fake

import (
	v1 "git.cloud2go.cn/reedition/turtle/api/pkg/client/clientset/versioned/typed/turtle/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeTurtleV1 struct {
	*testing.Fake
}

func (c *FakeTurtleV1) Clusters() v1.ClusterInterface {
	return &FakeClusters{c}
}

func (c *FakeTurtleV1) CronJobs() v1.CronJobInterface {
	return &FakeCronJobs{c}
}

func (c *FakeTurtleV1) DaemonSets() v1.DaemonSetInterface {
	return &FakeDaemonSets{c}
}

func (c *FakeTurtleV1) DataCenters() v1.DataCenterInterface {
	return &FakeDataCenters{c}
}

func (c *FakeTurtleV1) Deployments() v1.DeploymentInterface {
	return &FakeDeployments{c}
}

func (c *FakeTurtleV1) Events() v1.EventInterface {
	return &FakeEvents{c}
}

func (c *FakeTurtleV1) HPAs() v1.HPAInterface {
	return &FakeHPAs{c}
}

func (c *FakeTurtleV1) Ingresses() v1.IngressInterface {
	return &FakeIngresses{c}
}

func (c *FakeTurtleV1) Jobs() v1.JobInterface {
	return &FakeJobs{c}
}

func (c *FakeTurtleV1) Namespaces() v1.NamespaceInterface {
	return &FakeNamespaces{c}
}

func (c *FakeTurtleV1) Nodes() v1.NodeInterface {
	return &FakeNodes{c}
}

func (c *FakeTurtleV1) PVs() v1.PVInterface {
	return &FakePVs{c}
}

func (c *FakeTurtleV1) PVCs() v1.PVCInterface {
	return &FakePVCs{c}
}

func (c *FakeTurtleV1) Pods() v1.PodInterface {
	return &FakePods{c}
}

func (c *FakeTurtleV1) ReplicaSets() v1.ReplicaSetInterface {
	return &FakeReplicaSets{c}
}

func (c *FakeTurtleV1) Services() v1.ServiceInterface {
	return &FakeServices{c}
}

func (c *FakeTurtleV1) StatefulSets() v1.StatefulSetInterface {
	return &FakeStatefulSets{c}
}

func (c *FakeTurtleV1) Zones() v1.ZoneInterface {
	return &FakeZones{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeTurtleV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
