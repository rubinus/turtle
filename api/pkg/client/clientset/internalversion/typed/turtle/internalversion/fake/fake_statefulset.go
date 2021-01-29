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

// FakeStatefulSets implements StatefulSetInterface
type FakeStatefulSets struct {
	Fake *FakeTurtle
}

var statefulsetsResource = schema.GroupVersionResource{Group: "turtle", Version: "", Resource: "statefulsets"}

var statefulsetsKind = schema.GroupVersionKind{Group: "turtle", Version: "", Kind: "StatefulSet"}

// Get takes name of the statefulSet, and returns the corresponding statefulSet object, and an error if there is any.
func (c *FakeStatefulSets) Get(name string, options v1.GetOptions) (result *turtle.StatefulSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(statefulsetsResource, name), &turtle.StatefulSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*turtle.StatefulSet), err
}

// List takes label and field selectors, and returns the list of StatefulSets that match those selectors.
func (c *FakeStatefulSets) List(opts v1.ListOptions) (result *turtle.StatefulSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(statefulsetsResource, statefulsetsKind, opts), &turtle.StatefulSetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &turtle.StatefulSetList{ListMeta: obj.(*turtle.StatefulSetList).ListMeta}
	for _, item := range obj.(*turtle.StatefulSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested statefulSets.
func (c *FakeStatefulSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(statefulsetsResource, opts))
}

// Create takes the representation of a statefulSet and creates it.  Returns the server's representation of the statefulSet, and an error, if there is any.
func (c *FakeStatefulSets) Create(statefulSet *turtle.StatefulSet) (result *turtle.StatefulSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(statefulsetsResource, statefulSet), &turtle.StatefulSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*turtle.StatefulSet), err
}

// Update takes the representation of a statefulSet and updates it. Returns the server's representation of the statefulSet, and an error, if there is any.
func (c *FakeStatefulSets) Update(statefulSet *turtle.StatefulSet) (result *turtle.StatefulSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(statefulsetsResource, statefulSet), &turtle.StatefulSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*turtle.StatefulSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStatefulSets) UpdateStatus(statefulSet *turtle.StatefulSet) (*turtle.StatefulSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(statefulsetsResource, "status", statefulSet), &turtle.StatefulSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*turtle.StatefulSet), err
}

// Delete takes name of the statefulSet and deletes it. Returns an error if one occurs.
func (c *FakeStatefulSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(statefulsetsResource, name), &turtle.StatefulSet{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStatefulSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(statefulsetsResource, listOptions)

	_, err := c.Fake.Invokes(action, &turtle.StatefulSetList{})
	return err
}

// Patch applies the patch and returns the patched statefulSet.
func (c *FakeStatefulSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *turtle.StatefulSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(statefulsetsResource, name, pt, data, subresources...), &turtle.StatefulSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*turtle.StatefulSet), err
}
