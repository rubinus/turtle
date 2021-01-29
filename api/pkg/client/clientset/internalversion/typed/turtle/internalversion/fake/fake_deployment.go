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
	turtle "git.cloud2go.cn/reedition/turtle/api/pkg/apis/turtle"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDeployments implements DeploymentInterface
type FakeDeployments struct {
	Fake *FakeTurtle
}

var deploymentsResource = schema.GroupVersionResource{Group: "turtle", Version: "", Resource: "deployments"}

var deploymentsKind = schema.GroupVersionKind{Group: "turtle", Version: "", Kind: "Deployment"}

// Get takes name of the deployment, and returns the corresponding deployment object, and an error if there is any.
func (c *FakeDeployments) Get(name string, options v1.GetOptions) (result *turtle.Deployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(deploymentsResource, name), &turtle.Deployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*turtle.Deployment), err
}

// List takes label and field selectors, and returns the list of Deployments that match those selectors.
func (c *FakeDeployments) List(opts v1.ListOptions) (result *turtle.DeploymentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(deploymentsResource, deploymentsKind, opts), &turtle.DeploymentList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &turtle.DeploymentList{ListMeta: obj.(*turtle.DeploymentList).ListMeta}
	for _, item := range obj.(*turtle.DeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested deployments.
func (c *FakeDeployments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(deploymentsResource, opts))
}

// Create takes the representation of a deployment and creates it.  Returns the server's representation of the deployment, and an error, if there is any.
func (c *FakeDeployments) Create(deployment *turtle.Deployment) (result *turtle.Deployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(deploymentsResource, deployment), &turtle.Deployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*turtle.Deployment), err
}

// Update takes the representation of a deployment and updates it. Returns the server's representation of the deployment, and an error, if there is any.
func (c *FakeDeployments) Update(deployment *turtle.Deployment) (result *turtle.Deployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(deploymentsResource, deployment), &turtle.Deployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*turtle.Deployment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDeployments) UpdateStatus(deployment *turtle.Deployment) (*turtle.Deployment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(deploymentsResource, "status", deployment), &turtle.Deployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*turtle.Deployment), err
}

// Delete takes name of the deployment and deletes it. Returns an error if one occurs.
func (c *FakeDeployments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(deploymentsResource, name), &turtle.Deployment{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDeployments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(deploymentsResource, listOptions)

	_, err := c.Fake.Invokes(action, &turtle.DeploymentList{})
	return err
}

// Patch applies the patch and returns the patched deployment.
func (c *FakeDeployments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *turtle.Deployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(deploymentsResource, name, pt, data, subresources...), &turtle.Deployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*turtle.Deployment), err
}
