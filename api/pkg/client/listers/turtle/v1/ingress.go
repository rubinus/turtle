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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "git.cloud2go.cn/reedition/turtle/api/pkg/apis/turtle/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IngressLister helps list Ingresses.
type IngressLister interface {
	// List lists all Ingresses in the indexer.
	List(selector labels.Selector) (ret []*v1.Ingress, err error)
	// Get retrieves the Ingress from the index for a given name.
	Get(name string) (*v1.Ingress, error)
	IngressListerExpansion
}

// ingressLister implements the IngressLister interface.
type ingressLister struct {
	indexer cache.Indexer
}

// NewIngressLister returns a new IngressLister.
func NewIngressLister(indexer cache.Indexer) IngressLister {
	return &ingressLister{indexer: indexer}
}

// List lists all Ingresses in the indexer.
func (s *ingressLister) List(selector labels.Selector) (ret []*v1.Ingress, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Ingress))
	})
	return ret, err
}

// Get retrieves the Ingress from the index for a given name.
func (s *ingressLister) Get(name string) (*v1.Ingress, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("ingress"), name)
	}
	return obj.(*v1.Ingress), nil
}
