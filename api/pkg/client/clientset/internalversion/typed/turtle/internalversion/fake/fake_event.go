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

// FakeEvents implements EventInterface
type FakeEvents struct {
	Fake *FakeTurtle
}

var eventsResource = schema.GroupVersionResource{Group: "turtle", Version: "", Resource: "events"}

var eventsKind = schema.GroupVersionKind{Group: "turtle", Version: "", Kind: "Event"}

// Get takes name of the event, and returns the corresponding event object, and an error if there is any.
func (c *FakeEvents) Get(name string, options v1.GetOptions) (result *turtle.Event, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(eventsResource, name), &turtle.Event{})
	if obj == nil {
		return nil, err
	}
	return obj.(*turtle.Event), err
}

// List takes label and field selectors, and returns the list of Events that match those selectors.
func (c *FakeEvents) List(opts v1.ListOptions) (result *turtle.EventList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(eventsResource, eventsKind, opts), &turtle.EventList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &turtle.EventList{ListMeta: obj.(*turtle.EventList).ListMeta}
	for _, item := range obj.(*turtle.EventList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested events.
func (c *FakeEvents) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(eventsResource, opts))
}

// Create takes the representation of a event and creates it.  Returns the server's representation of the event, and an error, if there is any.
func (c *FakeEvents) Create(event *turtle.Event) (result *turtle.Event, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(eventsResource, event), &turtle.Event{})
	if obj == nil {
		return nil, err
	}
	return obj.(*turtle.Event), err
}

// Update takes the representation of a event and updates it. Returns the server's representation of the event, and an error, if there is any.
func (c *FakeEvents) Update(event *turtle.Event) (result *turtle.Event, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(eventsResource, event), &turtle.Event{})
	if obj == nil {
		return nil, err
	}
	return obj.(*turtle.Event), err
}

// Delete takes name of the event and deletes it. Returns an error if one occurs.
func (c *FakeEvents) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(eventsResource, name), &turtle.Event{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEvents) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(eventsResource, listOptions)

	_, err := c.Fake.Invokes(action, &turtle.EventList{})
	return err
}

// Patch applies the patch and returns the patched event.
func (c *FakeEvents) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *turtle.Event, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(eventsResource, name, pt, data, subresources...), &turtle.Event{})
	if obj == nil {
		return nil, err
	}
	return obj.(*turtle.Event), err
}
