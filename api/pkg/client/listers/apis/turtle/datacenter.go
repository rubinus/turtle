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

package turtle

import (
	turtle "git.cloud2go.cn/reedition/turtle/api/pkg/apis/turtle"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DataCenterLister helps list DataCenters.
type DataCenterLister interface {
	// List lists all DataCenters in the indexer.
	List(selector labels.Selector) (ret []*turtle.DataCenter, err error)
	// Get retrieves the DataCenter from the index for a given name.
	Get(name string) (*turtle.DataCenter, error)
	DataCenterListerExpansion
}

// dataCenterLister implements the DataCenterLister interface.
type dataCenterLister struct {
	indexer cache.Indexer
}

// NewDataCenterLister returns a new DataCenterLister.
func NewDataCenterLister(indexer cache.Indexer) DataCenterLister {
	return &dataCenterLister{indexer: indexer}
}

// List lists all DataCenters in the indexer.
func (s *dataCenterLister) List(selector labels.Selector) (ret []*turtle.DataCenter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*turtle.DataCenter))
	})
	return ret, err
}

// Get retrieves the DataCenter from the index for a given name.
func (s *dataCenterLister) Get(name string) (*turtle.DataCenter, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(turtle.Resource("datacenter"), name)
	}
	return obj.(*turtle.DataCenter), nil
}
