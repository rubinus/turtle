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

// HPALister helps list HPAs.
type HPALister interface {
	// List lists all HPAs in the indexer.
	List(selector labels.Selector) (ret []*v1.HPA, err error)
	// Get retrieves the HPA from the index for a given name.
	Get(name string) (*v1.HPA, error)
	HPAListerExpansion
}

// hPALister implements the HPALister interface.
type hPALister struct {
	indexer cache.Indexer
}

// NewHPALister returns a new HPALister.
func NewHPALister(indexer cache.Indexer) HPALister {
	return &hPALister{indexer: indexer}
}

// List lists all HPAs in the indexer.
func (s *hPALister) List(selector labels.Selector) (ret []*v1.HPA, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.HPA))
	})
	return ret, err
}

// Get retrieves the HPA from the index for a given name.
func (s *hPALister) Get(name string) (*v1.HPA, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("hpa"), name)
	}
	return obj.(*v1.HPA), nil
}