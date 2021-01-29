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

// DaemonSetLister helps list DaemonSets.
type DaemonSetLister interface {
	// List lists all DaemonSets in the indexer.
	List(selector labels.Selector) (ret []*v1.DaemonSet, err error)
	// Get retrieves the DaemonSet from the index for a given name.
	Get(name string) (*v1.DaemonSet, error)
	DaemonSetListerExpansion
}

// daemonSetLister implements the DaemonSetLister interface.
type daemonSetLister struct {
	indexer cache.Indexer
}

// NewDaemonSetLister returns a new DaemonSetLister.
func NewDaemonSetLister(indexer cache.Indexer) DaemonSetLister {
	return &daemonSetLister{indexer: indexer}
}

// List lists all DaemonSets in the indexer.
func (s *daemonSetLister) List(selector labels.Selector) (ret []*v1.DaemonSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.DaemonSet))
	})
	return ret, err
}

// Get retrieves the DaemonSet from the index for a given name.
func (s *daemonSetLister) Get(name string) (*v1.DaemonSet, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("daemonset"), name)
	}
	return obj.(*v1.DaemonSet), nil
}
