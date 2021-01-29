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

package externalversions

import (
	"fmt"

	turtle "git.cloud2go.cn/reedition/turtle/api/pkg/apis/turtle"
	v1 "git.cloud2go.cn/reedition/turtle/api/pkg/apis/turtle/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=turtle, Version=turtle
	case turtle.SchemeGroupVersion.WithResource("clusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().Turtle().Clusters().Informer()}, nil
	case turtle.SchemeGroupVersion.WithResource("cronjobs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().Turtle().CronJobs().Informer()}, nil
	case turtle.SchemeGroupVersion.WithResource("daemonsets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().Turtle().DaemonSets().Informer()}, nil
	case turtle.SchemeGroupVersion.WithResource("datacenters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().Turtle().DataCenters().Informer()}, nil
	case turtle.SchemeGroupVersion.WithResource("deployments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().Turtle().Deployments().Informer()}, nil
	case turtle.SchemeGroupVersion.WithResource("events"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().Turtle().Events().Informer()}, nil
	case turtle.SchemeGroupVersion.WithResource("hpas"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().Turtle().HPAs().Informer()}, nil
	case turtle.SchemeGroupVersion.WithResource("ingresses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().Turtle().Ingresses().Informer()}, nil
	case turtle.SchemeGroupVersion.WithResource("jobs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().Turtle().Jobs().Informer()}, nil
	case turtle.SchemeGroupVersion.WithResource("namespaces"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().Turtle().Namespaces().Informer()}, nil
	case turtle.SchemeGroupVersion.WithResource("nodes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().Turtle().Nodes().Informer()}, nil
	case turtle.SchemeGroupVersion.WithResource("pvs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().Turtle().PVs().Informer()}, nil
	case turtle.SchemeGroupVersion.WithResource("pvcs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().Turtle().PVCs().Informer()}, nil
	case turtle.SchemeGroupVersion.WithResource("pods"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().Turtle().Pods().Informer()}, nil
	case turtle.SchemeGroupVersion.WithResource("replicasets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().Turtle().ReplicaSets().Informer()}, nil
	case turtle.SchemeGroupVersion.WithResource("services"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().Turtle().Services().Informer()}, nil
	case turtle.SchemeGroupVersion.WithResource("statefulsets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().Turtle().StatefulSets().Informer()}, nil
	case turtle.SchemeGroupVersion.WithResource("zones"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().Turtle().Zones().Informer()}, nil

		// Group=turtle, Version=v1
	case v1.SchemeGroupVersion.WithResource("clusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().V1().Clusters().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("cronjobs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().V1().CronJobs().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("daemonsets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().V1().DaemonSets().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("datacenters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().V1().DataCenters().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("deployments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().V1().Deployments().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("events"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().V1().Events().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("hpas"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().V1().HPAs().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("ingresses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().V1().Ingresses().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("jobs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().V1().Jobs().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("namespaces"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().V1().Namespaces().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("nodes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().V1().Nodes().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("pvs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().V1().PVs().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("pvcs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().V1().PVCs().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("pods"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().V1().Pods().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("replicasets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().V1().ReplicaSets().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("services"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().V1().Services().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("statefulsets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().V1().StatefulSets().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("zones"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Turtle().V1().Zones().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}