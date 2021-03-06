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

// FakeReplicaSets implements ReplicaSetInterface
type FakeReplicaSets struct {
	Fake *FakeTurtle
}

var replicasetsResource = schema.GroupVersionResource{Group: "turtle", Version: "", Resource: "replicasets"}

var replicasetsKind = schema.GroupVersionKind{Group: "turtle", Version: "", Kind: "ReplicaSet"}

// Get takes name of the replicaSet, and returns the corresponding replicaSet object, and an error if there is any.
func (c *FakeReplicaSets) Get(name string, options v1.GetOptions) (result *turtle.ReplicaSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(replicasetsResource, name), &turtle.ReplicaSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*turtle.ReplicaSet), err
}

// List takes label and field selectors, and returns the list of ReplicaSets that match those selectors.
func (c *FakeReplicaSets) List(opts v1.ListOptions) (result *turtle.ReplicaSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(replicasetsResource, replicasetsKind, opts), &turtle.ReplicaSetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &turtle.ReplicaSetList{ListMeta: obj.(*turtle.ReplicaSetList).ListMeta}
	for _, item := range obj.(*turtle.ReplicaSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested replicaSets.
func (c *FakeReplicaSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(replicasetsResource, opts))
}

// Create takes the representation of a replicaSet and creates it.  Returns the server's representation of the replicaSet, and an error, if there is any.
func (c *FakeReplicaSets) Create(replicaSet *turtle.ReplicaSet) (result *turtle.ReplicaSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(replicasetsResource, replicaSet), &turtle.ReplicaSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*turtle.ReplicaSet), err
}

// Update takes the representation of a replicaSet and updates it. Returns the server's representation of the replicaSet, and an error, if there is any.
func (c *FakeReplicaSets) Update(replicaSet *turtle.ReplicaSet) (result *turtle.ReplicaSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(replicasetsResource, replicaSet), &turtle.ReplicaSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*turtle.ReplicaSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeReplicaSets) UpdateStatus(replicaSet *turtle.ReplicaSet) (*turtle.ReplicaSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(replicasetsResource, "status", replicaSet), &turtle.ReplicaSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*turtle.ReplicaSet), err
}

// Delete takes name of the replicaSet and deletes it. Returns an error if one occurs.
func (c *FakeReplicaSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(replicasetsResource, name), &turtle.ReplicaSet{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeReplicaSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(replicasetsResource, listOptions)

	_, err := c.Fake.Invokes(action, &turtle.ReplicaSetList{})
	return err
}

// Patch applies the patch and returns the patched replicaSet.
func (c *FakeReplicaSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *turtle.ReplicaSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(replicasetsResource, name, pt, data, subresources...), &turtle.ReplicaSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*turtle.ReplicaSet), err
}
