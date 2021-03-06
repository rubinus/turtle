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
	turtlev1 "git.cloud2go.cn/reedition/turtle/api/pkg/apis/turtle/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePVCs implements PVCInterface
type FakePVCs struct {
	Fake *FakeTurtleV1
}

var pvcsResource = schema.GroupVersionResource{Group: "turtle", Version: "v1", Resource: "pvcs"}

var pvcsKind = schema.GroupVersionKind{Group: "turtle", Version: "v1", Kind: "PVC"}

// Get takes name of the pVC, and returns the corresponding pVC object, and an error if there is any.
func (c *FakePVCs) Get(name string, options v1.GetOptions) (result *turtlev1.PVC, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(pvcsResource, name), &turtlev1.PVC{})
	if obj == nil {
		return nil, err
	}
	return obj.(*turtlev1.PVC), err
}

// List takes label and field selectors, and returns the list of PVCs that match those selectors.
func (c *FakePVCs) List(opts v1.ListOptions) (result *turtlev1.PVCList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(pvcsResource, pvcsKind, opts), &turtlev1.PVCList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &turtlev1.PVCList{ListMeta: obj.(*turtlev1.PVCList).ListMeta}
	for _, item := range obj.(*turtlev1.PVCList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pVCs.
func (c *FakePVCs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(pvcsResource, opts))
}

// Create takes the representation of a pVC and creates it.  Returns the server's representation of the pVC, and an error, if there is any.
func (c *FakePVCs) Create(pVC *turtlev1.PVC) (result *turtlev1.PVC, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(pvcsResource, pVC), &turtlev1.PVC{})
	if obj == nil {
		return nil, err
	}
	return obj.(*turtlev1.PVC), err
}

// Update takes the representation of a pVC and updates it. Returns the server's representation of the pVC, and an error, if there is any.
func (c *FakePVCs) Update(pVC *turtlev1.PVC) (result *turtlev1.PVC, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(pvcsResource, pVC), &turtlev1.PVC{})
	if obj == nil {
		return nil, err
	}
	return obj.(*turtlev1.PVC), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePVCs) UpdateStatus(pVC *turtlev1.PVC) (*turtlev1.PVC, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(pvcsResource, "status", pVC), &turtlev1.PVC{})
	if obj == nil {
		return nil, err
	}
	return obj.(*turtlev1.PVC), err
}

// Delete takes name of the pVC and deletes it. Returns an error if one occurs.
func (c *FakePVCs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(pvcsResource, name), &turtlev1.PVC{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePVCs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(pvcsResource, listOptions)

	_, err := c.Fake.Invokes(action, &turtlev1.PVCList{})
	return err
}

// Patch applies the patch and returns the patched pVC.
func (c *FakePVCs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *turtlev1.PVC, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(pvcsResource, name, pt, data, subresources...), &turtlev1.PVC{})
	if obj == nil {
		return nil, err
	}
	return obj.(*turtlev1.PVC), err
}
