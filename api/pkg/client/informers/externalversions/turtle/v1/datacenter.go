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

// DataCenterInformer provides access to a shared informer and lister for
// DataCenters.
type DataCenterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.DataCenterLister
}

type dataCenterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewDataCenterInformer constructs a new informer for DataCenter type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDataCenterInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDataCenterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredDataCenterInformer constructs a new informer for DataCenter type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDataCenterInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TurtleV1().DataCenters().List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TurtleV1().DataCenters().Watch(options)
			},
		},
		&turtlev1.DataCenter{},
		resyncPeriod,
		indexers,
	)
}

func (f *dataCenterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDataCenterInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *dataCenterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&turtlev1.DataCenter{}, f.defaultInformer)
}

func (f *dataCenterInformer) Lister() v1.DataCenterLister {
	return v1.NewDataCenterLister(f.Informer().GetIndexer())
}
