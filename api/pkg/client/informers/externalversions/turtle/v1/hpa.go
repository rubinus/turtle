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
	time "time"

	turtlev1 "git.cloud2go.cn/reedition/turtle/api/pkg/apis/turtle/v1"
	versioned "git.cloud2go.cn/reedition/turtle/api/pkg/client/clientset/versioned"
	internalinterfaces "git.cloud2go.cn/reedition/turtle/api/pkg/client/informers/externalversions/internalinterfaces"
	v1 "git.cloud2go.cn/reedition/turtle/api/pkg/client/listers/turtle/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// HPAInformer provides access to a shared informer and lister for
// HPAs.
type HPAInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.HPALister
}

type hPAInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewHPAInformer constructs a new informer for HPA type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewHPAInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredHPAInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredHPAInformer constructs a new informer for HPA type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredHPAInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TurtleV1().HPAs().List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TurtleV1().HPAs().Watch(options)
			},
		},
		&turtlev1.HPA{},
		resyncPeriod,
		indexers,
	)
}

func (f *hPAInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredHPAInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *hPAInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&turtlev1.HPA{}, f.defaultInformer)
}

func (f *hPAInformer) Lister() v1.HPALister {
	return v1.NewHPALister(f.Informer().GetIndexer())
}