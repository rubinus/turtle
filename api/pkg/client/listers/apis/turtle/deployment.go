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

// DeploymentLister helps list Deployments.
type DeploymentLister interface {
	// List lists all Deployments in the indexer.
	List(selector labels.Selector) (ret []*turtle.Deployment, err error)
	// Get retrieves the Deployment from the index for a given name.
	Get(name string) (*turtle.Deployment, error)
	DeploymentListerExpansion
}

// deploymentLister implements the DeploymentLister interface.
type deploymentLister struct {
	indexer cache.Indexer
}

// NewDeploymentLister returns a new DeploymentLister.
func NewDeploymentLister(indexer cache.Indexer) DeploymentLister {
	return &deploymentLister{indexer: indexer}
}

// List lists all Deployments in the indexer.
func (s *deploymentLister) List(selector labels.Selector) (ret []*turtle.Deployment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*turtle.Deployment))
	})
	return ret, err
}

// Get retrieves the Deployment from the index for a given name.
func (s *deploymentLister) Get(name string) (*turtle.Deployment, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(turtle.Resource("deployment"), name)
	}
	return obj.(*turtle.Deployment), nil
}
