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

// FakePVs implements PVInterface
type FakePVs struct {
	Fake *FakeTurtle
}

var pvsResource = schema.GroupVersionResource{Group: "turtle", Version: "", Resource: "pvs"}

var pvsKind = schema.GroupVersionKind{Group: "turtle", Version: "", Kind: "PV"}

// Get takes name of the pV, and returns the corresponding pV object, and an error if there is any.
func (c *FakePVs) Get(name string, options v1.GetOptions) (result *turtle.PV, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(pvsResource, name), &turtle.PV{})
	if obj == nil {
		return nil, err
	}
	return obj.(*turtle.PV), err
}

// List takes label and field selectors, and returns the list of PVs that match those selectors.
func (c *FakePVs) List(opts v1.ListOptions) (result *turtle.PVList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(pvsResource, pvsKind, opts), &turtle.PVList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &turtle.PVList{ListMeta: obj.(*turtle.PVList).ListMeta}
	for _, item := range obj.(*turtle.PVList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pVs.
func (c *FakePVs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(pvsResource, opts))
}

// Create takes the representation of a pV and creates it.  Returns the server's representation of the pV, and an error, if there is any.
func (c *FakePVs) Create(pV *turtle.PV) (result *turtle.PV, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(pvsResource, pV), &turtle.PV{})
	if obj == nil {
		return nil, err
	}
	return obj.(*turtle.PV), err
}

// Update takes the representation of a pV and updates it. Returns the server's representation of the pV, and an error, if there is any.
func (c *FakePVs) Update(pV *turtle.PV) (result *turtle.PV, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(pvsResource, pV), &turtle.PV{})
	if obj == nil {
		return nil, err
	}
	return obj.(*turtle.PV), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePVs) UpdateStatus(pV *turtle.PV) (*turtle.PV, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(pvsResource, "status", pV), &turtle.PV{})
	if obj == nil {
		return nil, err
	}
	return obj.(*turtle.PV), err
}

// Delete takes name of the pV and deletes it. Returns an error if one occurs.
func (c *FakePVs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(pvsResource, name), &turtle.PV{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePVs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(pvsResource, listOptions)

	_, err := c.Fake.Invokes(action, &turtle.PVList{})
	return err
}

// Patch applies the patch and returns the patched pV.
func (c *FakePVs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *turtle.PV, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(pvsResource, name, pt, data, subresources...), &turtle.PV{})
	if obj == nil {
		return nil, err
	}
	return obj.(*turtle.PV), err
}
