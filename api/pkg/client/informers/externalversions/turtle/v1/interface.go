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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "git.cloud2go.cn/reedition/turtle/api/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Clusters returns a ClusterInformer.
	Clusters() ClusterInformer
	// CronJobs returns a CronJobInformer.
	CronJobs() CronJobInformer
	// DaemonSets returns a DaemonSetInformer.
	DaemonSets() DaemonSetInformer
	// DataCenters returns a DataCenterInformer.
	DataCenters() DataCenterInformer
	// Deployments returns a DeploymentInformer.
	Deployments() DeploymentInformer
	// Events returns a EventInformer.
	Events() EventInformer
	// HPAs returns a HPAInformer.
	HPAs() HPAInformer
	// Ingresses returns a IngressInformer.
	Ingresses() IngressInformer
	// Jobs returns a JobInformer.
	Jobs() JobInformer
	// Namespaces returns a NamespaceInformer.
	Namespaces() NamespaceInformer
	// Nodes returns a NodeInformer.
	Nodes() NodeInformer
	// PVs returns a PVInformer.
	PVs() PVInformer
	// PVCs returns a PVCInformer.
	PVCs() PVCInformer
	// Pods returns a PodInformer.
	Pods() PodInformer
	// ReplicaSets returns a ReplicaSetInformer.
	ReplicaSets() ReplicaSetInformer
	// Services returns a ServiceInformer.
	Services() ServiceInformer
	// StatefulSets returns a StatefulSetInformer.
	StatefulSets() StatefulSetInformer
	// Zones returns a ZoneInformer.
	Zones() ZoneInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Clusters returns a ClusterInformer.
func (v *version) Clusters() ClusterInformer {
	return &clusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// CronJobs returns a CronJobInformer.
func (v *version) CronJobs() CronJobInformer {
	return &cronJobInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// DaemonSets returns a DaemonSetInformer.
func (v *version) DaemonSets() DaemonSetInformer {
	return &daemonSetInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// DataCenters returns a DataCenterInformer.
func (v *version) DataCenters() DataCenterInformer {
	return &dataCenterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Deployments returns a DeploymentInformer.
func (v *version) Deployments() DeploymentInformer {
	return &deploymentInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Events returns a EventInformer.
func (v *version) Events() EventInformer {
	return &eventInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// HPAs returns a HPAInformer.
func (v *version) HPAs() HPAInformer {
	return &hPAInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Ingresses returns a IngressInformer.
func (v *version) Ingresses() IngressInformer {
	return &ingressInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Jobs returns a JobInformer.
func (v *version) Jobs() JobInformer {
	return &jobInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Namespaces returns a NamespaceInformer.
func (v *version) Namespaces() NamespaceInformer {
	return &namespaceInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Nodes returns a NodeInformer.
func (v *version) Nodes() NodeInformer {
	return &nodeInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// PVs returns a PVInformer.
func (v *version) PVs() PVInformer {
	return &pVInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// PVCs returns a PVCInformer.
func (v *version) PVCs() PVCInformer {
	return &pVCInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Pods returns a PodInformer.
func (v *version) Pods() PodInformer {
	return &podInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ReplicaSets returns a ReplicaSetInformer.
func (v *version) ReplicaSets() ReplicaSetInformer {
	return &replicaSetInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Services returns a ServiceInformer.
func (v *version) Services() ServiceInformer {
	return &serviceInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// StatefulSets returns a StatefulSetInformer.
func (v *version) StatefulSets() StatefulSetInformer {
	return &statefulSetInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Zones returns a ZoneInformer.
func (v *version) Zones() ZoneInformer {
	return &zoneInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
